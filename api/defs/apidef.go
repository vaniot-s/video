package defs

//request 数据结构
type UserCredential struct {
	Username string `json:"user_name"` //tag名为use_name
	Pwd      string `json:"pwd"`
}
type NewComment struct {
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
}

type NewVideo struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

/**
tag 在反序列化时会生成
{
 user_name xxx
}
**/

//response 返回数据结构
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type UserSession struct {
	Username  string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type UserInfo struct {
	Id int `json:"id"`
}

type SignedIn struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

//data model
//video
type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}
type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type Comment struct {
	Id      string `json:"id"`
	VideoId string `json:"video_id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

//sessionMap
type SimpleSession struct {
	Username string //login name
	TTL      int64
}
