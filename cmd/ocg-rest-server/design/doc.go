// Package design holds a goa design (https://goa.design/) which specifies the OCG REST-API.
// The goagen tool is used to generate different sources like the application logic, controllers and swagger files.
// Execute the goagen commands from project root every time the design changes!
// $ goagen app -d github.com/aweisser/ocg-poc/cmd/ocg-rest-server/design -o cmd/ocg-rest-server/
// $ goagen controller -d github.com/aweisser/ocg-poc/cmd/ocg-rest-server/design -o cmd/ocg-rest-server/controller/ --app-pkg github.com/aweisser/ocg-poc/cmd/ocg-rest-server/app
// $ goagen swagger -d github.com/aweisser/ocg-poc/cmd/ocg-rest-server/design -o cmd/ocg-rest-server/
//
// All available source (including a generated entry point of the application) is placed in the "generated" folder, for better diff'ing
// $ goagen bootstrap -d github.com/aweisser/ocg-poc/cmd/ocg-rest-server/design -o cmd/ocg-rest-server/generated --force
package design
