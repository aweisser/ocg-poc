package githubaccount

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// NewClient returns an authenticated github client or an error if the authentication failed
func NewClient(token string) (*MultiAPIClient, error) {
	// assemble github.Client (REST API)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	restClient := github.NewClient(tc)

	// send a probe to verify connectibility
	_, _, err := restClient.Octocat(ctx, "")
	if err != nil {
		return nil, errors.New("Could not create a GitHub client. Perhaps your token is invalid")
	}

	return &MultiAPIClient{
		rest: restClient,
	}, nil
}

// MultiAPIClient provides access to the GitHub account
type MultiAPIClient struct {
	rest *github.Client
}

// User returns the authenticated github.User
func (sc *MultiAPIClient) User() (*github.User, error) {
	u, _, err := sc.rest.Users.Get(context.Background(), "")
	return u, err
}

// Stats from this authenticated GitHub account
type Stats struct {
	Login *string
}

// client interface defines the local contract for the main function "StatsOf"
type client interface {
	User() (*github.User, error)
}

// StatsOf uses a client collect stats of the authenticated GitHub user
func StatsOf(c client) (*Stats, error) {
	u, err := c.User()
	if err != nil {
		return nil, err
	}
	return &Stats{
		Login: u.Login,
	}, nil
}
