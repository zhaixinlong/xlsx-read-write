package xlsx

import (
	"github.com/emirpasic/gods/sets/linkedhashset"
	"github.com/xuri/excelize/v2"
)

func ReadExcelToSet(filename string, set *linkedhashset.Set) (*linkedhashset.Set, error) {
	f, err := excelize.OpenFile(filename) //
	if err != nil {
		return nil, err
	}
	firstSheet := f.GetSheetName(0)
	rows, err := f.GetRows(firstSheet)

	for i, row := range rows {
		if i == 0 {
			continue
		}
		for _, colCell := range row {
			set.Add(colCell)
			break
		}
	}
	return set, err
}
