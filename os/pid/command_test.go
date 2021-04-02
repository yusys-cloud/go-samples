// Author: yangzq80@gmail.com
// Date: 2020-10-22
//
package pid

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestJava(t *testing.T) {
	//cmd := exec.Command("java -jar /Users/zqy/soft/zipkin-server-2.15.0-exec.jar")
	cmd := exec.Command("./script.sh")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)

	time.Sleep(5 * 1000)
}
