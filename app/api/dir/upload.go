package dir

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func Upload(r *ghttp.Request) {
	file:=r.GetUploadFile("file")
	path:=r.Get("path")
	var filePath string
	if path==nil {
		filePath = genv.Get("rootPath")
	}else{
		filePath = genv.Get("rootPath")+"/"+gconv.String(path)
	}
	glog.Line().Println("上传文件:"+genv.Get("rootPath")+"/"+gconv.String(path))
	name,err:=file.Save(filePath)
	if err!=nil {
		glog.Debug(err)
	}
	r.Response.WriteExit("upload successfully: ", name)
}
