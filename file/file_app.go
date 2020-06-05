// Author: yangzq80@gmail.com
// Date: 2020-06-05
//
package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func AppendLine(name string, line string) {

	f, err := os.OpenFile(name,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(line + "\n"); err != nil {
		log.Println(err)
	}
}

//模拟生成随机行数的文本文件
func MockTextFiles(name string, randLines int) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < rand.Intn(randLines); i++ {
		f.WriteString(strconv.Itoa(i) + " File handling is easy.\n")
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type HandleLine func(line string)

//读取文件内容按行进行处理
func HandleLines(name string, callback HandleLine) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type HandleFile func(file *os.File)

func HandleDirFiles(dirname string, callback HandleFile) {
	files, _ := ioutil.ReadDir(dirname)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			f, _ := os.Open(dirname + "/" + file.Name())
			defer f.Close()
			callback(f)
		}
	}
}

func WriteFile(name string, reader io.Reader) {
	// open output file
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := file.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}
