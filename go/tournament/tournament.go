package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type tallyTable struct {
	mp, w, d, l, p map[string]int
}

type sortedTeam struct {
	name   string
	points int
}
type sortedTeams []sortedTeam

func (p sortedTeams) Len() int      { return len(p) }
func (p sortedTeams) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p sortedTeams) Less(i, j int) bool {
	return p[i].points > p[j].points || (p[i].points == p[j].points && p[i].name < p[j].name)
}

// Tally creates tallyTable struct from input and prints formatted results to output.
func Tally(r io.Reader, w io.Writer) error {
	tt, err := parse(r)
	if err != nil {
		return err
	}
	//w.Write([]byte(fmt.Sprintf("\nParsed tally table: %+v\n", tt)))
	return tt.print(w)
}

func parse(r io.Reader) (tallyTable, error) {
	t := tallyTable{
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())
		f := strings.Split(row, ";")

		// validation
		switch {
		case row == "": // skip empty rows
			continue
		case row[0] == '#': // skip comments
			continue
		case len(f) < 3:
			return t, fmt.Errorf("Invalid row:\n%s", row)
		case f[2] != "win" && f[2] != "loss" && f[2] != "draw":
			return t, fmt.Errorf("Unrecongnized match result \"%s\" in a row:\n%s", f[2], row)
		}

		team1, team2, result := f[0], f[1], f[2]
		t.mp[team1]++
		t.mp[team2]++
		switch result {
		case "win":
			t.w[team1]++
			t.l[team2]++
			t.p[team1] += 3
		case "loss":
			t.w[team2]++
			t.l[team1]++
			t.p[team2] += 3
		case "draw":
			t.d[team1]++
			t.d[team2]++
			t.p[team1]++
			t.p[team2]++
		}
	}

	return t, nil
}

func (t tallyTable) print(w io.Writer) (err error) {
	const format = "%-30v |%3v |%3v |%3v |%3v |%3v\n"

	// head
	_, err = fmt.Fprintf(w, format, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}

	// body
	for _, team := range t.sorted() {
		name := team.name
		_, err := fmt.Fprintf(w, format,
			name,
			t.mp[name],
			t.w[name],
			t.d[name],
			t.l[name],
			t.p[name],
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// sorted sorts given tallyTable and returns sortedTeams slice of sortedTeam struct
func (t tallyTable) sorted() sortedTeams {
	slice := make(sortedTeams, 0, len(t.p))
	for name := range t.mp {
		slice = append(slice, sortedTeam{name, t.p[name]})
	}
	sort.Sort(slice)
	return slice
}
