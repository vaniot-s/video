package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")

	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/*********************local*******************/
	//vid := p.ByName("vid-id")
	//vl := VIDEO_DIR + vid
	//
	//video, err := os.Open(vl)
	//if err != nil {
	//	sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	//	return
	//}
	//w.Header().Set("Content-Type", "video/mp4")
	//http.ServeContent(w, r, "", time.Now(), video)
	//defer video.Close()
	/************aliyun oss**********/
	//targetUrl:="https://vaniot.oss-cn-beijing.aliyuncs.com/videos/"+p.ByName("vid-id")
	//http.Redirect(w,r,targetUrl,301)
	/************qiniu oss***********/
	//http://pnq0o42mg.bkt.clouddn.com/0843795f-cae2-472b-a136-fcacc23ba24c
	targetUrl := "http://pq5cmm2db.bkt.clouddn.com/" + p.ByName("vid-id")
	http.Redirect(w, r, targetUrl, 301)
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
