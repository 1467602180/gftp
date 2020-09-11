package router

import (
	"gftp/app/api/dir"
    "github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		group.ALL("/dir", dir.Dir)
		group.ALL("/video", dir.Video)
		group.ALL("/download", dir.Download)
		group.ALL("/upload", dir.Upload)
	})
}
