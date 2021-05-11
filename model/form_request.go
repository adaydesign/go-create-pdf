package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

type FormRequest struct {
	ID                    int
	FormLocal             string
	FormDate              *time.Time
	RequesterFullName     string
	RequesterPositionName string
	LevelID               uint
	RequesterLevelName    string // !
	OfficeCode            string
	OfficeName            string // !
	Tel                   string
	MaritalStatusID       uint
	SpouseFullName        string
	SpousePositionName    string
	SpouseOfficeName      string
	Children              int
	Address               string
	Soi                   string
	Subdistrict           string
	District              string
	Province              string
	HouseOwnerID          uint
	HouseStatusID         uint
	HouseOtherStatusName  string
	ResidentID            uint
	HeadFullName          string
	HeadPositionName      string
	ScanDocuments         []int
	ScanDocumentOtherName string
}

func (frm FormRequest) GeneratePDFDocument() error {

	// current date
	monthTHList := "มกราคม,กุมภาพันธ์,มีนาคม,เมษายน,พฤษภาคม,มิถุนายน,กรกฎาคม,สิงหาคม,กันยายน,ตุลาคม,พฤศจิกายน,ธันวาคม"
	monthTH := strings.Split(monthTHList, ",")
	year, month, day := time.Now().Date()

	var err error
	// PDF setting
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("THSarabunIT9", "", "/pdf-template/THSarabunIT9.ttf")
	pdf.SetFont("THSarabunIT9", "", 16)

	// Import example-pdf.pdf with gofpdi free pdf document importer
	// fix local of form1.pdf -- tempate
	tpl1 := gofpdi.ImportPage(pdf, "./pdf-template/form1.pdf", 1, "/MediaBox")

	// page 1
	pdf.AddPage()
	gofpdi.UseImportedTemplate(pdf, tpl1, 0, 0, 210, 0)

	// เขียนที่
	pdf.Text(128, 29, frm.FormLocal)
	// วันที่
	pdf.Text(110, 37.5, fmt.Sprintf("%d", day))
	pdf.Text(131, 37.5, monthTH[int(month)])
	pdf.Text(175, 37.5, fmt.Sprintf("%d", year+543))
	// ข้าพเจ้า
	pdf.Text(60, 57.5, frm.RequesterFullName)
	pdf.Text(132, 57.5, frm.RequesterPositionName)
	// ระดับ
	pdf.Text(35, 65, frm.RequesterLevelName)
	pdf.Text(69, 65, frm.OfficeName)
	pdf.Text(150, 65, frm.Tel)
	// สถานะภาพ
	if frm.MaritalStatusID == 1 {
		pdf.Text(51.5, 71.5, "/")
	} else if frm.MaritalStatusID == 2 {
		pdf.Text(73, 71.5, "/")
	} else if frm.MaritalStatusID == 3 {
		pdf.Text(118, 71.5, "/")
	}
	// คู่สมรส
	pdf.Text(43, 78, frm.SpouseFullName)
	pdf.Text(125, 78, frm.SpousePositionName)
	// สถานที่ปฏิบัตังาน
	pdf.Text(55, 85, frm.SpouseOfficeName)
	pdf.Text(178, 85, fmt.Sprintf("%d", frm.Children))
	// ปัจจุบัน
	pdf.Text(84, 92, frm.Address)
	pdf.Text(125, 92, frm.Soi)
	// แขวง
	pdf.Text(46, 99, frm.Subdistrict)
	pdf.Text(105, 99, frm.District)
	pdf.Text(154, 99, frm.Province)
	// เป็นของ
	if frm.HouseOwnerID == 1 {
		pdf.Text(65, 105.5, "/")
	} else if frm.HouseOwnerID == 2 {
		pdf.Text(86.5, 105.5, "/")
	} else if frm.HouseOwnerID == 3 {
		pdf.Text(107, 105.5, "/")
	}
	// โดยเป็น
	if frm.HouseStatusID == 1 {
		pdf.Text(139.5, 105.5, "/")
	} else if frm.HouseStatusID == 2 {
		pdf.Text(139.5, 112, "/")
	} else if frm.HouseStatusID == 3 {
		pdf.Text(139.5, 118.5, "/")
	} else if frm.HouseStatusID == 4 {
		pdf.Text(139.5, 125.5, "/")
		pdf.Text(162, 125.5, frm.HouseOtherStatusName)
	}
	// อาคาร
	if frm.ResidentID == 1 {
		pdf.Text(65.5, 141, "/")
	} else if frm.ResidentID == 2 {
		pdf.Text(65.5, 148, "/")
	} else if frm.ResidentID == 3 {
		pdf.Text(65.5, 155, "/")
	} else if frm.ResidentID == 4 {
		pdf.Text(65.5, 161.5, "/")
	}
	// เอกสาร
	for i := 0; i < len(frm.ScanDocuments); i++ {
		docID := frm.ScanDocuments[i]
		if docID == 1 {
			pdf.Text(65.5, 175, "/")
		}
		if docID == 2 {
			pdf.Text(65.5, 182, "/")
		}
		if docID == 3 {
			pdf.Text(65.5, 189, "/")
		}
		if docID == 4 {
			pdf.Text(65.5, 196, "/")
		}
		if docID == 5 {
			pdf.Text(65.5, 202.5, "/")
		}
		if docID == 6 {
			pdf.Text(65.5, 215.5, "/")
			pdf.Text(90, 215.5, frm.ScanDocumentOtherName)
		}
	}

	// ลงชื่อ
	pdf.Text(122, 267, frm.RequesterFullName)
	pdf.Text(120, 274, frm.RequesterPositionName)

	//page 2
	pdf.AddPage()
	tpl2 := gofpdi.ImportPage(pdf, "./pdf-template/form1.pdf", 2, "/MediaBox")
	gofpdi.UseImportedTemplate(pdf, tpl2, 0, 0, 210, 0)

	// ข้าพเจ้า
	pdf.Text(64, 45, frm.HeadFullName)
	pdf.Text(139, 45, frm.HeadPositionName)
	// ของ
	pdf.Text(78, 53, frm.RequesterFullName)
	// ลงชื่อ
	pdf.Text(123, 112.5, frm.HeadFullName)
	pdf.Text(121, 120, frm.HeadPositionName)

	// file output : YYYYMMDD_ID.pdf
	fileName := fmt.Sprintf("./pdf-output/%d%02d%02d_%05d.pdf", year, int(month), day, frm.ID)

	err = pdf.OutputFileAndClose(fileName)
	if err != nil {
		return err
	}

	return nil
}
