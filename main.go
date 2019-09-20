package main

import (
	"fmt"
	"io/ioutil"

	"github.com/quiqueporta/b17-crew-generator/pdf"
)

func main() {
	ioutil.WriteFile("B-17_crew.pdf", pdf.GeneratePDF(), 0644)
	fmt.Println("B-17_crew.pdf was created.")
}
