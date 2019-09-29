package authentication

type Method int

const (
	UnknownAuth   Method = 0
	BasicAuth     Method = 1
	TwoFactorAuth Method = 2
)

// should be const
var authNames = [...]string{
		"unknownAuth",
		"basicAuth",
		"twoFactorAuth"}

/**
 * to String function for authentication methods
 */
func (method Method) String() string {

	if len(authNames) < int(method) {

		return authNames[UnknownAuth]
	}

	return authNames[method]
}

/**
 * parse string to authentication method enum
 */
func ParseString (rawValue string) (authenticationMethod Method) {

	for index, name := range authNames {

		if name == rawValue {

			return Method(index)
		}
	}

	return UnknownAuth
}
