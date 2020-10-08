package tournament

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type tallyTable struct {
	mp, w, d, l, p map[string]int
}

func (tallyTable) print(w io.Writer) {

}

// Tally writes to specified writer the number of points, goals, runs, etc. achieved in a game or by a team.
func Tally(r io.Reader, w io.Writer) error {
	parse(r).print(w)
	return nil
}

func format(table tallyTable, w io.Writer) {
	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	w.Write([]byte(fmt.Sprintf("\nTable: %+v\n", table)))
}

func parse(r io.Reader) tallyTable {
	t := tallyTable{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Text()
		f := strings.Split(row, ";")
		fmt.Printf("%#v\n", f)
	}
	t.w = map[string]int{"Allegoric Alaskians": 1, "Devastating Donkeys": 2}

	return t
}
