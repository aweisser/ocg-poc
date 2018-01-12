package test

// Call tests with local env variables set like:
// GITHUB_TOKEN="..." GITHUB_LOGIN="..." GITHUB_EMAIL="..." go test ./io/githubstats/...
import (
	"context"
	"os"
	"testing"

	"github.com/shurcooL/githubql"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/oauth2"
)

func TestGitHubGraphQlAPI(t *testing.T) {

	Convey("Given a github token, login and an authenticated github GraphQL client", t, func() {

		token := os.Getenv("GITHUB_TOKEN")
		login := os.Getenv("GITHUB_LOGIN")
		email := os.Getenv("GITHUB_EMAIL")
		So(token, ShouldNotBeEmpty)
		So(login, ShouldNotBeEmpty)
		So(email, ShouldNotBeEmpty)

		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient := oauth2.NewClient(context.Background(), src)
		client := githubql.NewClient(httpClient)
		So(client, ShouldNotBeNil)

		Convey("When requesting for the users login and email", func() {
			var query struct {
				Viewer struct {
					Login githubql.String
					Email githubql.String
				}
			}
			err := client.Query(context.Background(), &query, nil)
			Convey("No error is returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The result corresponds to the environment variable GITHUB_LOGIN", func() {
				So(query.Viewer.Login, ShouldEqual, login)
			})

			Convey("The result corresponds to the environment variable GITHUB_EMAIL", func() {
				So(query.Viewer.Email, ShouldEqual, email)
			})

		})

	})

}

/*
{
  viewer {
    login
    name
    contributedRepositories(first:100){
      totalCount
      edges {
        node {
          name
          stargazers(first:0) {
            totalCount
            edges {
              node {
                name
              }
            }
          }
          watchers(first:0) {
            totalCount
            edges {
              node {
                id
              }
            }
          }
          forkCount
          isArchived
          isLocked
          deployments(first:0) {
            totalCount
            edges {
              node {
                id
              }
            }
          }
          commitComments(first:0) {
            totalCount
            edges {
              node {
                id
              }
            }
          }
        }
      }
    }
  }
}

*/
