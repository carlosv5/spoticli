package credentials

import (
	"testing"
)

func TestGet(t *testing.T) {
	credentials := Get()
	if len(credentials.Identifier) == 0 {
		t.Errorf("No App Identifier read in credentials")
	}
	if len(credentials.Secret) == 0 {
		t.Errorf("No App secret read in credentials")
	}
}
