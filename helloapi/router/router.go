package router

import (
	"fmt"
	"net/http"
	"yoozooAssignment/helloapi/controller"
)

//
//BootStrap the Server
//
func BootStrap() {
	//Home Page route
	http.HandleFunc("/", controller.HomePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
