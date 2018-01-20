package cert

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImmutableCertificate(t *testing.T) {

	Convey("Given a certificate that is already assigned", t, func() {
		cert := Certificate{ownerID: ID("User1")}
		Convey("When trying to assign a new owner", func() {
			err := cert.AssignTo(ID("User2"))
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a certificate that is already assigned", t, func() {
		cert := Certificate{ownerID: ID("User1")}
		Convey("When trying to assign a new owner", func() {
			err := cert.AssignTo(ID("User2"))
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given a certificate that is not asssigned yet", t, func() {
		cert := Certificate{}
		Convey("Assigning it to an owner", func() {
			err := cert.AssignTo(ID("User1"))
			Convey("Should not return an error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should set the owner of the certificate to the assigned ownerID", func() {
				So(cert.OwnerID(), ShouldEqual, ID("User1"))
			})
		})

	})

	Convey("Given a certificate with an activity", t, func() {
		originalName := "A1"
		originalFrom := time.Date(2018, time.Month(1), 1, 0, 0, 0, 0, time.Local)
		originalTo := time.Date(2018, time.Month(1), 2, 0, 0, 0, 0, time.Local)
		originalCat := "Category1"
		originalActivity := Activity{
			Name:     originalName,
			From:     originalFrom,
			To:       originalTo,
			category: originalCat,
		}
		cert := Certificate{
			activity: originalActivity,
		}
		Convey("When trying to change the attributes of the activity", func() {
			a := cert.Activity()
			a.Name = "A2"
			a.From = time.Date(2007, time.Month(3), 5, 0, 0, 0, 0, time.Local)
			a.To = time.Date(2007, time.Month(3), 6, 0, 0, 0, 0, time.Local)
			a.category = "Category2"
			Convey("The activity in the certificate should not change", func() {
				So(cert.Activity().Name, ShouldEqual, originalName)
				So(cert.Activity().From, ShouldEqual, originalFrom)
				So(cert.Activity().To, ShouldEqual, originalTo)
				So(cert.Activity().category, ShouldEqual, originalCat)
			})

		})
		Convey("The category of the activity", func() {
			a := cert.Activity()
			Convey("Should be accessable via an exported method", func() {
				So(a.Category(), ShouldEqual, originalCat)
			})
		})
	})

}

func TestCertificateFactory(t *testing.T) {

	Convey("Given a Certificate with an Activity, prefilled with category C1", t, func() {
		originalActivity := &Activity{Name: "Original Name", category: "Preset category", From: time.Now(), To: time.Now()}
		cert, _ := NewCertificate("Certified category", originalActivity)
		Convey("The certified activity", func() {
			certifiedActivity := cert.Activity()
			Convey("Should not contain the preset category", func() {
				So(certifiedActivity.Category(), ShouldEqual, "Certified category")
				So(certifiedActivity.category, ShouldEqual, "Certified category")
			})
		})
		Convey("When trying to modify the certified activity via the original activity", func() {
			originalActivity.Name = "Modified Name"
			Convey("The name of the certified activity should not change", func() {
				So(cert.Activity().Name, ShouldEqual, "Original Name")
			})
		})
	})

	Convey("Given a Certificate with an Activity having an empty category", t, func() {
		cert, _ := NewCertificate("Certified category", &Activity{From: time.Now(), To: time.Now()})
		Convey("The category of the certified activity", func() {
			certifiedActivity := cert.Activity()
			Convey("Should not be empty", func() {
				So(certifiedActivity.Category(), ShouldEqual, "Certified category")
			})
		})

	})

	Convey("Giving a nil Activity", t, func() {
		Convey("When trying to create a Certificate", func() {
			_, err := NewCertificate("Certified category", nil)
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Giving an Activity without From dates", t, func() {
		a := &Activity{
			Name:     "A name",
			category: "A Category",
		}
		Convey("When trying to create a Certificate", func() {
			_, err := NewCertificate("Certified category", a)
			Convey("An error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given two Certificates", t, func() {
		a := Activity{Name: "Original Name", From: time.Now(), To: time.Now()}
		cert1, _ := NewCertificate("Certified category", &a)
		cert2, _ := NewCertificate("Certified category", &a)
		Convey("The two IDs", func() {
			id1 := cert1.id
			id2 := cert2.id
			Convey("Should be different", func() {
				So(id1, ShouldNotEqual, id2)
			})
		})
	})

}
