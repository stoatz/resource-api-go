package main

import (
   "net/http"

   "myapi/controller"
   "myapi/model/repository"
)

//DI
var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main(){
   http.HandleFunc("/api/users/", ro.HandleTodosRequest)
   http.ListenAndServe(":9090", nil)
}
