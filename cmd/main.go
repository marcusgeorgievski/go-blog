package main

import (
	"goblog/database"
	. "goblog/logger"
	"goblog/routes"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	muxCors := corsMiddleware(mux)


	// Setup routes
	routes.InitBlogRoutes(mux)

	// Connect to database
	db, err := database.InitDB()
	if err != nil {
		Error.Println(err)
		return
	}
	database.SetDB(db)
	// defer db?


	// Start sever
	Info.Println("Server has started on port "+Clr.Yellow+"http://localhost:3000" + Clr.Reset)
	errHttp := http.ListenAndServe(":3000", muxCors); 
	
	if errHttp != nil {
		Error.Println("Error with server")
	}
}


func corsMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Call the next handler
        handler.ServeHTTP(w, r)
    })
}