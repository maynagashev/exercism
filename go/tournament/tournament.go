package tournament

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Tally writes to specified writer the number of points, goals, runs, etc. achieved in a game or by a team.
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for scanner.Scan() {
		row := scanner.Text()
		f := strings.Split(row, ";")
		fmt.Println(row)
		fmt.Sprintf("%+v\n", f)
		w.Write([]byte(row+"\n"))
	}
	return nil
}
