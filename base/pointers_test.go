// Author: yangzq80@gmail.com
// Date: 2021-08-31
//
package base

import (
	"fmt"
	"testing"
)

func TestPointers(t *testing.T) {
	var a int = 20 /* 声明实际变量 */
	fmt.Println(" a = ", a)
	var ip *int /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}
