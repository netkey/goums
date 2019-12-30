package serial

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"

	"github.com/tsingson/goums/apis/flatums"
)

// ReadList read terminal sn list from excel file
func ReadList(excelFileName string) (v *flatums.TerminalListT, err error) {
	xlFile, er1 := xlsx.OpenFile(excelFileName)
	if er1 != nil {
		fmt.Println("文件错误啦")
		return
	}

	count := 0
	list := make([]*flatums.TerminalProfileT, 0)

	//
	for _, sheet := range xlFile.Sheets {
		for n, row := range sheet.Rows {
			if n == 0 {
				continue
			}
			var serial, code, role string

			for i, cell := range row.Cells {
				// i == 0 是序号, 跳过
				if i == 1 {
					serial = cell.String()
				}

				if i == 2 {
					code = cell.String()
				}
				if i == 3 {
					role = cell.String()
				}

			}
			serial = strings.TrimSpace(serial)
			if len(serial) < 8 {
				continue
			}
			a := &flatums.TerminalProfileT{
				SerialNumber: serial,
				ActiveCode:   code,
				AccessRole:   role,
				Operation:    flatums.NotifyTypeinsert,
			}
			list = append(list, a)
			count++
		}
	}

	v = &flatums.TerminalListT{
		Count: int64(count),
		List:  list,
	}

	return
}
