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
	return p[i].points > p[j].points || (p[i].points == p[j].points && p[i].name > p[j].name)
}

func newTallyTable(r io.Reader) (tallyTable, error) {
	t := tallyTable{
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
		map[string]int{},
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Text()
		f := strings.Split(row, ";")
		//fmt.Printf("%#v\n", f)
		if len(f) < 3 {
			continue
		}
		t.mp[f[0]]++
		t.mp[f[1]]++
		switch f[2] {
		case "win":
			t.w[f[0]]++
			t.p[f[0]] += 3
			t.l[f[1]]++
		case "loss":
			t.l[f[0]]++
			t.w[f[1]]++
			t.p[f[1]] += 3
		case "draw":
			t.d[f[0]]++
			t.d[f[1]]++
			t.p[f[0]] += 1
			t.p[f[1]] += 1
		}
	}

	return t, nil
}

func (t tallyTable) print(w io.Writer) (err error) {

	// head
	_, err = w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	if err != nil {
		return err
	}

	// body
	for _, team := range t.sorted() {
		name := team.name
		w.Write([]byte(fmt.Sprintf("%s\t       |  %d |  %d |  %d |  %d |  %d\n",
			name,
			t.mp[name],
			t.w[name],
			t.d[name],
			t.l[name],
			t.p[name],
		)))
		if err != nil {
			return err
		}
	}

	// debug
	//w.Write([]byte(fmt.Sprintf("\nDebug table: %+v\n", t)))

	return nil
}

// sorted iterates through all participated teams (presented in "match played" map) and build slice for sort
func (t tallyTable) sorted() sortedTeams {
	slice := make(sortedTeams, 0, len(t.p))
	for name, _ := range t.mp {
		slice = append(slice, sortedTeam{name, t.p[name]})
	}
	sort.Sort(slice)
	return slice
}

// Tally creates tallyTable struct from input and prints formatted results to output.
func Tally(r io.Reader, w io.Writer) error {
	tally, err := newTallyTable(r)
	if err != nil {
		return err
	}
	return tally.print(w)
}
