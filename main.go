package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Auth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginSuccess struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type Person struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Photo string `json:"photo"`
	NIF   string `json:"nif"`
}

var authFire = Auth{
	Name:     "miguel",
	Password: "123456",
}

var tokenFire = "tokenFire"

const PHOTO = "https://source.unsplash.com/user/erondu/300x300"
const PORT = 8989
const RED = "\033[31m"
const GREEN = "\033[32m"
const BLUE = "\033[34m"

func auth(rw http.ResponseWriter, r *http.Request) {
	print(RED)
	if r.Method != "POST" {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	dt := time.Now()
	tokenFire = fmt.Sprintf("TheTokenIs%dMoreOrLess", dt.Nanosecond())

	b, _ := io.ReadAll(r.Body)
	bf := Auth{}

	if err := json.Unmarshal(b, &bf); err != nil {
		fmt.Println("Bad Request")
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	if bf.Name != authFire.Name || bf.Password != authFire.Password {
		log.Println("Unauthorized 🤢")
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}

	loginSuccess := LoginSuccess{Message: "Login success!", Token: tokenFire}
	js, _ := json.Marshal(loginSuccess)

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
	print(GREEN)
	log.Println(loginSuccess.Message, "👍")
}

func data(rw http.ResponseWriter, r *http.Request) {
	print(RED)
	if r.Method != "GET" {
		log.Println("Not Found 🧐")
		http.Error(rw, "Not Found", http.StatusNotFound)
		return
	}

	if fmt.Sprintf("Bearer %s", tokenFire) != r.Header.Get("Authorization") {
		log.Println("Forbidden 💩")
		http.Error(rw, "Forbidden", http.StatusForbidden)
		return
	}

	js, _ := json.Marshal([]Person{
		{Id: 1, Name: "YugarekPowa", Age: 67, Photo: PHOTO, NIF: "T2000"},
		{Id: 2, Name: "Kikirri", Age: 7, Photo: PHOTO, NIF: "enserio"},
		{Id: 3, Name: "Payete", Age: 17, Photo: PHOTO, NIF: "tekilea"},
		{Id: 4, Name: "Kerrunxi", Age: 62, Photo: PHOTO, NIF: "sinpc"},
		{Id: 5, Name: "Raulitill", Age: 97, Photo: PHOTO, NIF: "tw0wife"},
		{Id: 6, Name: "Borchelini", Age: 12, Photo: PHOTO, NIF: "estasiqsi"},
		{Id: 7, Name: "Maikirri", Age: 50, Photo: PHOTO, NIF: "rf4ever"},
		{Id: 8, Name: "Durkele", Age: 35, Photo: PHOTO, NIF: "gaferdealambrer"},
	})

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
	print(GREEN)
	log.Println("Authorization success! 👍")
}

func main() {
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/data", data)
	print(BLUE)
	log.Print(fmt.Sprintf("🥳 Server running on: %d", PORT))
	fmt.Println("\nROUTES")
	fmt.Println("-------------------------------")
	fmt.Println("POST http://localhost:8989/auth")
	fmt.Println("GET  http://localhost:8989/data")
	fmt.Println()
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
