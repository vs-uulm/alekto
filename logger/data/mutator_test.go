package data

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/time"
	"testing"
)

func TestLogAuthAttempt(t *testing.T) {

	authAttempt := structs.AuthAttempt{
		Subject: structs.Subject{
			Kind: fields.User,
			Name: "bender"},
		Timestamp: time.NowTimestamp(),
		Failed:    true,
		IPAddress: "127.0.0.1"}

	err := LogAuthAttempt(authAttempt)

	if err != nil {
		t.Error("failed to log auth attempt into log file")
	}
}
