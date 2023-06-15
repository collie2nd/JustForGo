package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("D:/GolandProjects/JustForGo/excel/read_excel/HealthCheck.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println("Excel 文件关闭失败: " + err.Error())
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("监测点列表")
	if err != nil || len(rows) == 0 {
		fmt.Println("Excel 文件读取失败: " + err.Error())
		return
	}

	for _, row := range rows[1:] {
		// 在表格末尾有空格的时候，每一行实际的 row 并不等长
		// 但是行首有空格并不会导致 row 长度变短
		rowTmp := make([]string, 6)
		for k, v := range row {
			rowTmp[k] = v
		}
		fmt.Printf("%#v\n", row)
		fmt.Printf("%#v\n", rowTmp)
	}

	fmt.Println("----------------------------")

	r, _ := f.Rows("监测点列表")
	for {
		if r.Next() {
			rst, _ := r.Columns()
			fmt.Printf("%#v\n", rst)
		} else {
			break
		}
	}

	fmt.Println("----------------------------")

	// 检查行的存在性
	visible1, _ := f.GetRowVisible("Sheet1", 2)
	visible2, _ := f.GetRowVisible("代理列表", 2)
	visible3, _ := f.GetRowVisible("代理列表", 10)
	fmt.Println(visible1, visible2, visible3)

	fmt.Println("----------------------------")
}
