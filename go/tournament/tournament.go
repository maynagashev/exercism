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

func (t team) won() team {
	t.matches++
	t.wins++
	t.points += 3
	return t
}
func (t team) lost() team {
	t.matches++
	t.losses++
	return t
}

func (t team) drew() team {
	t.matches++
	t.draws++
	t.points++
	return t
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

type teamMap map[string]team

func (teams teamMap) print(w io.Writer) (err error) {
	const format = "%-30v |%3v |%3v |%3v |%3v |%3v\n"

	_, err = fmt.Fprintf(w, format, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}
	for _, s := range teams.sorted() {
		t := teams[s.name]
		_, err := fmt.Fprintf(w, format, s.name, t.matches, t.wins, t.draws, t.losses, t.points)
		if err != nil {
			return err
		}
	}

	return nil
}

func (teams teamMap) sorted() sortedTeams {
	slice := make(sortedTeams, 0, len(teams))
	for name, team := range teams {
		slice = append(slice, sortedTeam{name, team.points})
	}
	sort.Sort(slice)
	return slice
}

func parse(r io.Reader) (teamMap, error) {
	teams := teamMap{}
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

		team1, team2, result := f[0], f[1], f[2]

		switch result {
		case "win":
			teams[team1] = teams[team1].won()
			teams[team2] = teams[team2].lost()
		case "loss":
			teams[team2] = teams[team2].won()
			teams[team1] = teams[team1].lost()
		case "draw":
			teams[team1] = teams[team1].drew()
			teams[team2] = teams[team2].drew()
		default:
			return teams, fmt.Errorf("Unrecongnized match result \"%s\" in a row:\n%s", result, row)
		}
	}

	return teams, nil
}

// Tally creates teams map with team structs
func Tally(r io.Reader, w io.Writer) error {
	teams, err := parse(r)
	if err != nil {
		return err
	}
	return teams.print(w)
}
