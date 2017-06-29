package main

import (
	"encoding/json"
	"fmt"

	"github.com/ledongthuc/linkedin_pdf_cv_parser/parser"
)

func main() {
	resumeProfile, err := parser.ParsePDFContent("example/my_cv.pdf")
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(resumeProfile, "", " ")
	fmt.Println(string(b))
}
