package speaking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeakerKnowsHowToSpeaks(t *testing.T) {
	assert.Equal(t, SpeakUp(), "Saying hello..", "SpeakUp to return the correct string")
}

func TestSpeakerKnowsHowToSayMyName(t *testing.T) {
	assert.Equal(t, SpeakUpWithV2("Itamar Arjuan"), "Hello Itamar Arjuan", "Say my name correctly")
}
