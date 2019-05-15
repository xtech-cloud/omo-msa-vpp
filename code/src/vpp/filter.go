package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func isFilter(_url string) bool {
	if config.Filter.Mode == "blacklist" {
		for _, url := range config.Filter.URL {
			if "" == url {
				continue
			}
			// 通配符后缀
			if strings.HasSuffix(url, "*") {
				// 过滤前缀
				if strings.HasPrefix(_url, url[0:len(url)-1]) {
					return true
				}
			} else {
				if url == _url {
					return true
				}
			}
		}
		// 未在黑名单中找到类似项,放行
		return false
	}

	if config.Filter.Mode == "whitelist" {
		for _, url := range config.Filter.URL {
			if "" == url {
				continue
			}
			// 通配符后缀
			if strings.HasSuffix(url, "*") {
				// 放行前缀
				if strings.HasPrefix(_url, url[0:len(url)-1]) {
					return false
				}
			} else {
				if url == _url {
					return false
				}
			}
		}
		// 未在白名单中找到类似项,过滤
		return true
	}

	// 模式配置错误,过滤
	return true
}

func filterMiddleware() gin.HandlerFunc {
	return func(_context *gin.Context) {
		url := _context.Request.URL.Path
		if isFilter(url) {
			_context.JSON(405, gin.H{
				"code":    405,
				"message": "method not allowed",
			})
			_context.Abort()
			return
		}

		_context.Next()
	}
}
