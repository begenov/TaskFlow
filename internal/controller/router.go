package controller

import "net/http"

func (c *controller) Router() http.Handler {
	mux := http.NewServeMux()
	return mux
}
