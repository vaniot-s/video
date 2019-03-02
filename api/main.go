package main

import (
	"github.com/julienschmidt/httprouter"
	"video/api/session"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

//注册路由
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)           //用户注册
	router.POST("/user/:username", Login)      //登录
	router.GET("/user/:username", GetUserInfo) //用户信息
	router.GET("/test", Test)
	router.POST("/user/:username/videos", AddNewVideo)

	router.GET("/user/:username/videos", ListAllVideos)

	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	router.POST("/videos/:vid-id/comments", PostComment)

	router.GET("/videos/:vid-id/comments", ShowComments)
	return router
}
func Prepare() {
	session.LoadSessionsFromDB()
}
func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":9003", mh) //handler
}

//main->middleware->defs(message,err)->handlers->dbops->response
