// Package cert represents the certification subdomain.
// A "Certificate" is a uniquely identifiable and immutable struct
// which certifies that the owner of the certificate has executed a certain "Activity".
// All attributes of a "Certificate" are semi final, means they are immutable after being set.
// All attributes are also package private, so if you want to create a "Certificate"
// you have to use the exportet factory functions of this package.
package cert
