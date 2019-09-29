package data

import (
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/yaml"
)

const (
	authAttemptFilepath = "data/authAttempts.yaml"
)

func LogAuthAttempt(attempt structs.AuthAttempt) (err error) {

	attempts := []structs.AuthAttempt{attempt}

	err = yaml.AppendStructToFile(authAttemptFilepath, attempts)
	return
}
