package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"github.com/hecomp/taskmanager/routers"
	"github.com/hecomp/taskmanager/common"
)

func main() {
	common.StartUp()
	r := routers.InitRouters()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Print("Listening...")
	http.ListenAndServe(":8080", n)
}
