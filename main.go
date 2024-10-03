package main

import (
	"fmt"
	"net/http"

	"github.com/echen805/web-development-go/controllers"
	"github.com/echen805/web-development-go/models"
	"github.com/echen805/web-development-go/templates"
	"github.com/echen805/web-development-go/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.MustParse(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(
		views.MustParse(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.MustParse(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup controllers
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.MustParse(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	usersC.Templates.SignIn = views.MustParse(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix before deploying
		csrf.Secure(false),
	)
	http.ListenAndServe(":3000", csrfMw(r))
}
