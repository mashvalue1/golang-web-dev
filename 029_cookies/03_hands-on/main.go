package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	// cookieを読み込み
	c, err := req.Cookie("my-cookie")

	// cookieがなかった場合は作成
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	// cookieのvalueをintにキャストしてcountに代入
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	// count増やす
	count++
	// cookiemのvalueをstringにキャストしたcountに更新
	c.Value = strconv.Itoa(count)

	// cookieにセット
	http.SetCookie(w, c)

	io.WriteString(w, c.Value)
}
