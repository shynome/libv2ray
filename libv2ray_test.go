package libv2ray

import (
	_ "embed"
	"os/exec"
	"testing"
)

//go:embed test/socks5.json
var testV2rayConfig string

func testProxyConnection() (err error) {
	_, err = exec.Command("curl", "-x", "socks5://127.0.0.1:1080", "https://ip.sb").CombinedOutput()
	return
}

func TestServer(t *testing.T) {
	err := Start(testV2rayConfig)
	if err != nil {
		t.Error(err)
		return
	}
	// test Start
	if err = testProxyConnection(); err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 3; i++ {
		// test Stop
		Stop()
		if err = testProxyConnection(); err == nil {
			t.Error(i, " Stop Server fail")
			return
		}
		// test Start agin
		Start(testV2rayConfig)
		if err = testProxyConnection(); err != nil {
			t.Error(i, " Second start Server fail", err)
			return
		}
	}
	return
}
