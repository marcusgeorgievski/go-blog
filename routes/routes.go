package routes

import (
	// . "goblog/logger"
	"goblog/routes/api"
	"goblog/routes/templates"
	"net/http"
)


func InitRoutes(mux *http.ServeMux) {
	// mux.HandleFunc("GET /{$}", Home)
	// mux.HandleFunc("GET /blogs", GetBlogs)
	// mux.HandleFunc("GET /blogs/{id}", GetBlog)
	// mux.HandleFunc("POST /blogs", CreateBlog)
	// mux.HandleFunc("DELETE /blogs/{id}", DeleteBlog)
	// mux.HandleFunc("GET /{slug...}", Catch)

	api.InitAPIRoutes(mux)
	templates.InitTemplateRoutes(mux)
}


// func Home(w http.ResponseWriter, r *http.Request) {
// 	Req.Println("GET /")
// 	http.ServeFile(w, r, "../static/index.html")
// }

// func Catch(w http.ResponseWriter, r *http.Request) {
// 	slug := r.PathValue("slug")
// 	Req.Println("GET /{" + slug + "}")
// 	Warn.Println("Redirecting to /")

// 	http.Redirect(w, r, "/", http.StatusSeeOther)

// }


// func GetBlogs(w http.ResponseWriter, r *http.Request) {
// 	Req.Println("GET /blogs")

// 	blogs, err := db.GetBlogs()

// 	if err != nil {

// 	}

// 	if blogs == nil {
// 		w.Write([]byte("No blogs found"))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(blogs)
// }

// func GetBlog(w http.ResponseWriter, r *http.Request) {
// 	slug := r.PathValue("id")
// 	Req.Println("GET /blogs/" + slug)

// 	id, err := strconv.Atoi(slug)

// 	if err != nil {}

// 	blog, err := db.GetBlog(id)

// 	if blog == nil {
// 		w.Write([]byte("No blog found"))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(blog)
// }

// func CreateBlog(w http.ResponseWriter, r *http.Request) {
// 	Req.Println("POST /blogs")

// 	err := r.ParseForm()

//     if err != nil {
//         http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
//         Error.Println("Failed to parse form data:", err)
//         return
//     }


// 	// Get form data
//     title := r.FormValue("title")
//     author := r.FormValue("author")
//     content := r.FormValue("content")

//     // Create BlogInput struct
//     b := db.BlogInput{
//         Title:  title,
//         Author: author,
// 		Content: content,
//     }

	
// 	blog, err := db.CreateBlog(&b)

// 	if err != nil {
// 		http.Error(w, "Failed to create blog post", http.StatusInternalServerError)
//         Error.Println("Failed to create blog post:", err)
//         return
// 	}
	
// 	Info.Println("New blog created")

// 	http.Redirect(w, r, "/blogs/" + strconv.Itoa(blog.ID), http.StatusSeeOther)
// }


// func DeleteBlog(w http.ResponseWriter, r *http.Request) {
// 	slug := r.PathValue("id")
// 	Req.Println("DELETE /blogs/" + slug)

// 	id, err := strconv.Atoi(slug)

// 	if err != nil {}

// 	blog, err := db.DeleteBlog(id)

// 	if err != nil {
// 		Error.Println(err.Error())
// 	}

// 	if blog == nil {
// 		w.Write([]byte("No blog found"))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(blog)
// }