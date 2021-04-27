// Author: yangzq80@gmail.com
// Date: 2021-04-26
//
package file

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestCsv(t *testing.T) {
	csvFile1, _ := os.Open("/Users/zqy/codes/python/tmp/cx-1.csv")
	defer csvFile1.Close()
	csvFile2, _ := os.Open("/Users/zqy/codes/python/tmp/cx-2.csv")

	outFile, _ := os.OpenFile("/Users/zqy/codes/python/tmp/cx-out.csv", os.O_CREATE|os.O_RDWR, 0644)
	wCsv := csv.NewWriter(outFile)

	csv1Lines, _ := csv.NewReader(csvFile1).ReadAll()
	csv2Lines, _ := csv.NewReader(csvFile2).ReadAll()

	for _, l1 := range csv1Lines {
		tl := []string{l1[0], ""}
		for _, l2 := range csv2Lines {
			if l1[0] == l2[0] {
				tl[1] = l2[1]
				break
			}
		}
		wCsv.Write(tl)
	}

	wCsv.Flush()
}
