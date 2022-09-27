package demo

import (
	"net"
	"strings"

	"go.minekube.com/common/minecraft/color"
	. "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/ping"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/netutil"
)

const (
	domainSuffix = ".demo.minekube.net"
	domainFormat = "<server>" + domainSuffix
)

func (p *plugin) onPing(e *proxy.PingEvent) {
	endpoint, ok := extractEndpoint(e.Connection().VirtualHost())
	if !ok {
		p.regularPong(e.Ping())
		return
	}

	server := p.Server(endpoint)
	if server == nil {
		endpointNotConnectedPong(e.Ping(), endpoint)
		return
	}

	endpointPong(e.Ping(), endpoint)
}

func extractEndpoint(addr net.Addr) (string, bool) {
	host := strings.ToLower(netutil.Host(addr))

	// Remove forge client specific suffix
	host = strings.Split(host, "\000")[0]

	if !strings.HasSuffix(host, domainSuffix) {
		return "", false
	}
	return strings.TrimSuffix(host, domainSuffix), true
}

func (p *plugin) regularPong(pong *ping.ServerPing) {
	pong.Description = text(
		Style{Color: color.Gold},
		"Demo Connect Network - Join %d servers with\ntheir domain %s",
		len(p.Servers()), domainFormat,
	)
}

func endpointNotConnectedPong(pong *ping.ServerPing, endpoint string) {
	pong.Description = endpointUnavailableText(endpoint)
}

func endpointPong(pong *ping.ServerPing, endpoint string) {
	pong.Description = successText("Server %s is connected\nJoin now to connect to it!", endpoint)
}
