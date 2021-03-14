package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	beego "github.com/beego/beego/v2/server/web"
)

// 准备数据
var datalist []string
var classes []string
var workat int

func init() {
	res, err := filepath.Glob("DATA/data/*")
	if err != nil {
		panic(err)
	}
	datalist = res
	r, err := filepath.Glob("DATA/out/*")
	if err != nil {
		panic(err)
	}
	for _, i := range r {
		classes = append(classes, filepath.Base(i))
	}
	workat = 0
	fmt.Println(res)
}

type ClassesImg struct {
	beego.Controller
}

func (c *ClassesImg) Get() {
	res := c.GetString("classes")
	for _, it := range classes {
		if it == res {
			savetoP := "DATA/out/" + res + "/" + filepath.Base(datalist[workat])
			os.Rename(datalist[workat], savetoP)
			break
		}
	}
	workat++

	if workat >= len(datalist) {
		c.TplName = "over.html"
	} else {
		c.Data["classes"] = classes
		c.Data["url"] = "/imgdata/" + filepath.Base(datalist[workat])
		c.TplName = "simple.html"
	}
}
