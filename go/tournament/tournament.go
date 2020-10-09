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

func newTallyTable(r io.Reader) (tallyTable, error) {
	t := tallyTable{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Text()
		f := strings.Split(row, ";")
		fmt.Printf("%#v\n", f)
		// switch f[2]
	}
	t.w = map[string]int{"Allegoric Alaskians": 1, "Devastating Donkeys": 2}

	return t, nil
}

func (t tallyTable) print(w io.Writer) (err error) {
	
	_, err = w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(fmt.Sprintf("\nTable: %+v\n", t)))
	if err != nil {
		return err
	}

	return nil
}

// Tally creates tallyTable struct from input and prints formatted results to output.
func Tally(r io.Reader, w io.Writer) error {
	tally, err := newTallyTable(r)
	if err != nil {
		return err
	}
	return tally.print(w)
}
