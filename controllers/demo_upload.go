package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/astaxie/beego"
)

type UploadController struct {
	FFBaseController
}

func (c *UploadController) Get() {
	c.FFBaseController.Get()

	c.TplName = "demo_upload.html"
}

func (c *UploadController) Post() {
	c.FFBaseController.Post()

	beego.Info(c.Input())

	f, h, err := c.GetFile("myfile")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	// 保存位置在 static/upload, 没有文件夹要先创建
	if saveErr := c.SaveToFile("myfile", "static/upload/"+h.Filename); saveErr != nil {
		beego.Error(saveErr)
	}

	c.Data["json"] = h.Filename
	c.ServeJSON()
}

type ReadController struct {
	FFBaseController
}

//上传下载文件的页面
func (c *ReadController) Get() {

	c.TplName = "demo_upload.html"
}

//创建文件
type CreateController struct {
	FFBaseController
}

func (c *CreateController) Post() {
	//创建文件
	file, error := os.OpenFile("static/txtfile", os.O_CREATE|os.O_RDWR, 0666)
	//文件关闭
	defer file.Close()
	if error != nil {
		fmt.Println("创建文件失败")
		beego.Error(error)
	}
	c.Data["json"] = map[string]interface{}{"data": file.Name()}
	c.ServeJSON()
}

//写入文件
type WriteController struct {
	FFBaseController
}

func (c *WriteController) Post() {
	confPath := c.GetString("path")
	info := c.GetString("info")
	content, err := parseWriteConfig(confPath, info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
	c.Data["json"] = map[string]interface{}{"data": string(content)}
	c.ServeJSON()
}

//写入text文件内容
func parseWriteConfig(confPath, info string) ([]byte, error) {
	fl, err := os.OpenFile(confPath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer fl.Close()
	byteinfo := []byte(info)
	n, err := fl.Write(byteinfo)
	if err == nil && n < len(byteinfo) {
		fmt.Println("写入失败")
		fmt.Println(err)
	}
	return byteinfo, err
}

//读取文件内容
func (c *ReadController) Post() {
	confPath := c.GetString("path")
	fmt.Println("文件的地址:")
	fmt.Println(confPath)
	content, err := ReadFile(confPath)
	if err != nil {
		c.Data["data"] = ""
		fmt.Println(err)
	} else {
		c.Data["data"] = content
	}
	fmt.Println(content)
	c.Data["json"] = map[string]interface{}{"data": content}
	c.ServeJSON()
}

//解析text文件内容
func ReadFile(path string) (str string, err error) {
	//打开文件的路径
	fi, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件失败")
		fmt.Println(err)
	}
	defer fi.Close()
	//读取文件的内容
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println("读取文件失败")
		fmt.Println(err)
	}
	str = string(fd)
	return str, err
}

//删除文件
type DeleteController struct {
	FFBaseController
}

func (c *DeleteController) Post() {
	isdel := false
	file := c.GetString("path") //源文件路径
	err := os.Remove(file)      //删除文件
	if err != nil {
		//删除失败,输出错误详细信息
		fmt.Println(err)
	} else {
		//如果删除成功则输出
		isdel = true
	}
	c.Data["json"] = map[string]interface{}{"data": isdel}
	c.ServeJSON()
}
