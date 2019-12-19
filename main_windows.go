package main

import (
	"flag"
	"net/http"

	"github.com/mattn/go-ieproxy"
)

var (
	ieproxyFlag bool
)

func init() {
	flag.BoolVar(&ieproxyFlag, "ieproxy", true, `detect the proxy settings on Windows platform`)

	postFlagParseFuncs = append(postFlagParseFuncs, func() {
		if ieproxyFlag {
			http.DefaultTransport.(*http.Transport).Proxy = ieproxy.GetProxyFunc()
		}
	})
}
