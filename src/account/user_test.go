package account

import (
	"fmt"
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/location"
)

func TestValidate(t *testing.T) {
	_, err := NewUser(NewUserId(), "", "asdasd", "asdasd", "email", "asdasd", time.Now().UTC(), "asdasd",
		true, &location.Address{
			Country: "tr",
			City:    "tr",
			County:  "tr",
		})

	if len(err) == 0 {
		t.Errorf("Validation errors are not throwed!")
	}
}

func TestGetFieldValue(t *testing.T) {
	user, _ := NewUser(NewUserId(), "", "asdasd", "asdasd", "a@a.com", "asdasd", time.Now().UTC(), "asdasd",
		true, &location.Address{
			Country: "tr",
			City:    "tr",
			County:  "tr",
		})

	fieldValue := user.GetFieldValue("Email")
	fmt.Printf("%v", fieldValue)
}
