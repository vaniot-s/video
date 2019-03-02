package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)

	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)
	//
	router.POST("/userhome", userHomeHandler)

	//api代理转发
	router.POST("/api", apiHandler)

	//proxy solution: crosss resource sharing
	router.GET("/videos/:vid-id", proxyVideoHandler)

	//
	router.POST("/upload/:vid-id", proxyUploadHandler)

	//静态文件定向
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":9006", r)
}
