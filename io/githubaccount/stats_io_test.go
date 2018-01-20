// +build io

package githubaccount

// Call tests with local env variables set like:
// GITHUB_TOKEN="..." GITHUB_LOGIN="..." GITHUB_EMAIL="..." go test ./...
import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGitHubStatsClient(t *testing.T) {

	Convey("Given an invalid GitHub token", t, func() {
		token := "invalid token"

		Convey("When creating a StatsClient", func() {
			client, err := NewClient(token)

			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("The client should be nil", func() {
				So(client, ShouldBeNil)
			})

		})

	})

	Convey("Given an valid GitHub token", t, func() {
		token := os.Getenv("GITHUB_TOKEN")
		login := os.Getenv("GITHUB_LOGIN")
		So(token, ShouldNotBeEmpty)
		So(login, ShouldNotBeEmpty)

		client, err := NewClient(token)
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)

		Convey("When requesting the users GitHub stats", func() {
			stats, err := StatsOf(client)

			Convey("No error is returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The github login of the user is part of the stats", func() {
				So(*stats.Login, ShouldEqual, login)
			})

		})

	})

}
