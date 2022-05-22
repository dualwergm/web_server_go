package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request){
			flag := true
			fmt.Println("Checking auth")
			if flag {
				f(w,r)
			}
		}
	}
}

func Loggin() Middleware{
	return func(f http.HandlerFunc) http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request){
			start := time.Now()
			defer func(){
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w,r)
		}
	}
}