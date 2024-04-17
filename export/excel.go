package export

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gofreelance/internals/index"
	"gofreelance/model"
	"os"
	"strings"
)

//go:embed clitt-templ.xlsx
var template []byte

func SaveAsExcel(r *model.RecordDataSet) error {
	f, _ := os.Create("test.xlsx")
	excelF, err := excelize.OpenReader(bytes.NewBuffer(template))
	if err != nil {
		return fmt.Errorf("failed to open excel template: %s", err.Error())
	}
	sheetName := index.Current()
	sheetID, _ := excelF.GetSheetIndex(sheetName)
	// sheet does not exist, let's create a new one
	if sheetID == -1 {
		sheetID, err = excelF.NewSheet(sheetName)
		if err != nil {
			return fmt.Errorf("failed to set sheet name: %s", err.Error())
		}
		// copy the template to the new sheet
		if err := excelF.CopySheet(0, sheetID); err != nil {
			return fmt.Errorf("failed to copy sheet: %s", err.Error())
		}
	}
	excelF.SetActiveSheet(sheetID)

	workTime := 0.0
	currentEntries := r.Entries[sheetName]
	for _, e := range currentEntries {
		offset := 11
		recordTime := 0.0
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("A%d", offset), fmt.Sprintf("%s", e.Started.Format("2006-01-02")))
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("C%d", offset), getTaskEntries(e.Tasks))
		if e.Ended != nil {
			minutes := e.Ended.Sub(*e.Started).Minutes()
			recordTime += minutes
		}
		var taskList string
		for _, t := range e.Tasks {
			taskList += fmt.Sprintf("%s - ", t.Text)
		}
		_ = excelF.SetCellValue(sheetName, fmt.Sprintf("D%d", offset), recordTime/60)
		offset++
		workTime += recordTime
	}
	_ = excelF.SetCellValue(sheetName, fmt.Sprintf("B8"), workTime/60)

	if _, err := excelF.WriteTo(f); err != nil {
		return fmt.Errorf("failed to write excel: %s", err)
	}
	return f.Close()
}

func getTaskEntries(r []model.Task) string {
	var descriptions []string
	for _, d := range r {
		descriptions = append(descriptions, d.Text)
	}

	return strings.Join(descriptions, " / ")
}
