package main

import (
	// Connectors
	_ "github.com/go-gost/x/connector/direct"
	_ "github.com/go-gost/x/connector/http"
	_ "github.com/go-gost/x/connector/relay"

	// Dialers
	_ "github.com/go-gost/x/dialer/mtcp"
	_ "github.com/go-gost/x/dialer/mtls"
	_ "github.com/go-gost/x/dialer/mws"
	_ "github.com/go-gost/x/dialer/tcp"
	_ "github.com/go-gost/x/dialer/tls"
	_ "github.com/go-gost/x/dialer/udp"
	_ "github.com/go-gost/x/dialer/ws"

	// Handlers
	_ "github.com/go-gost/x/handler/forward/local"
	_ "github.com/go-gost/x/handler/relay"

	// Listeners
	_ "github.com/go-gost/x/listener/mtcp"
	_ "github.com/go-gost/x/listener/mtls"
	_ "github.com/go-gost/x/listener/mws"
	_ "github.com/go-gost/x/listener/tcp"
	_ "github.com/go-gost/x/listener/tls"
	_ "github.com/go-gost/x/listener/udp"
	_ "github.com/go-gost/x/listener/ws"
)
