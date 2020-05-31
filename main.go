package main

import (
	"net/http"

	"github.com/zedonora/myapp"
)

func main() {

	http.ListenAndServe(":4000", myapp.NewHttpHandler())
}
