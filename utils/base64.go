package utils

import (
	"encoding/base64"

	"github.com/iskraman/golang-modules/syslog"
)

func Base64_Encoding(msg string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	return encoded
}

func Base64_Decoding(msg string) string {
	decoded, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		syslog.WAR("Base64 decoding error:", err)
		return ""
	}
	return string(decoded)
}
