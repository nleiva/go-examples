// https://play.golang.org/p/QNWdEIfLsXU

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n])) //should handle any remainding bytes.
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}
