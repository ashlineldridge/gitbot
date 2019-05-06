package github

import "github.com/palantir/go-githubapp/githubapp"

type gitHubApp struct {
  IntegrationID int
  WebhookSecret string
  PrivateKey    string
}

func LoadConfig() *githubapp.Config {
  return &githubapp.Config{
    WebURL:   "",
    V3APIURL: "",
    V4APIURL: "",
    App: struct {
      IntegrationID int    `yaml:"integration_id" json:"integrationId"`
      WebhookSecret string `yaml:"webhook_secret" json:"webhookSecret"`
      PrivateKey    string `yaml:"private_key" json:"privateKey"`
    }{
      IntegrationID: 0,
      WebhookSecret: "",
      PrivateKey:    "",
    },
    OAuth: struct {
      ClientID     string `yaml:"client_id" json:"clientId"`
      ClientSecret string `yaml:"client_secret" json:"clientSecret"`
    }{
      ClientID:     "",
      ClientSecret: "",
    },
  }
}
