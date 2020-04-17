package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrenter_Run(t *testing.T) {

		arr := make([]interface{}, 0)
		for i:=0; i<100;i++ {
			arr = append(arr, i)
		}
	Conncurrent.Run(arr, 2, func(data interface{}) {
			fmt.Println("data===", data, "time===", time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(2*time.Second)
		})


}