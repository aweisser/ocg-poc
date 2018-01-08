// Package design holds a goa design (https://goa.design/) which specifies the OCG REST-API.
// The goagen tool is used to generate different sources like the application logic, controllers and swagger files.
// Execute the goagen commands from project root every time the design changes!
// $ goagen app -d github.com/aweisser/ocg-poc/io/rest/design -o io/rest/
// $ goagen controller -d github.com/aweisser/ocg-poc/io/rest/design -o io/rest/controller/ --app-pkg github.com/aweisser/ocg-poc/io/rest/app
// $ goagen swagger -d github.com/aweisser/ocg-poc/io/rest/design -o io/rest/
package design
