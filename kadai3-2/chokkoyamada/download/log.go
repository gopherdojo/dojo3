package download

import (
	"net/http"
	"os"
	"log"
)

func logHTTPRequest(req *http.Request) {
	if os.Getenv("DEBUG") != "" {
		log.Printf("<- %s %s", req.Method, req.URL)
		for key, values := range req.Header {
			for _, value := range values {
				log.Printf("<- %s: %s", key, value)
			}
		}
	}
}

func logHTTPResponse(res *http.Response) {
	if os.Getenv("DEBUG") != "" {
		log.Printf("=> %s, %s", res.Proto, res.Status)
		for key, values := range res.Header {
			for _, value := range values {
				log.Printf("-> %s: %s", key, value)
			}
		}
	}
}