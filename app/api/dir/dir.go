package dir

import (
	"gftp/app/service"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func Dir(r *ghttp.Request) {
	path:=r.Get("path")
	var files *garray.Array
	var err error
	rootPath:=genv.Get("rootPath")
	if rootPath == "./" {
		rootPath = gfile.Pwd()
	}
	if path==nil {
		files, err = service.ListDir(rootPath, "")
	}else{
		files, err = service.ListDir(rootPath+"/"+gconv.String(r.Get("path")), "")
	}
	glog.Line().Println("请求路径:"+rootPath+"/"+gconv.String(r.Get("path")))
	if err == nil {
		r.Response.WriteJson(g.Map{
			"status": "success",
			"data":   files,
		})
	} else {
		r.Response.WriteJson(g.Map{
			"status": "fail",
			"data":   g.Array{},
		})
	}
}
