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
)

func init() {
	flag.BoolVar(&verboseFlag, "verbose", true, "should every proxy request be logged to stdout")
	flag.StringVar(&port, "p", ":1080", `HTTP proxy service address (e.g., ":1080")`)
}

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verboseFlag

	var cnt uint64
	proxy.OnRequest().HandleConnectFunc(func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
		val := atomic.AddUint64(&cnt, 1)
		fmt.Printf("[%d] HandleConnect :%s\n", val, ctx.Req.URL)
		return nil, ""
	})

	//	FYI:
	//	http://qiita.com/ono_matope/items/60e96c01b43c64ed1d18
	// DefaultTransportの制限を変更する場合。
	// DefaultTransportはhttp.DefaultClientのTransportとして使われる。
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 32

	log.Println("port", port)
	log.Fatal(http.ListenAndServe(port, proxy))
}
