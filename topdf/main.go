package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	file := strings.Join(os.Args[1:], "")

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	// OrientationStr, UnitStr, SizeStr, FontDirStr
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	//FamilyStr, StyleStr, Size
	pdf.SetFont("Arial", "B", 14)

	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("output.pdf")

	fmt.Println("PDF Creato")

}
