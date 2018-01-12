package github

// Call tests with local env variables set like:
// GITHUB_TOKEN="your personal github token" GITHUB_LOGIN="the corresponding user name" go test ./io/...
import (
	"context"
	"os"
	"testing"

	"github.com/shurcooL/githubql"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/oauth2"
)

func TestGitHubGraphQlAPI(t *testing.T) {

	Convey("Given a github token, login and an authenticated github client", t, func() {

		token := os.Getenv("GITHUB_TOKEN")
		login := os.Getenv("GITHUB_LOGIN")
		So(token, ShouldNotBeEmpty)
		So(login, ShouldNotBeEmpty)

		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient := oauth2.NewClient(context.Background(), src)
		client := githubql.NewClient(httpClient)
		So(client, ShouldNotBeNil)

		Convey("When requesting for the users login", func() {
			var query struct {
				Viewer struct {
					Login     githubql.String
					CreatedAt githubql.DateTime
				}
			}
			err := client.Query(context.Background(), &query, nil)
			Convey("No error is returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The result corresponds to the GITHUB_LOGIN environment variable", func() {
				So(query.Viewer.Login, ShouldEqual, os.Getenv("GITHUB_LOGIN"))
			})

		})

	})

}
