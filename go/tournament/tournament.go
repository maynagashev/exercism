package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team representation of single team data
type Team struct {
	name                                 string
	matches, wins, losses, draws, points int
}

// ReadMatchDataFrom parses match data using team map
func ReadMatchDataFrom(r io.Reader) (map[string]Team, error) {
	teamMap := map[string]Team{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// parse row
		row := strings.TrimSpace(scanner.Text())
		if row == "" || strings.HasPrefix(row, "#") {
			continue
		}
		parts := strings.Split(row, ";")
		if len(parts) != 3 {
			return nil, fmt.Errorf("bad line %q: expected 'teamA;teamB;result'", row)
		}
		name1, name2, result := parts[0], parts[1], parts[2]

		// resolve both teams from a map, update attributes and write back to the map
		a, b := teamMap[name1], teamMap[name2]
		a.name, b.name = name1, name2
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
			return teamMap, fmt.Errorf("Unrecongnized match result \"%s\" in a row:\n%s", result, row)
		}
		teamMap[name1], teamMap[name2] = a, b
	}

	return teamMap, nil
}

// WriteTeamsAsTableTo prints teams as a table in given order
func WriteTeamsAsTableTo(w io.Writer, teams []Team) (err error) {
	const format = "%-30v |%3v |%3v |%3v |%3v |%3v\n"

	_, err = fmt.Fprintf(w, format, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}
	for _, t := range teams {
		_, err := fmt.Fprintf(w, format, t.name, t.matches, t.wins, t.draws, t.losses, t.points)
		if err != nil {
			return err
		}
	}

	return nil
}

// Tally reads a file of match results and outputs statistics for each team as formatted table.
func Tally(r io.Reader, w io.Writer) error {
	teamMap, err := ReadMatchDataFrom(r)
	if err != nil {
		return err
	}

	//sorting
	teams := make([]Team, 0, len(teamMap))
	for _, team := range teamMap {
		teams = append(teams, team)
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})

	return WriteTeamsAsTableTo(w, teams)
}
