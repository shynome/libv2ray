package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Required features. Can't remove unless there is replacements.
	_ "github.com/v2fly/v2ray-core/v4/app/dispatcher"
	_ "github.com/v2fly/v2ray-core/v4/app/proxyman/inbound"
	_ "github.com/v2fly/v2ray-core/v4/app/proxyman/outbound"

	// Other optional features.
	_ "github.com/v2fly/v2ray-core/v4/app/dns"
	_ "github.com/v2fly/v2ray-core/v4/app/dns/fakedns"
	_ "github.com/v2fly/v2ray-core/v4/app/log"
	_ "github.com/v2fly/v2ray-core/v4/app/policy"
	_ "github.com/v2fly/v2ray-core/v4/app/reverse"
	_ "github.com/v2fly/v2ray-core/v4/app/router"
	_ "github.com/v2fly/v2ray-core/v4/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/tagged/taggedimpl"

	// Inbound and outbound proxies.
	_ "github.com/v2fly/v2ray-core/v4/proxy/blackhole"
	_ "github.com/v2fly/v2ray-core/v4/proxy/dns"
	_ "github.com/v2fly/v2ray-core/v4/proxy/dokodemo"
	_ "github.com/v2fly/v2ray-core/v4/proxy/freedom"
	_ "github.com/v2fly/v2ray-core/v4/proxy/http"
	_ "github.com/v2fly/v2ray-core/v4/proxy/mtproto"
	_ "github.com/v2fly/v2ray-core/v4/proxy/shadowsocks"
	_ "github.com/v2fly/v2ray-core/v4/proxy/socks"
	_ "github.com/v2fly/v2ray-core/v4/proxy/trojan"
	_ "github.com/v2fly/v2ray-core/v4/proxy/vless/inbound"
	_ "github.com/v2fly/v2ray-core/v4/proxy/vless/outbound"
	_ "github.com/v2fly/v2ray-core/v4/proxy/vmess/inbound"
	_ "github.com/v2fly/v2ray-core/v4/proxy/vmess/outbound"

	// Transports
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/domainsocket"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/grpc"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/http"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/kcp"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/quic"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/tcp"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/tls"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/udp"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/websocket"

	// Transport headers
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/headers/http"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/headers/noop"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/headers/srtp"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/headers/tls"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/headers/utp"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/headers/wechat"
	_ "github.com/v2fly/v2ray-core/v4/transport/internet/headers/wireguard"

	// JSON config support. Choose only one from the two below.
	// The following line loads JSON from v2ctl
	// _ "github.com/v2fly/v2ray-core/v4/main/json"
	// The following line loads JSON internally
	_ "github.com/v2fly/v2ray-core/v4/main/jsonem"
)
