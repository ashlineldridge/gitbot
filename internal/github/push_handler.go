package github

import (
  "context"
  "encoding/json"
  hub "github.com/google/go-github/github"
  "github.com/palantir/go-githubapp/githubapp"
  "github.com/pkg/errors"
  "github.com/rs/zerolog/log"
)

type PushHandler struct {
  githubapp.ClientCreator
}

func (h *PushHandler) Handles() []string {
  return []string{"issue_comment"}
}

func (h *PushHandler) Handle(ctx context.Context, eventType, deliveryID string, payload []byte) error {
  var event hub.PushEvent
  if err := json.Unmarshal(payload, &event); err != nil {
    return errors.Wrap(err, "failed to parse push event payload")
  }

  log.Printf("Received push event for ref %s on repo %s", *event.Ref, *event.Repo.GitURL)

  return nil
}
