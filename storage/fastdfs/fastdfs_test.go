// Author: yangzq80@gmail.com
// Date: 2020-06-05
//
package fastdfs

import (
	"fmt"
	fastdfs "github.com/wodog/fastdfs-client"
	"go-samples/file"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

var root = "/Users/zqy/test/batch-files"
var trackerAddr = "192.168.251.157:22122"

func TestFastdfsDownload(t *testing.T) {

	client := fastdfs.Default()
	client.AddTracker(trackerAddr)

	fileId := "group2/M00/00/00/wKj7nV7Y5yWAPZ9pAAAAD3y2II81726415"

	// download file
	reader, err := client.Open(fileId)

	if err != nil {
		log.Println(err)
		return
	}

	file.WriteFile("remote-dat", reader)
}

func TestFastdUpload(t *testing.T) {

	client := fastdfs.Default()
	client.AddTracker(trackerAddr)

	// upload file
	file, _ := os.Open(root + "/t1-output.txt")

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	fileId, _ := client.Upload(file)

	fmt.Println(fileId)

}

//测试fastdfs性能
func TestFastdfsPerformance(t *testing.T) {

	client := fastdfs.Default()
	client.AddTracker(trackerAddr)

	start := time.Now()

	//create files
	initBatchFiles(5000)

	timeTrack(start, "Init files")
	start = time.Now()

	//upload files
	file.HandleDirFiles(root, func(f *os.File) {
		fileId, _ := client.Upload(f)
		file.AppendLine(root+"/list.dat", fileId)
	})

	timeTrack(start, "Upload files")
	start = time.Now()

	//delete files
	file.HandleLines(root+"/list.dat", func(line string) {
		err := client.Delete(line)
		if err != nil {
			log.Println("delete error", err, line)
		}
	})
	timeTrack(start, "Delete files")

}

func initBatchFiles(num int) {
	for i := 0; i < num; i++ {
		file.MockTextFiles(root+"/mock-file-"+strconv.Itoa(i), 100)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
