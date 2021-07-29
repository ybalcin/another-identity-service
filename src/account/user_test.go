package account

import (
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/location"
)

func TestValidate(t *testing.T) {
	_, err := CreateNewUser(NewUserId(), "", "asdasd", "asdasd", "email", "asdasd", time.Now().UTC(), "asdasd",
		true, &location.Address{
			Country: "tr",
			City:    "tr",
			County:  "tr",
		})

	if len(err) == 0 {
		t.Errorf("Validation errors are not throwed!")
	}
}
