// Author: yangzq80@gmail.com
// Date: 2020-11-23
//
package os_exec

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	//cmd := exec.Command("/bin/bash", "-c", "docker exec -i redis redis-cli --cluster check 172.16.20.223:7002")
	cmd := exec.Command("/bin/bash", "-c", "df -lh")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func TestStdinPipe(t *testing.T) {
	cmd := exec.Command("/bin/bash", "-c", "docker exec -i redis redis-cli --cluster check 172.16.20.223:7002")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range strings.Split(string(out), "\n") {
		fmt.Println(l)
	}
	//fmt.Printf("%s\n", out)
}
