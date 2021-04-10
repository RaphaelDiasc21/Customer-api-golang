package main

import (
	"github.com/gorilla/mux"
	customerRouters "github.com/customers/routers"
	"net/http"
)

func main() {	
	r := mux.NewRouter()
	customerRouters.CustomersRouter(r)
	
	http.Handle("/",r)
	http.ListenAndServe(":4000",nil)
}
