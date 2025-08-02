package utils

import (
	"fmt"
	"kumande/models"

	"github.com/jung-kurt/gofpdf"
)

func GetStringNoData(s *string) string {
	if s != nil {
		return *s
	}
	return "-"
}

func GeneratePDFErrorAudit(c []models.ErrorAudit, filename string) error {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetTitle("Kumande", false)
	pdf.AddPage()

	// Set Header
	pdf.SetFont("Arial", "B", 20)
	pdf.SetTextColor(0, 102, 204)
	pdf.CellFormat(0, 12, "Kumande", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "I", 12)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 10, "Daily Consume Management", "", 1, "C", false, 0, "")
	pdf.Ln(4)

	// Set Letterhead
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Audit - Error")
	pdf.Ln(8)

	// Set header
	pdf.SetFillColor(200, 200, 200)
	pdf.CellFormat(175, 10, "Error Message", "1", 0, "C", true, 0, "")
	pdf.CellFormat(60, 10, "Datetime", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 10, "Total", "1", 1, "C", true, 0, "")

	// Set body
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 255)
	for _, dt := range c {
		pdf.CellFormat(175, 8, dt.Message, "1", 0, "L", false, 0, "")
		pdf.CellFormat(60, 8, dt.CreatedAt, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 8, fmt.Sprintf("%d", dt.Total), "1", 1, "C", false, 0, "")
	}

	return pdf.OutputFileAndClose(filename)
}
