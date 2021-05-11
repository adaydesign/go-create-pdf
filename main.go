package main

import (
	"fmt"
	"go-pdf-1/model"
)

func main() {
	// mockup data
	officer1 := &model.FormRequest{
		ID:                    1,
		FormLocal:             "สำนักเทคโนโลยีสารสนเทศ",
		RequesterFullName:     "นายทดสอบ แบบคำร้องทดสอบ",
		RequesterPositionName: "นักวิชาการทั่วไป",
		RequesterLevelName:    "ปฏิบัติการ",
		OfficeName:            "สำนักเทคโนโลยีสารสนเทศ",
		Tel:                   "089 999 9999",
		MaritalStatusID:       1,
		SpouseFullName:        "-",
		SpousePositionName:    "-",
		SpouseOfficeName:      "-",
		Children:              0,
		Address:               "111/12",
		Soi:                   "ซอย 11 ถนนกว้าง",
		Subdistrict:           "จตุจักร",
		District:              "ลาดพร้าว",
		Province:              "กรุงเทพมหานคร",
		HouseOwnerID:          0,
		HouseStatusID:         1,
		ResidentID:            3,
		HeadFullName:          "นายหัวหน้า อยู่เย็น",
		HeadPositionName:      "นักวิชาการชำนาญการพิเศษ",
		ScanDocuments:         []int{1, 2, 3},
	}

	officer1.GeneratePDFDocument()
	fmt.Println("generate finished -- officer1")

	officer2 := &model.FormRequest{
		ID:                    2,
		FormLocal:             "สำนักเทคโนโลยีสารสนเทศ",
		RequesterFullName:     "นายทดสอบ แบบคำร้องทดสอบ",
		RequesterPositionName: "นักวิชาการทั่วไป",
		RequesterLevelName:    "ปฏิบัติการ",
		OfficeName:            "สำนักเทคโนโลยีสารสนเทศ",
		Tel:                   "089 999 9999",
		MaritalStatusID:       2,
		SpouseFullName:        "นางสาวคู่สมรส แบบคำร้องทดสอบ",
		SpousePositionName:    "นักข่าว",
		SpouseOfficeName:      "สำนักงานข่าวแห่งชาติ",
		Children:              2,
		Address:               "111/12",
		Soi:                   "ซอย 11 ถนนกว้าง",
		Subdistrict:           "จตุจักร",
		District:              "ลาดพร้าว",
		Province:              "กรุงเทพมหานคร",
		HouseOwnerID:          2,
		HouseStatusID:         4,
		HouseOtherStatusName:  "บ้านแบ่งเช่า",
		ResidentID:            4,
		HeadFullName:          "นายหัวหน้า อยู่เย็น",
		HeadPositionName:      "นักวิชาการชำนาญการพิเศษ",
		ScanDocuments:         []int{1, 2, 3, 4, 5, 6},
		ScanDocumentOtherName: "เอกสารความลับ",
	}

	officer2.GeneratePDFDocument()
	fmt.Println("generate finished -- officer2")

}
