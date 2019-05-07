package main

import (
  "log"
  "net/http"

  "github.com/ashlineldridge/gitbot/internal/github"
  "github.com/palantir/go-githubapp/githubapp"
)

func main() {
  githubConfig, err := github.LoadConfig()
  if err != nil {
    panic(err)
  }

  cc, err := githubapp.NewDefaultCachingClientCreator(
    *githubConfig,
    githubapp.WithClientUserAgent("gitbot/1.0.0"),
  )
  if err != nil {
    panic(err)
  }

  pushHandler := &github.PushHandler{
    ClientCreator: cc,
  }

  webhookHandler := githubapp.NewDefaultEventDispatcher(*githubConfig, pushHandler)

  http.Handle("/github/hook", webhookHandler)

  log.Printf("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
