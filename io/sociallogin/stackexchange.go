package sociallogin

import (
	"net/http"

	"github.com/danilopolani/gocialite"
	"github.com/danilopolani/gocialite/structs"
	"golang.org/x/oauth2"
)

// StackexchangeDriverName under which the driver is registered
const StackexchangeDriverName = "stackexchange"

// StackExchangeEndpoint is Stackexchange's implicit OAuth 2.0 endpoint.
// See https://api.stackexchange.com/docs/authentication for more information.
var StackExchangeEndpoint = oauth2.Endpoint{
	AuthURL:  "https://stackexchange.com/oauth",
	TokenURL: "https://stackexchange.com/oauth/access_token",
}

func init() {
	gocialite.RegisterNewDriver(
		StackexchangeDriverName,
		StackexchangeDefaultScopes,
		StackexchangeUserFn,
		StackExchangeEndpoint,
		StackexchangeAPIMap,
		StackexchangeUserMap,
	)
}

// StackexchangeUserMap is the map to create the User struct
var StackexchangeUserMap = map[string]string{
	"user_id":      "ID",
	"display_name": "FullName",
}

// StackexchangeAPIMap is the map for API endpoints
var StackexchangeAPIMap = map[string]string{
	"endpoint":     "https://api.stackexchange.com",
	"userEndpoint": "/2.2/me",
}

// StackexchangeUserFn is a callback to parse additional fields for User
var StackexchangeUserFn = func(client *http.Client, u *structs.User) {
	//u.Avatar = data["items"].([]interface{})[0].(map[string]interface{})["profile_image"].(string)
	u.Username = u.FullName
}

// StackexchangeDefaultScopes contains the default scopes
// Default scope is empty.
// Authentication will only allow an application to identify a user via the /me method.
// See https://api.stackexchange.com/docs/authentication#scope for more scopes.
var StackexchangeDefaultScopes = []string{}
