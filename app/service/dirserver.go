package service

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/text/gstr"
	"io/ioutil"
	"strings"
)

func ListDir(dirPth string, suffix string) (files *garray.Array, err error) {
	files = garray.NewArray(true)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		fileData := gmap.New(true)
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			if fi.IsDir() { // 忽略目录
				fileData.Set("isDir", true)
			} else {
				fileData.Set("isDir", false)
			}
			if gstr.SubStr(dirPth, len(dirPth)-1, len(dirPth)) != "/" {
				dirPth += "/"
			}
			fileData.Set("name", fi.Name())
			fileData.Set("pwd", dirPth)
			files.Append(fileData)
		}
	}
	return files, nil
}
