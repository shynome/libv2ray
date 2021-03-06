package libv2ray

import (
	"strings"
	"time"

	core "github.com/v2fly/v2ray-core/v4"

	// include modules
	_ "github.com/shynome/libv2ray/distro/all"
)

var server *core.Instance
var startCount int = 0

func increaseStartCount() {
	startCount++
	if startCount > 5_000_000 {
		startCount = 0
	}
}

// Start v2ray Sevice
func Start(jsonConfig string) (err error) {
	if server != nil {
		return
	}
	config, err := core.LoadConfig("json", "stdin:", strings.NewReader(jsonConfig))
	if err != nil {
		return
	}
	server, err = core.New(config)
	if err != nil {
		return
	}
	c := make(chan error)
	increaseStartCount()
	go func() {
		c <- server.Start()
		server = nil
	}()
	go func() {
		// 如果 1s 后 Start 未退出则认为启动成功
		time.Sleep(time.Second)
		c <- nil
	}()
	err = <-c
	return
}

// Stop v2ray Sevice
func Stop() {
	if server == nil {
		return
	}
	s := server
	server = nil
	s.Close()
}

func Status() int {
	if server == nil {
		return -1
	}
	return startCount
}
