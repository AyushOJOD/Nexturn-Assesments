package main

import (
	db "blog_exercise/config"
	"blog_exercise/controller"
	"blog_exercise/middleware"
	"blog_exercise/repository"
	"blog_exercise/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Routes:-
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			response := map[string]string{
				"App":     "Welcome to the Blog Management System!",
				"message": "Please Register & Login to Explore blogs or to contribute.",
			}
			json.NewEncoder(w).Encode(response)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Register Route
	http.Handle("/register", middleware.InputValidationMiddleware([]string{"name", "email", "password"})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		userController.RegisterUser(w, r)
	})))

	// Login Route
	http.Handle("/login", middleware.InputValidationMiddleware([]string{"email", "password"})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		userController.AuthenticateUser(w, r)
	})))

	// Protected routes
	http.Handle("/home", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, _ := middleware.GetUserIDFromContext(r)
		response := map[string]string{
			"App":     "Welcome to the Blog Management System!!",
			"message": "You have accessed a protected route!",
			"userID":  userID,
		}
		json.NewEncoder(w).Encode(response)
	})))

	// AuthMiddleware wraps the handler for authentication
	http.Handle("/blogs", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Input validation for POST
			middleware.InputValidationMiddleware([]string{"title", "content", "author"})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				blogController.CreateBlog(w, r)
			})).ServeHTTP(w, r)
		case http.MethodGet:
			blogController.GetAllBlogs(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})))

	// AuthMiddleware wraps the handler for authentication with ID-based operations
	http.Handle("/blogs/", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			blogController.GetBlog(w, r)
		case http.MethodPut:
			// Input validation for PUT
			middleware.InputValidationMiddleware([]string{"title", "content", "author"})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				blogController.UpdateBlog(w, r)
			})).ServeHTTP(w, r)
		case http.MethodDelete:
			blogController.DeleteBlog(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})))

	// Rate limiter setup
	rateLimiter := middleware.NewRateLimiter(10, time.Second)

	loggedAndRateLimitedMux := middleware.LoggingMiddleware(
		middleware.RateLimitingMiddleware(rateLimiter, http.DefaultServeMux),
	)

	// Start the server
	fmt.Println("Server is running on port 8081")
	if err := http.ListenAndServe(":8081", loggedAndRateLimitedMux); err != nil {
		fmt.Println("Error Starting Server:", err)
	}
}