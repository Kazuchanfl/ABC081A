package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var input string

	fmt.Fscan(r, &input)

	fmt.Fprintln(w, strings.Count(input, "1"))
}
