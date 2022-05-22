package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router{
	return &Router {
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, existPath := r.rules[path]
	handler, existMethod := r.rules[path][method]
	return handler, existMethod, existPath
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request){
	handler, existMethod, existPath := r.FindHandler(request.URL.Path, request.Method)
	if !existPath {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !existMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}