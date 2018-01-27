package sociallogin

import (
	"net/http"

	"github.com/danilopolani/gocialite"
	"github.com/danilopolani/gocialite/structs"
	"golang.org/x/oauth2"
)

const MediumDriverName = "medium"

// Endpoint is Mediums's explicit OAuth 2.0 endpoint.
// https://github.com/Medium/medium-api-docs/blob/master/README.md#21-browser-based-authentication
var MediumEndpoint = oauth2.Endpoint{
	AuthURL:  "https://medium.com/m/oauth/authorize",
	TokenURL: "https://api.medium.com/v1/tokens",
}

func init() {
	gocialite.RegisterNewDriver(
		MediumDriverName,
		MediumDefaultScopes,
		MediumUserFn,
		MediumEndpoint,
		MediumAPIMap,
		MediumUserMap,
	)
}

// MediumUserMap is the map to create the User struct
var MediumUserMap = map[string]string{
	"id":       "ID",
	"name":     "FullName",
	"username": "Username",
}

// MediumAPIMap is the map for API endpoints
var MediumAPIMap = map[string]string{
	"endpoint":     "https://api.medium.com",
	"userEndpoint": "/v1/me",
}

// MediumUserFn is a callback to parse additional fields for User
var MediumUserFn = func(client *http.Client, u *structs.User) {
}

// MediumDefaultScopes contains the default scopes
var MediumDefaultScopes = []string{}
