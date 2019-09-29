package data

import (
	"fmt"
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"testing"
)

func TestGetAuthAttemptsOfSubject(t *testing.T) {

	subject := structs.Subject{
		Kind: fields.User,
		Name: "bender"}

	data, err := GetAuthAttemptsOfSubject(subject)

	if err != nil {
		t.Error("Failed to load auth attempts of subject")
	}

	fmt.Println(data)
}
