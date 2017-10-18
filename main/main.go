package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"encoding/json"
)

type User struct {
	Uid  int
	Name string
}

type Conversation struct {
	Uid       int
	CreatedAt int
	UpdatedAt int
	Messages  []Message
	User1     int
	User2     int
}

type Message struct {
	Uid       int
	CreatedAt int
	Unread    bool
	SentBy    int
	SentTo    int
	Body   string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var outJson, err = json.Marshal(usersSample)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Fprint(w, string(outJson))
}

func getUserById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}
func getConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var outJson, err = json.Marshal(conversationsSample)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Fprint(w, string(outJson))
}

func getConversationById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}
func getMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var outJson, err = json.Marshal(conversationsSample)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Fprint(w, string(outJson))
}

func getMessageById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are modify user %s", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	fmt.Fprintf(w, "you are add user %s", uid)
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)

	router.GET("/user/", getUsers)
	router.GET("/user/:uid", getUserById)

	router.GET("/conversation/", getConversations)
	router.GET("/conversation/:uid", getConversationById)

	router.GET("/message/", getConversations)
	router.GET("/message/:uid", getConversationById)

	router.POST("/adduser/:uid", adduser)
	router.DELETE("/deluser/:uid", deleteuser)
	router.PUT("/moduser/:uid", modifyuser)

	log.Fatal(http.ListenAndServe(":8080", router))
}

var currentUser = User{0, "Me"}
var otherUser1 = User{1, "Other"}
var otherUser2 = User{2, "Bob"}
var usersSample = []User{currentUser, otherUser1, otherUser2}

var message1 = Message{0, 1507967480, false, 0, 2, "Why do you do this"}
var message2 = Message{1, 1507967471, false, 0, 1, "Hey hows it going"}
var message3 = Message{2, 1507967500, false, 1, 0, "Good, and you"}
var message4 = Message{3, 1507967600, true, 2, 0, "I don't know"}
var messagesSample = []Message{message1, message2, message3, message4}

var conversation1 = Conversation{0, 1507967471, 1507967471, []Message{message1, message4}, 0, 2}
var conversation2 = Conversation{1, 1507967480, 1507967490, []Message{message2, message3}, 0, 1}

var conversationsSample = []Conversation{conversation1, conversation2}
