package export

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gofreelance/model"
	"os"
	"strings"
)

//go:embed template.xlsx
var template []byte

func SaveAsExcel(r *model.RecordDataSet) error {
	f, _ := os.Create("test.xlsx")
	excelF, err := excelize.OpenReader(bytes.NewBuffer(template))
	if err != nil {
		return fmt.Errorf("failed to open excel template: %s", err.Error())
	}
	for k, entries := range r.Entries {
		id, err := excelF.NewSheet(k)
		if err := excelF.CopySheet(0, id); err != nil {
			return err
		}
		if err != nil {
			return err
		}
		excelF.DeleteSheet("Sheet1")
		err = excelF.SetColWidth(k, "A", "D", 30)
		if err != nil {
			return fmt.Errorf("failed to set col width: %s", err.Error())
		}
		workTime := 0.0
		_ = excelF.SetColStyle(k, "A", id)
		_ = excelF.SetCellValue(k, "A1", "Date")
		_ = excelF.SetCellValue(k, "B1", "Description")
		_ = excelF.SetCellValue(k, "C1", "Duration (Minutes)")
		_ = excelF.SetCellValue(k, "D", "Total (Hours)")

		for idx, e := range entries {
			_ = excelF.SetCellValue(k, fmt.Sprintf("A%d", idx+2), e.Started)
			_ = excelF.SetCellValue(k, fmt.Sprintf("B%d", idx+2), getDescriptions(e.Descriptions))
			if e.Ended != nil {
				minutes := e.Ended.Sub(*e.Started).Minutes()
				workTime += minutes
				_ = excelF.SetCellValue(k, fmt.Sprintf("C%d", idx+2), minutes)
			}
			_ = excelF.SetCellValue(k, fmt.Sprintf("D%d", idx+2), workTime/60)
		}

	}

	if _, err := excelF.WriteTo(f); err != nil {
		return fmt.Errorf("failed to write excel: %s", err)
	}
	return f.Close()
}

func getDescriptions(r []model.RecordDescriptionEntry) string {
	var descriptions []string
	for _, d := range r {
		descriptions = append(descriptions, d.Text)
	}

	return strings.Join(descriptions, " / ")
}
