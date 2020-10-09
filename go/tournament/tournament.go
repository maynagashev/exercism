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
		fmt.Printf("%#v\n", f)
		if len(f) < 3 {
			continue
		}
		t.mp[f[0]]++
		t.mp[f[1]]++
		switch f[2] {
		case "win":
			t.w[f[0]]++
			t.l[f[1]]++
		case "loss":
			t.l[f[0]]++
			t.w[f[1]]++
		case "draw":
			t.d[f[0]]++
			t.d[f[1]]++
		}
	}

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
