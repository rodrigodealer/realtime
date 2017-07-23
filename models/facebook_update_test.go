package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacebookUpdateUnmarshall(t *testing.T) {
	jsonText := "{\"object\":\"user\",\"entry\":[{\"uid\":\"123456789\",\"id\":\"123456789\",\"time\":1232313,\"changed_fields\":[\"friends\"]}]}"
	facebookUpdate := &FacebookUpdate{}
	facebookUpdate.FromJson(jsonText)
	assert.Equal(t, "user", facebookUpdate.Object)
	assert.Equal(t, "123456789", facebookUpdate.Entry[0].UID)
	assert.Equal(t, "123456789", facebookUpdate.Entry[0].ID)
}

func TestFacebookUpdateUnmarshallWithError(t *testing.T) {
	jsonText := "{\"object\":false?}"
	facebookUpdate := &FacebookUpdate{}
	err := facebookUpdate.FromJson(jsonText)
	assert.Contains(t, err.Error(), "invalid character")
}
