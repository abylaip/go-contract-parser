package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/ledongthuc/pdf"
)

type contract struct {
	header string
	first string
	second string
	third string
	fourth string
	fifth string
	conditions string
	client_signature string
	client_date string
	contractor_signature string
	contractor_date string
}

func main() {
	deleteHello()
	
	pdf.DebugOn = true
	content, err := readPdf("Contract.pdf")
	if err != nil {
		panic(err)
	}

	fmt.Println(content)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	addHeader(pdf, content, 16)
	pdf.Ln(-1)

	erro := pdf.OutputFileAndClose("hello.pdf")

	if erro != nil {
		panic(erro)
	}

	return
}

func addHeader(pdf *gofpdf.Fpdf, text string, size int) {
	if pdf == nil {
		return
	}
	pdf.SetFillColor(255, 255, 255)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", (float64)(size))
	pdf.CellFormat((float64)(len(text)*5), (float64)(size/2), text, "0", 0, "L", true, 0, "")
	pdf.Ln(-1)
}


func takePart(from string, to string, content string) string {
	var result bytes.Buffer
	i := strings.Index(content, from)
	j := strings.Index(content, to)

	for k := i + len(from); k < j; k++ {
		result.WriteString(string(content[k]))
	}

	return result.String()
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
    defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
    b, err := r.GetPlainText()
    if err != nil {
        return "", err
    }
    buf.ReadFrom(b)

	return buf.String(), nil
}

func deleteHello() {
	e := os.Remove("hello.pdf")
    if e != nil {
        panic(e)
    }
}