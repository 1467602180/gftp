package dir

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func Download(r *ghttp.Request) {
	path:=r.Get("path")
	if path==nil {
		r.Response.WriteStatus(500,"缺少参数")
	}else{
		glog.Line().Println("请求下载:"+genv.Get("rootPath")+"/"+gconv.String(path))
		path = genv.Get("rootPath")+"/"+gconv.String(path)
		r.Response.ServeFileDownload(gconv.String(path))
	}
}
