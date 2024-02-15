package database

import (
	"time"
)

type Blog struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Author    string    `json:"author"`
    LikeCount int       `json:"like_count"`
    CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
    ID         int       `json:"id"`
    BlogID     int       `json:"blog_id"`
    CommentText string    `json:"comment_text"`
    Author     string    `json:"author"`
    CreatedAt  time.Time `json:"created_at"`
}

type BlogInput struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}


func GetBlogs() ([]Blog, error) {
    rows, err := db.Query("SELECT * FROM blogs")

    if err != nil {
		return nil, err
	}
    defer rows.Close()

    var blogs []Blog

    for rows.Next() {
        var blog Blog
        err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.LikeCount, &blog.CreatedAt)
        if err != nil {
            return nil, err
        }
        blogs = append(blogs, blog)
    }

    return blogs, nil
}

func GetBlog(id int) (*Blog, error) {
    row := db.QueryRow("SELECT * FROM blogs WHERE id = $1", id)

    var blog Blog

    err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.LikeCount, &blog.CreatedAt)
    if err != nil {
        return nil, err
    }

    return &blog, nil
}


func CreateBlog(b *BlogInput) (*Blog, error) {
    var newBlog Blog
    err := db.QueryRow("INSERT INTO blogs (title, author, content) VALUES ($1, $2, $3) RETURNING *",
		b.Title, b.Author, b.Content).Scan(&newBlog.ID, &newBlog.Title, &newBlog.Content, &newBlog.Author, &newBlog.LikeCount, &newBlog.CreatedAt)

    if err != nil {
        return nil, err
    }

    return &newBlog, nil
}

func DeleteBlog(id int) (*Blog, error) {
    row := db.QueryRow("DELETE FROM blogs WHERE id = $1 RETURNING *", id)

    var blog Blog

    err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.LikeCount, &blog.CreatedAt)
    if err != nil {
        return nil, err
    }

    return &blog, nil
}