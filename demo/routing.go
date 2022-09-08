package demo

import (
	"fmt"
	"sort"

	"go.minekube.com/common/minecraft/color"
	. "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func (p *plugin) onPreLogin(e *proxy.PreLoginEvent) {
	endpoint, ok := extractEndpoint(e.Conn().VirtualHost())
	if !ok {
		e.Deny(&Text{Extra: []Component{
			errText("You can connect to servers with\ntheir domain %s\n\n", domainFormat),
			serversText(p.Servers()),
		}})
		return
	}

	server := p.Server(endpoint)
	if server == nil {
		e.Deny(endpointUnavailableText(endpoint))
		return
	}
}

func (p *plugin) onChooseServer(e *proxy.PlayerChooseInitialServerEvent) {
	endpoint, ok := extractEndpoint(e.Player().VirtualHost())
	if !ok {
		return
	}

	server := p.Server(endpoint)
	if server == nil {
		e.Player().Disconnect(endpointUnavailableText(endpoint))
		return
	}

	e.SetInitialServer(server)
}

func (p *plugin) onKickedFromServer(e *proxy.KickedFromServerEvent) {
	e.SetResult(&proxy.DisconnectPlayerKickResult{
		Reason: &Text{Extra: []Component{
			errText("You were kicked from the server %s:\n\n", e.Server().ServerInfo().Name()),
			e.OriginalReason(),
		}},
	})
}

const maxServersShown = 10

func serversText(a []proxy.RegisteredServer) *Text {
	if len(a) == 0 {
		return &Text{}
	}

	// Sort by most players first
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].Players().Len() > a[j].Players().Len()
	})

	builder := &Text{
		S:       Style{Color: color.Yellow},
		Content: fmt.Sprintf("Available Servers (%d):\n\n", len(a)),
	}
	n := min(len(a), maxServersShown)
	for i := 0; i < n; i++ {
		svr := a[i]
		builder.Extra = append(builder.Extra, &Text{
			Content: serverDomain(svr.ServerInfo().Name()),
			Extra: []Component{
				&Text{Content: fmt.Sprintf(" - %d players\n", svr.Players().Len())},
			},
		})
	}
	if len(a) > n {
		builder.Extra = append(builder.Extra, &Text{
			Content: fmt.Sprintf("and %d more...", len(a)-n),
		})
	}

	return builder
}

func errText(format string, a ...interface{}) *Text {
	return text(Style{
		Color: color.Red,
	}, format, a...)
}

func successText(format string, a ...interface{}) *Text {
	return text(Style{
		Color: color.Green,
	}, format, a...)
}

func text(s Style, format string, a ...interface{}) *Text {
	return &Text{
		Content: fmt.Sprintf(format, a...),
		S:       s,
	}
}

func endpointUnavailableText(endpoint string) *Text {
	return errText("Server %s is not connected to\nthis Demo Connect network", endpoint)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func serverDomain(endpoint string) string {
	return endpoint + domainSuffix
}
