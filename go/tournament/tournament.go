package tournament

import (
	"fmt"
	"io"
)

func Tally(r io.Reader, w io.Writer) error {
	p := make([]byte, 100)
	fmt.Println(r.Read(p))
	fmt.Println(string(p))
	fmt.Println(r.Read(p))
	fmt.Println(string(p))
	return nil
}
