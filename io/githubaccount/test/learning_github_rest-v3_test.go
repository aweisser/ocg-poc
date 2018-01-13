package test

// Call tests with local env variables set like:
// GITHUB_TOKEN="..." GITHUB_LOGIN="..." GITHUB_EMAIL="..." go test ./io/githubstats/...
import (
	"context"
	"os"
	"testing"

	gh "github.com/google/go-github/github"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/oauth2"
)

func TestGitHubRestAPI(t *testing.T) {

	Convey("Given a github token, login and an authenticated github rest client", t, func() {

		token := os.Getenv("GITHUB_TOKEN")
		login := os.Getenv("GITHUB_LOGIN")
		email := os.Getenv("GITHUB_EMAIL")
		So(token, ShouldNotBeEmpty)
		So(login, ShouldNotBeEmpty)
		So(email, ShouldNotBeEmpty)

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := gh.NewClient(tc)
		So(client, ShouldNotBeNil)

		Convey("When requesting the users login", func() {
			u, _, err := client.Users.Get(ctx, login)

			Convey("No error is returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The user is not nil", func() {
				So(u, ShouldNotBeNil)
			})

			Convey("The user email corresponds to the env variable GITHUB_EMAIL", func() {
				So(*u.Email, ShouldEqual, email)
			})

		})

	})

}
