package main

import (
  "net/http"
  "log"
  "os"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
)

func main() {
  port := "8080"

  if fromEnv := os.Getenv("PORT"); fromEnv != "" {
    port = fromEnv
  }

  log.Printf("Starting up on http://localhost:%s", port)

  r := chi.NewRouter()

  r.Use(middleware.Logger)

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
  })

  r.Mount("/posts", postsResource{}.Routes())

  log.Fatal(http.ListenAndServe(":"+port, r))
}