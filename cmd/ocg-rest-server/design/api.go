package design

import (
	d "github.com/goadesign/goa/design"
	a "github.com/goadesign/goa/design/apidsl"
)

var _ = a.API("ocg", func() {
	a.Title("OCG Rest API")
	a.Description("OCG Rest API")
	a.Version("1.0")
	a.Host("localhost")
	a.Scheme("http")
	a.BasePath("/api")
	a.Consumes("application/json")
	a.Produces("application/json")

	a.License(func() {
		a.Name("GNU GPLv3")
		a.URL("https://www.gnu.org/licenses/gpl-3.0.txt")
	})
	a.Origin("/[localhost]/", func() {
		a.Methods("GET", "POST", "PUT", "PATCH", "DELETE")
		a.Headers("X-Request-Id", "Content-Type", "Authorization", "If-None-Match", "If-Modified-Since")
		a.MaxAge(600)
		a.Credentials()
	})

	a.Trait("GenericLinksTrait", func() {
		a.Attribute("self", d.String)
		a.Attribute("meta", a.HashOf(d.String, d.Any))
	})

	a.Trait("jsonapi-media-type", func() {
		a.ContentType("application/vnd.api+json")
	})

	a.Trait("conditional", func() {
		a.Headers(func() {
			a.Header("If-Modified-Since", d.String)
			a.Header("If-None-Match", d.String)
		})
	})

	a.ResponseTemplate(d.OK, func() {
		a.Description("Resource created")
		a.Status(200)
		a.Headers(func() {
			a.Header("Last-Modified", d.DateTime)
			a.Header("ETag")
			a.Header("Cache-Control")
		})
	})

	a.ResponseTemplate(d.Created, func(pattern string) {
		a.Description("Resource created")
		a.Status(201)
		a.Headers(func() {
			a.Header("Location", d.String, "href to created resource", func() {
				a.Pattern(pattern)
			})
		})
	})
})
