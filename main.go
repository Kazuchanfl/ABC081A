package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var input int

	fmt.Fscan(r, &input)

	var castedInput string
	castedInput = strconv.Itoa(input)

	fmt.Fprintln(w, castedInput)

	//fmt.Println("Hello, World!")
}
