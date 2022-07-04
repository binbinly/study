package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func HelloHandler(w http.ResponseWriter, request *http.Request)  {
	name := request.FormValue("name")
	w.Write([]byte(fmt.Sprintf("<h1>hello %s</h1>", name)))
}

func Index(w http.ResponseWriter, req *http.Request)  {
	bytes, err := ioutil.ReadFile("index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("没有找到文件")
		return
	}
	io.WriteString(w, string(bytes))
}

const (
	// 2 MB
	maxUploadSize = 2 * 1024 * 2014
	uploadPath    = "upload"
)

func Upload(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("文件大小超过限制")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("无效的file")
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("读取文件失败")
		return
	}
	// 获取文件类型
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpg" && fileType != "image/png" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("此文件类型是:%s,目前仅支持image/jpg和image/png格式 \n", fileType)
		return
	}
	// 组装文件保存路径
	fileName := strconv.FormatInt(time.Now().Unix(), 10)
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("获取文件扩展名失败")
		return
	}
	newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
	log.Printf("Get file type:%s, path: %s", fileType, newPath)
	// 判断上传的文件夹是否存在，不存在则创建
	_, err = os.Stat(uploadPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(uploadPath, 0666)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("文件夹不存在，创建失败")
				return
			}
		}
	}
	// 复制文件
	newFile, err := os.Create(newPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("创建文件失败")
		return
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileBytes);err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("写入文件失败")
		return
	}
	w.Write([]byte("上传成功"))
}

func Download(w http.ResponseWriter, r *http.Request)  {
	filePath := r.FormValue("filePath")
	file, err := os.Open(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("打开文件失败,err:%s \n", err.Error())
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("读取文件失败,err:%s \n", err.Error())
		return
	}
	w.Header().Add("Content-Disposition", "attachment;filename=\""+filePath+"\"")
	w.Header().Add("Content-Type", "application/octect-stream")
	w.Write(bytes)
}