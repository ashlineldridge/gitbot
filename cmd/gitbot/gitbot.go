package main

import (
  "github.com/ashlineldridge/gitbot/internal/github"
  "log"
	"net/http"

  "github.com/palantir/go-githubapp/githubapp"
)

func main() {
  githubConfig := github.LoadConfig()

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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
