package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	bb, err := os.ReadFile("../samples/invoice.pdf")
	if err != nil {
		panic(err)
	}

	r, err := pdf.NewReader(bytes.NewReader(bb), int64(len(bb)))
	if err != nil {
		panic(err)
	}

	var b strings.Builder
	for i := 1; ; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			break
		}

		content := page.Content()
		for _, text := range content.Text {
			b.WriteString(text.S)
		}
		b.WriteByte('\n')
	}
	fmt.Println(b.String())
}
