// Author: yangzq80@gmail.com
// Date: 2021-03-30
//
package json

import (
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestJsonToMaps(t *testing.T) {

	var jsonStr = `
{
  "chaos": [
    {
      "node": { "ip": "192.168.1.1", "name": "主-Redis-M1" },
      "blades": [
        {
          "cmd": "blade create jvm throwCustomException",
          "desc": "指定类方法抛自定义异常，命令可以简写为 blade c jvm tce",
          "id": "blade:1377556306784686080",
          "name": "java-方法-异常",
          "params": [
            {
              "desc": "指定 java 进程名，如果同时填写pid优先",
              "id": "--process",
              "label": "推荐填写该值",
              "name": "java 进程名",
              "value": "330"
            },
            {
              "desc": "限制影响百分比",
              "id": "--effect-percent",
              "name": "影响百分比"
            }
          ],
          "type": "blade"
        },
        {
          "cmd": "blade create jvm OutOfMemoryError",
          "desc": "内存溢出场景，命令可以简写为：blade c jvm oom",
          "id": "blade:1377556719797800960",
          "name": "java-内存-溢出",
          "params": [
            {
              "desc": "设定运行时长，单位是秒，通用参数",
              "id": " --timeout",
              "label": "3000",
              "name": "运行时长"
            },
            {
              "desc": "开启DEBUG日志打印，当故障未生效，可以通过DEBUG日志来观察执行过程和相关信息。\n（公共日志目录：/opt/chaosblade/logs/chaosblade.log，\nJAVA相关场景的日志目录：<USER_HOME>/logs/chaosblade/chaosblade.log）",
              "id": "--debug",
              "label": "true或者false",
              "name": "开启DEBUG"
            }
          ],
          "type": "blade"
        }
      ]
    },
    {
      "node": { "ip": "192.168.2.2:8080", "name": "主-ZK-F1" },
      "blades": [
	  {
		"name": "线程等待",
		"cmd": "thread-wait",
		"params": [
		  {
			"name": "等待时间(秒)",
			"lable": "60",
			"value": "5"
		  }
		]
	  },
        {
          "cmd": "blade create jvm OutOfMemoryError",
          "desc": "内存溢出场景，命令可以简写为：blade c jvm oom",
          "id": "blade:1377556719797800960",
          "name": "java-内存-溢出",
          "params": [
            {
              "desc": "设定运行时长，单位是秒，通用参数",
              "id": " --timeout",
              "label": "3000",
              "name": "运行时长"
            },
            {
              "desc": "开启DEBUG日志打印，当故障未生效，可以通过DEBUG日志来观察执行过程和相关信息。\n（公共日志目录：/opt/chaosblade/logs/chaosblade.log，\nJAVA相关场景的日志目录：<USER_HOME>/logs/chaosblade/chaosblade.log）",
              "id": "--debug",
              "label": "true或者false",
              "name": "开启DEBUG"
            }
          ],
          "type": "blade"
        }
      ]
    }
  ],
  "name": "测试",
  "desc": "测试",
  "status": 0
}

`
	j, _ := jsonvalue.Unmarshal([]byte(jsonStr))

	chaos, _ := j.GetArray("chaos")

	chaos.RangeArray(func(i int, v *jsonvalue.V) bool {
		node, _ := v.Get("node")
		ip, _ := node.GetString("ip")

		blades, _ := v.GetArray("blades")

		blades.RangeArray(func(i int, v *jsonvalue.V) bool {
			cmd, _ := v.GetString("cmd")

			//拼命令
			req := cmd
			logrus.Info(ip, "---", i, "---", cmd)

			if cmd == "thread-wait" {
				p, _ := v.GetArray("params")
				p.RangeArray(func(i int, v *jsonvalue.V) bool {
					t, _ := v.GetString("value")
					it, _ := strconv.Atoi(t)
					time.Sleep(time.Duration(it) * time.Second)
					return false
				})
				return true
			}

			bladeParams, _ := v.GetArray("params")

			//拼命令[参数]
			bladeParams.RangeArray(func(i int, v *jsonvalue.V) bool {

				id, _ := v.GetString("id")
				idv, _ := v.GetString("value")

				if idv != "" {

					req = fmt.Sprintf("%s %s %s", req, id, idv)

				}

				return true
			})

			req = url.PathEscape(req)
			//默认拼接:6666端口
			var freq string
			if !strings.Contains(ip, ":") {
				freq = fmt.Sprintf("http://%s:6666/chaosblade?cmd=%s", ip, req)
			} else {
				freq = fmt.Sprintf("http://%s/chaosblade?cmd=%s", ip, req)
			}
			go func() {
			}()
			logrus.Println("send...", freq)
			rs := jsonvalue.NewObject()
			rs.SetString(freq).At("reqUrl")
			//rs.SetString(getUrl(freq)).At("respBody")

			v.Set(rs).At("runResult")

			logrus.Println(chaos)

			return true
		})

		return true
	})

}

func getUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err.Error())
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}

func TestGetUrl(t *testing.T) {

	fmt.Println(getUrl("http://localhost:9999/api/kv/chaos/designer/designer:1377565223489114112"))

}
