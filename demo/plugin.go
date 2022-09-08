package demo

import (
	"context"

	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/runtime/event"
)

func init() {
	proxy.Plugins = append(proxy.Plugins, proxy.Plugin{
		Name: "Connect-Demo",
		Init: func(ctx context.Context, proxy *proxy.Proxy) error {
			return (&plugin{
				Proxy: proxy,
			}).init(ctx)
		},
	})
}

type plugin struct {
	*proxy.Proxy
}

func (p *plugin) init(ctx context.Context) error {
	event.Subscribe(p.Event(), 0, p.onPing)
	event.Subscribe(p.Event(), 0, p.onPreLogin)
	event.Subscribe(p.Event(), 0, p.onChooseServer)
	event.Subscribe(p.Event(), 0, p.onKickedFromServer)
	return nil
}
