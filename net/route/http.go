package net

import "net/http"

type HttpHandler struct {
	W http.ResponseWriter
	R *http.Request

}
