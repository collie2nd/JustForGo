package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func main() {
	f := excelize.NewFile()
	// 修改默认的 Sheet1 命名
	f.SetSheetName("Sheet1", "Test_Origin")

	// Create a new sheet.
	f.NewSheet("Test")

	// 填充单位格
	for i := 1; i <= 10; i++ {
		f.SetCellValue("Test", "A"+strconv.Itoa(i), strconv.Itoa(i*100))
	}

	// 设置列宽，可一次设置等宽多列
	f.SetColWidth("Test", "A", "H", 25)

	// 设置行高
	f.SetRowHeight("Test", 1, 100)

	// 设置列表大纲模式
	for i := 3; i <= 8; i++ {
		f.SetRowOutlineLevel("Test", i, 1)
	}

	// 设置下拉选项
	// 新建下拉数据验证数据结构，可以选空 allowBlank:true
	dvRangeAgent := excelize.NewDataValidation(true)
	// 此处设置需要添加数据验证的行
	dvRangeAgent.Sqref = "D2:F10"
	// 此处设置数据验证选项源
	dvRangeAgent.SetSqrefDropList("=Test!$A$1:$A$10")
	// 添加到对应Sheet
	f.AddDataValidation("Test", dvRangeAgent)
	f.AddDataValidation("Test_Origin", dvRangeAgent)

	// 设置单元格样式
	// 样式除了颜色填充以外还有边框，字体等等
	style, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#FF0000"}, Pattern: 1},
	})
	f.SetCellStyle("Test_Origin", "A1", "H10", style)

	// Save spreadsheet by the given path.
	if err := f.SaveAs("./test.xlsx"); err != nil {
		fmt.Println(err)
		return
	}

	return
}
