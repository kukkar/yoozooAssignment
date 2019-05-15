package router

import (
	"log"
	"net/http"
	"yoozooAssignment/helloapi/controller"
)

func BootStrap() {
	http.HandleFunc("/", controller.HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
