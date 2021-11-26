// Author: yangzq80@gmail.com
// Date: 2021-10-29
//
package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total/1024/1024/1024, v.Free/1024/1024/1024, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
