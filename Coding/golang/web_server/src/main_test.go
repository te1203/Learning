package main

import (
	"fmt"
	"io"
	"net"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/valyala/fasthttp"
)

const (
	GET  = "GET"
	POST = "POST"
)

func TestProtected(t *testing.T) {
	port := 8880

	router := RouterInit()
	closer := startServerOnPort(t, port, router.Handler)
	defer closer.Close()

	Convey("Test a protected resource", t, func() {

		Convey("with basic authentication", func() {
			requestBodyString := doHTTPRequest(fmt.Sprintf("http://127.0.0.1:%d/protected/", port),
				GET,
				"Basic dGVzdHVzZXI6dGVzdHVzZXIhISE=",
				nil)

			So(requestBodyString, ShouldStartWith, "Protected")
		})

		Convey("without basic authentication", func() {
			requestBodyString := doHTTPRequest(fmt.Sprintf("http://127.0.0.1:%d/protected/", port),
				GET,
				"",
				nil)
			So(requestBodyString, ShouldStartWith, "Unauthorized")
		})
	})

}

func doHTTPRequest(url, method, basicAuth string, postBody []byte) string {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// release the resources
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// set basic authentication in the header
	if basicAuth != "" {
		req.Header.Set("Authorization", basicAuth)
	}
	req.Header.SetMethod(method)

	req.SetRequestURI(url)

	if method == POST {
		req.SetBody(postBody)
	}

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("request failed with error: ", err.Error())
		return ""
	}

	body := resp.Body()
	requestBodyString := string(body)

	return requestBodyString
}

func startServerOnPort(t *testing.T, port int, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %d: %s", port, err)
	}
	go fasthttp.Serve(ln, h)
	return ln
}
