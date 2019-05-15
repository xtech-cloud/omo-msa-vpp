package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func getRemote(_url string) string {
	for _, forward := range config.Forward {
		endpoint := forward.Endpoint
		if strings.HasSuffix(endpoint, "*") {
			if strings.HasPrefix(_url, endpoint[0:len(endpoint)-1]) {
				return forward.Remote
			}
		} else {
			if _url == endpoint {
				return forward.Remote
			}
		}
	}
	return ""
}

func forward(_context *gin.Context) {
	endpoint := _context.Request.URL.Path
	remoteURL := getRemote(endpoint)
	if "" == remoteURL {
		_context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "remote not found",
		})
		return
	}
	remote, err := url.Parse(remoteURL)
	if nil != err {
		_context.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(_context.Writer, _context.Request)
}
