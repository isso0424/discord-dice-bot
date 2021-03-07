package tester

import (
	"bufio"
	"os"
)

func receive() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	return stdin.Text()
}
