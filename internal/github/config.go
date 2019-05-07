package github

import (
  "fmt"
  "os"
  "strconv"

  "github.com/palantir/go-githubapp/githubapp"
  "github.com/pkg/errors"
)

func LoadConfig() (*githubapp.Config, error) {
  appID, err := getIntEnv("GITHUB_APP_ID")
  if err != nil {
    return nil, err
  }

  secret, err := getEnv("GITHUB_APP_WEBHOOK_SECRET")
  if err != nil {
    return nil, err
  }

  privateKey, err := getEnv("GITHUB_APP_PRIVATE_KEY")
  if err != nil {
    return nil, err
  }

  return &githubapp.Config{
    WebURL:   "",
    V3APIURL: "",
    V4APIURL: "",
    App: struct {
      IntegrationID int    `yaml:"integration_id" json:"integrationId"`
      WebhookSecret string `yaml:"webhook_secret" json:"webhookSecret"`
      PrivateKey    string `yaml:"private_key" json:"privateKey"`
    }{
      IntegrationID: appID,
      WebhookSecret: secret,
      PrivateKey:    privateKey,
    },
    OAuth: struct {
      ClientID     string `yaml:"client_id" json:"clientId"`
      ClientSecret string `yaml:"client_secret" json:"clientSecret"`
    }{
      ClientID:     "",
      ClientSecret: "",
    },
  }, nil
}

func getEnv(key string) (string, error) {
  value, ok := os.LookupEnv(key)
  if !ok {
    return "", fmt.Errorf("required variable %s is not set", key)
  }
  return value, nil
}

func getIntEnv(key string) (int, error) {
  s, err := getEnv(key)
  if err != nil {
    return 0, err
  }

  i, err := strconv.Atoi(s)
  if err != nil {
    return 0, errors.Wrapf(err, "could not convert value of %s to an integer", key)
  }

  return i, nil
}
