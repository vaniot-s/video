package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}

//获取文件地址
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if os.Getenv("STORAGETYPE")=="OSS" {
		/************qiniu oss***********/
		targetUrl := os.Getenv("OSSURL") + p.ByName("vid-id")
		http.Redirect(w, r, targetUrl, 301)
		/************aliyun oss**********/
		//targetUrl:="https://vaniot.oss-cn-beijing.aliyuncs.com/videos/"+p.ByName("vid-id")
		//http.Redirect(w,r,targetUrl,301)
	}else{
		/*********************local*******************/
		vid := p.ByName("vid-id")
		vl := VIDEO_DIR + vid

		video, err := os.Open(vl)
		if err != nil {
			sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
			return
		}
		w.Header().Set("Content-Type", "video/mp4")
		http.ServeContent(w, r, "", time.Now(), video)
		defer video.Close()
	}
}

//上传video
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	//校验文件的大小
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil { //检查文件
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	//获取文件
	file, _, err := r.FormFile("file") //上传的文件 返回的第二个参数是handler
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	//读文件
	data, err := ioutil.ReadAll(file) //读出文件
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	//写文件
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666) //路径 数据 权限
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	//ossfn := "videos/" + fn
	//path := "./videos/" + fn
	//bn := "avenssi-videos2"
	//ret := UploadToOss(ossfn, path, bn)
	//if !ret {
	//	sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	//	return
	//}
	//
	//os.Remove(path)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
