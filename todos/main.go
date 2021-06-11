package main

import (
	"log"
	"net/http"
	"todos/app"

	"github.com/urfave/negroni"
)

func main() {

	mux := app.MakeHandler("./test.db")
	// mux := app.MakeHandler("./test.db") // 파일 경로를 안에 넣지 않아야 할 때는 flag 패키지를 사용 (method : flag.Args())
	defer mux.Close()

	n := negroni.Classic()

	n.UseHandler(mux)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)

	if err != nil {
		panic(err)
	}
}