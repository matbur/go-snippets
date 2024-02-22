package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	bb, err := os.ReadFile("../samples/invoice.pdf")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("pdftotext", "-layout", "-nopgbrk", "-enc", "UTF-8", "-eol", "unix", "-", "-")
	cmd.Stdin = bytes.NewReader(bb)
	body, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
