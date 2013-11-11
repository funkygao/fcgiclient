package main

import (
	"code.google.com/p/go-fastcgi-client"
	"flag"
	"fmt"
	"os"
)

func showVersion() {
	fmt.Fprintf(os.Stderr, "ver:%s (build: %s)\n", VERSION, BuildID)
}

// fastcgi client for locating error source
func main() {
	var (
		script  = flag.String("s", "/mnt/htdocs/royal/public/index.php", "server script absolute file path")
		uri     = flag.String("uri", "/", "request uri")
		reqData = flag.String("d", "", "send data in a POST request in form of k=v&k1=v1")
		host    = flag.String("h", "127.0.0.1", "fastcgi server host")
		port    = flag.Int("p", 9000, "fastcgi server port")
		version = flag.Bool("version", false, "show version")
	)
	flag.Parse()

	if *version {
		showVersion()
		os.Exit(0)
	}

	requestMethod := "GET"
	if *reqData != "" {
		requestMethod = "POST"
	}

	env := make(map[string]string)
	env["REQUEST_METHOD"] = requestMethod
	env["SCRIPT_FILENAME"] = *script
	env["REQUEST_URI"] = *uri
	env["SERVER_SOFTWARE"] = "go / fcgiclient "
	env["REMOTE_ADDR"] = "127.0.0.1"
	env["SERVER_PROTOCOL"] = "HTTP/1.1"
	env["QUERY_STRING"] = *reqData

	fcgi, err := fcgiclient.New(*host, *port)
	if err != nil {
		panic(err)
	}

	response, err := fcgi.Request(env, *reqData)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
