package main

import (
	"flag"
	"fmt"
	_ "gftp/boot"
	_ "gftp/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"net"
	"os/exec"
	"runtime"
)

func GetIp(port int)  {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		glog.Line().Println("获取ip地址失败，",err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						glog.Line().Println("开启服务成功，地址：http://"+ipnet.IP.String()+":"+gconv.String(port))
						Open("http://"+ipnet.IP.String()+":"+gconv.String(port))
					}
				}
			}
		}
	}
}

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

// Open calls the OS default program for uri
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

func main() {
	var path string
	var port int
	flag.StringVar(&path, "path", "./", "文件服务路径")
	flag.IntVar(&port, "port", 8199, "服务端口号")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	s := g.Server()
	genv.Set("rootPath",path)
	s.SetIndexFolder(true)
	s.SetServerRoot("template")
	s.SetPort(port)
	GetIp(port)
	s.Run()
}
