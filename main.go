package main

import (
	"fmt"
	"go-pdf-1/model"
	"log"

	"github.com/gofiber/fiber/v2"
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

	// run normal
	officer1.GeneratePDFDocument()
	fmt.Print("generate finished -- officer1 (normal)")

	// run server
	app := fiber.New()
	app.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "1" {
			officer1.GeneratePDFDocument()
			return c.SendString("generate finished -- officer1")
		} else if id == "2" {
			officer2.GeneratePDFDocument()
			return c.SendString("generate finished -- officer2")
		} else {
			return c.SendString("Hello, World!")
		}
	})

	log.Fatal(app.Listen(":3030"))
}
