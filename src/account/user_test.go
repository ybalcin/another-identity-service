package account

import (
	"fmt"
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/location"
)

func TestValidate(t *testing.T) {
	user := CreateNewUser(NewUserId(), "", "asdasd", "asdasd", "email", "asdasd", time.Now().UTC(), "asdasd",
		true, &location.Address{
			Country: "tr",
			City:    "tr",
			County:  "tr",
		})

	err := user.Validate()

	if err != nil || len(err) > 0 {
		for _, e := range err {
			fmt.Print(e)
			fmt.Printf("\n field: %v \n", e.Field())
			t.FailNow()
		}
	}
}
