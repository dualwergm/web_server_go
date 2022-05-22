package main

func main(){
	server := NewServer(":3000")
	server.AddHandle("/", "GET", HandlerRoot)
	server.AddHandle("/create", "POST", PostRequest)
	server.AddHandle("/user", "POST", UserPostRequest)
	server.AddHandle("/api", "POST", server.AddMiddleware(HandlerHome, CheckAuth(), Loggin()))
	server.Listen()
}