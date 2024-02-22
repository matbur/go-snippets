package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	bb, err := os.ReadFile("../samples/invoice.pdf")
	if err != nil {
		panic(err)
	}

	r, err := model.NewPdfReader(bytes.NewReader(bb))
	if err != nil {
		panic(err)
	}

	nPages, err := r.GetNumPages()
	if err != nil {
		panic(err)
	}

	var b strings.Builder
	for i := 1; i <= nPages; i++ {
		page, err := r.GetPage(i)
		if err != nil {
			panic(err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			panic(err)
		}

		text, err := ex.ExtractText()
		if err != nil {
			panic(err)
		}

		b.WriteString(text)
	}
	fmt.Println(b.String())
}
