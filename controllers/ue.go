package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var allowImageType = "gif|jpeg|jpg|png|bmp"
var allowFileType = "rar|doc|docx|zip|pdf|txt|swf|mkv|avi|rm|rmvb|mpeg|mpg|ogg|mov|wmv|mp4|webm"
var website = "http://localhost:8080/"
var uploadfile = "static/upload/"
var uploadimage = "static/img/"

type UEController struct {
	beego.Controller
}

func (this *UEController) UploadImage() {
	filename := this.Input().Get("Filename")
	ext := filename[strings.LastIndex(filename, ".")+1:]
	if !strings.Contains(allowImageType, ext) {
		this.Ctx.WriteString("{'state':'FAILED'}")
	}
	newname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename
	err := this.SaveToFile("upfile", uploadimage+newname)
	state := "SUCCESS"
	if err != nil {
		fmt.Println(err)
		state = "FAILED"
	}
	url := website + uploadimage + newname
	this.Ctx.WriteString("{'original':'" + filename + "','url':'" + url + "','title':'" + this.Input().Get("pictitle") + "','state':'" + state + "'}")
}

func (this *UEController) UploadFile() {
	filename := this.Input().Get("Filename")
	index := strings.LastIndex(filename, ".")
	filetype := ""
	if index == -1 {
		this.Ctx.WriteString("{'state':'FAILED'}")
	}
	filetype = filename[index:]
	ext := filetype[1:]
	if !strings.Contains(allowFileType, ext) {
		this.Ctx.WriteString("{'state':'FAILED'}")
	}
	newname := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename
	err := this.SaveToFile("upfile", uploadfile+newname)
	state := "SUCCESS"
	if err != nil {
		fmt.Println(err)
		state = "FAILED"
	}
	url := website + uploadfile + newname
	this.Ctx.WriteString("{'url':'" + url + "','fileType':'" + filetype + "','state':'" + state + "','original':'" + filename + "'}")
}

func (this *UEController) ImageManager() {
	strRet := ""
	err := filepath.Walk(uploadimage, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ext := path[strings.LastIndex(path, ".")+1:]
		if strings.Contains(allowImageType, ext) {
			strRet += (path + "ue_separate_ue")
			fmt.Println("allow:", path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	fmt.Println(strRet)
	this.Ctx.WriteString(strRet)
}
