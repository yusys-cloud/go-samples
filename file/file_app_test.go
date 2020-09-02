// Author: yangzq80@gmail.com
// Date: 2020-06-05
//
package file

import (
	"fmt"
	"io/ioutil"
	"log"
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

func TestFilenames(t *testing.T) {

	//var files []string
	//
	//root := "/Users/zqy/Downloads/tmp/najie"
	//err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	//	fmt.Println(path)
	//	files = append(files, path)
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}
	//for _, file := range files {
	//	fmt.Println(file)
	//}

	files, err := ioutil.ReadDir("/Users/zqy/Downloads/tmp/najie/yz1")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
