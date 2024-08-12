package main

import (
	"fmt"
	"net/http"
)

// registerWebUI is an empty stub because `webui` tag has not been specified.
func registerWebUI(hdl http.Handler) (http.Handler, error) {
	fmt.Println("testr")
	return hdl, nil
}
