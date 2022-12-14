# Minekube Connect Demo

Minekube Connect Demo is a tiny open source demo implementation of
[Minekube Connect](https://developers.minekube.com/connect/).

Minekube Connect Demo allows you to connect any Minecraft server,
whether online mode, public, behind your protected home network or
anywhere else in the world, with our hosted demo implementation.

## Getting Started

1. Install the Connect Java plugin on your Minecraft server or proxy
2. Start and join your server publicly at `<server>.demo.minekube.net`

## Notes

The demo provides a working Connect network (Watch and Tunnel service)
running on a single Gate instance
as [implemented here](https://github.com/minekube/gate/blob/master/pkg/util/connectutil/single/single.go).

Please note that the demo is not meant for production use and players
are disconnected on every new commit to this demo repository's main branch
as it deploys new versions automatically.