package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	matches, wins, losses, draws, points int
}
type TeamMap map[string]team

type sortedTeam struct {
	name   string
	points int
}
type League []sortedTeam

func (p League) Len() int      { return len(p) }
func (p League) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p League) Less(i, j int) bool {
	return p[i].points > p[j].points || (p[i].points == p[j].points && p[i].name < p[j].name)
}
func LeagueFromTeams(teams TeamMap) League {
	slice := make(League, 0, len(teams))
	for name, team := range teams {
		slice = append(slice, sortedTeam{name, team.points})
	}
	sort.Sort(slice)
	return slice
}

func ReadMatchDataFrom(r io.Reader) (TeamMap, error) {
	teams := TeamMap{}
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
			return teams, fmt.Errorf("Invalid row:\n%s", row)
		}

		name1, name2, result := f[0], f[1], f[2]

		a, b := teams[name1], teams[name2]

		a.matches++
		b.matches++
		switch result {
		case "win":
			a.wins++
			a.points += 3
			b.losses++
		case "loss":
			b.wins++
			b.points += 3
			a.losses++
		case "draw":
			a.draws++
			a.points++
			b.draws++
			b.points++
		default:
			return teams, fmt.Errorf("Unrecongnized match result \"%s\" in a row:\n%s", result, row)
		}
		teams[name1], teams[name2] = a, b
	}

	return teams, nil
}

func WriteLeagueTableTo(w io.Writer, league League) (err error) {
	const format = "%-30v |%3v |%3v |%3v |%3v |%3v\n"

	_, err = fmt.Fprintf(w, format, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}
	for _, s := range league {
		t := teams[s.name]
		_, err := fmt.Fprintf(w, format, s.name, t.matches, t.wins, t.draws, t.losses, t.points)
		if err != nil {
			return err
		}
	}

	return nil
}

// Tally creates teams map with team structs
func Tally(r io.Reader, w io.Writer) error {
	teams, err := ReadMatchDataFrom(r)
	if err != nil {
		return err
	}
	league := LeagueFromTeams(teams)
	return WriteLeagueTableTo(w, league)
}
