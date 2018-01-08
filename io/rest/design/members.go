package design

import (
	d "github.com/goadesign/goa/design"
	a "github.com/goadesign/goa/design/apidsl"
)

var _ = a.Resource("member", func() {
	a.BasePath("/members")

	a.Action("show", func() {
		a.Routing(
			a.GET("/:memberID"),
		)
		a.Description("Retrieve member (as JSONAPI) for the given ID.")
		a.Params(func() {
			a.Param("memberID", d.String, "ID of the member")
		})
		a.UseTrait("conditional")
		a.Response(d.OK, memberSingle)
		a.Response(d.NotModified)
		a.Response(d.BadRequest, JSONAPIErrors)
		a.Response(d.InternalServerError, JSONAPIErrors)
		a.Response(d.NotFound, JSONAPIErrors)
	})
})

var memberSingle = JSONSingle(
	"Member", "Holds a single response to a member request",
	member,
	nil)

var member = a.Type("Member", func() {
	a.Attribute("type", d.String, "The type of the related resource", func() {
		a.Enum("members")
	})
	a.Attribute("id", d.String, "The unique ID of the member", func() {
		a.Example("4711xyz")
	})
	a.Attribute("attributes", memberAttributes)
	a.Attribute("links", genericLinksForMember)
	a.Required("type", "attributes")
	a.Attribute("relationships", memberRelationships)
})

var memberAttributes = a.Type("memberAttributes", func() {
	a.Attribute("name", d.String, "Name of the member", func() {
		a.MaxLength(62) // maximum name length is 62 characters
		a.MinLength(1)  // minimum name length is 1 characters
		a.Pattern("^[^_|-].*")
		a.Example("Johndoe42")
	})
	a.Required("name")
})

var genericLinksForMember = a.Type("GenericLinksForMember", func() {
	a.Attribute("self", d.String)
	a.Attribute("meta", a.HashOf(d.String, d.Any))
})

var memberRelationships = a.Type("MemberRelationships", func() {
	a.Attribute("accounts", relationGeneric, "A Member can link external accounts like github or stackoverflow.")
})
