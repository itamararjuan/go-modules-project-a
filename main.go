// Package A essentially..
// The main project

package speaking

import (
	b "github.com/itamararjuan/go-modules-project-b"
	bb "github.com/itamararjuan/go-modules-project-b/v2"
)

// SpeakUp - allows speaking in certain languages
func SpeakUp() string {
	// The implementation of speak up needs to use the SayHello
	// Method from package-b
	return b.SayHelloInEnglish()
}

// SpeakUpWithV2 is a different implementation that uses a higher major version
// of the same Go module.
func SpeakUpWithV2(name string) string {
	return bb.SayHelloInEnglish(name)
}
