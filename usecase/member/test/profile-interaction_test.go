package test

import (
	"testing"

	"github.com/aweisser/ocg-poc/domain/ocg"
	"github.com/aweisser/ocg-poc/usecase/member"
	. "github.com/smartystreets/goconvey/convey"
)

type testRepo struct {
	byID func(id interface{}) *ocg.Member
}

func (r *testRepo) GetMemberByID(id interface{}) *ocg.Member {
	return r.byID(id)
}

func TestMemberProfileInteraction(t *testing.T) {

	Convey("Given an empty repository", t, func() {
		pi := member.ProfileInteraction{
			MemberRepo: &testRepo{
				byID: func(id interface{}) *ocg.Member {
					return nil
				},
			},
		}
		Convey("When trying to get a member profile by id", func() {
			member, found := pi.LoadByID("some id")
			Convey("The returned member should be nil", func() {
				So(member, ShouldBeNil)
			})
			Convey("A second return value should indicate that the member could not be found", func() {
				So(found, ShouldBeFalse)
			})
		})
	})

	Convey("Given a repository with an existing member", t, func() {
		pi := member.ProfileInteraction{
			MemberRepo: &testRepo{
				byID: func(id interface{}) *ocg.Member {
					if id == "existing id" {
						return &ocg.Member{
							Name: "johndoe",
						}
					}
					return nil
				},
			},
		}
		Convey("When trying to get this member profile by id", func() {
			member, found := pi.LoadByID("existing id")
			Convey("The returned member should not be nil", func() {
				So(member, ShouldNotBeNil)
			})
			Convey("A second return value should indicate that the member could be found", func() {
				So(found, ShouldBeTrue)
			})
		})
		Convey("When trying to get another member profile by id", func() {
			member, found := pi.LoadByID("another id")
			Convey("The returned member should be nil", func() {
				So(member, ShouldBeNil)
			})
			Convey("A second return value should indicate that the member could not be found", func() {
				So(found, ShouldBeFalse)
			})
		})
	})

}
