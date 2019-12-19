package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/elazarl/goproxy"
)

var (
	verboseFlag bool
	port        string

	postFlagParseFuncs []func()
)

func init() {
	flag.BoolVar(&verboseFlag, "verbose", true, "enable proxy request logging")
	flag.StringVar(&port, "p", ":1080", `HTTP proxy service address (e.g., ":1080")`)
}

func main() {
	flag.Parse()
	for _, f := range postFlagParseFuncs {
		f()
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verboseFlag

	var cnt uint64
	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		val := atomic.AddUint64(&cnt, 1)
		fmt.Printf("[%04d] HandleConnect :%s\n", val, ctx.Req.URL)
		return nil, ""
	})

	// FYI: [Goでnet/httpを使う時のこまごまとした注意 \- Qiita]( https://qiita.com/ono_matope/items/60e96c01b43c64ed1d18 )
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 32

	log.Printf("[log] port is '%s'\n", port)
	log.Fatal(http.ListenAndServe(port, proxy))
}
