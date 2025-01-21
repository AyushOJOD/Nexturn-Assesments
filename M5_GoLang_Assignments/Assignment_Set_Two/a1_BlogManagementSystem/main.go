package main

import (
	db "BlogManagementSystem/config"
	"BlogManagementSystem/controller"
	"BlogManagementSystem/middleware"
	"BlogManagementSystem/repository"
	"BlogManagementSystem/services"
	"fmt"
	"net/http"
)

func main() {
	db.InititalizeDB()

	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := services.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	blogRoutes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			blogController.CreateBlog(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			fmt.Println("id received:", id)
			if id == "" {
				blogController.GetAllBlogs(w, r)
			} else {
				blogController.GetBlog(w, r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	paramRoutes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			blogController.UpdateBlog(w, r)
		case http.MethodDelete:
			blogController.DeleteBlog(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	protectedMux := http.NewServeMux()

	protectedMux.Handle("/blog", blogRoutes)
	protectedMux.Handle("/blog/", paramRoutes)

	protectedRoutes := middleware.AuthorizationMiddleware(db.GetDB(), protectedMux)

	http.Handle("/blog", protectedRoutes)
	http.Handle("/blog/", protectedRoutes)

	loggedMux := middleware.LoggingMiddleware(http.DefaultServeMux)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Println("Server failed to start", err)
	}
}
