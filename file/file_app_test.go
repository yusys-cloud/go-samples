// Author: yangzq80@gmail.com
// Date: 2020-06-05
//
package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var root = "/Users/zqy/test"

func TestAppendLine(t *testing.T) {
	AppendLine(root+"/t1.txt", "appends contents.")
}

func TestHandle_lines(t *testing.T) {
	HandleLines(root+"/t1.txt", func(line string) {
		fmt.Println(line)
	})
}

func TestWrite_file(t *testing.T) {
	WriteFile(root+"/t2.dat", strings.NewReader("reader contents."))
}

func TestBatchFile(t *testing.T) {

	for i := 0; i < 10; i++ {
		MockTextFiles(root+"/batch-files/mock-file-"+strconv.Itoa(i), 100)
	}
}

func TestFiles(t *testing.T) {
	HandleDirFiles(root, func(file *os.File) {
		fmt.Println(file.Name())
	})
}
