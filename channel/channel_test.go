// Author: yangzq80@gmail.com
// Date: 2020-09-24
//
package channel

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	c := make(chan int)
	defer close(c)
	go func() { c <- 3 + 4 }()
	i := <-c
	fmt.Println(i)
}
