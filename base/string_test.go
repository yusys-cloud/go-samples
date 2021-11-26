// Author: yangzq80@gmail.com
// Date: 2021-09-15
//
package base

import (
	"fmt"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	fmt.Println(strings.Contains(strings.ToLower("UP"), "up"))

	var a int
	a = 10
	a = int(float64(a) / 30 * 100)
	fmt.Println(a)
}
