package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func writeCSV(header []string, data [][]string) (dataBytes *bytes.Buffer)  {
	dataBytes = &bytes.Buffer{}
	// 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码
	//dataBytes.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(dataBytes)

	_ = writer.Write(header)

	for _, row := range data {
		_ = writer.Write(row)
	}

	writer.Flush()

	return
}

func main()  {
	h := []string{"姓名", "年龄"}
	d := [][]string{
		{"张三", "20"},
		{"李四", "21"},
	}
	res := writeCSV(h, d)
	fmt.Println(res.String())
}
