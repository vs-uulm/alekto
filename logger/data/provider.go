package data

import (
	"github.com/ma-zero-trust-prototype/shared_lib/enum/fields"
	"github.com/ma-zero-trust-prototype/shared_lib/structs"
	"github.com/ma-zero-trust-prototype/shared_lib/time"
	"github.com/ma-zero-trust-prototype/shared_lib/yaml"
)

func GetAuthAttemptsOfSubject(subject structs.Subject) (data []structs.AuthAttempt, err error) {

	return GetAuthAttemptsOfSubjectSince(subject, time.YesterdayTimestamp())
}

func GetAuthAttemptsOfSubjectSince(subject structs.Subject, since int64) (
	data []structs.AuthAttempt, err error) {

	subjectIsDevice := subject.Kind == fields.Device

	if subjectIsDevice {
		data = getAttemptsByDeviceId(subject.Name, since)
	} else {
		data = getAttemptsBySubject(subject, since)
	}

	return
}

func getAttemptsBySubject(subject structs.Subject, since int64) (data []structs.AuthAttempt) {

	for _, attempt := range getAllAuthAttempts() {

		if len(data) > 100 {
			break
		}

		if subject.Equals(attempt.Subject) && attempt.Timestamp >= since {

			data = append(data, attempt)
		}
	}

	return
}

func getAttemptsByDeviceId(deviceId string, since int64) (data []structs.AuthAttempt) {

	for _, attempt := range getAllAuthAttempts() {

		if len(data) > 100 {
			break
		}

		if deviceId == attempt.DeviceId && attempt.Timestamp >= since {

			data = append(data, attempt)
		}
	}

	return
}

func getAllAuthAttempts() (attempts []structs.AuthAttempt) {

	yaml.LoadStructFromFile(authAttemptFilepath, &attempts)
	return
}
