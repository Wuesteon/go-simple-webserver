package main

import (
	"net/http"

	"github.com/Wuesteon/go-simple-webserver/controllers"
)

func main(){
	controllers.RegisterControllers()
	http.ListenAndServe(":3000",nil)

}