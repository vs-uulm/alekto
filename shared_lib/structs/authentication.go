package structs

import "github.com/ma-zero-trust-prototype/shared_lib/enum/authentication"

type AuthAttempt struct {
	Subject    Subject               `yaml:"subject"`
	UserAuth   authentication.Method `yaml:"userAuthMethod"`
	DeviceAuth string                `yaml:"deviceAuthMethod"`
	Timestamp  int64                 `yaml:"timestamp"`
	Failed     bool                  `yaml:"failed"`
	IPAddress  string                `yaml:"ipaddress"`
	DeviceId   string                `yaml:"deviceid"`
}
