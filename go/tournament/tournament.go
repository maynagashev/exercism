package tournament

import (
	"bufio"
	"fmt"
	"io"
)

// Tally writes to specified writer the number of points, goals, runs, etc. achieved in a game or by a team.
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Println(t)
		w.Write([]byte(t))
	}
	return nil
}
