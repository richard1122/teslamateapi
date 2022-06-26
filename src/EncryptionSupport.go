package main

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strings"
)

type Region string

const (
	China  Region = "China"
	Global        = "Global"
)

// Get tesla web API endpoint url from iss in accessToken
// AccessToken is a JWT
func GetCarRegion(accessToken string) Region {
	payload := strings.Split(accessToken, ".")
	if len(payload) != 3 {
		return Global
	}
	decodedStr, err := base64.RawStdEncoding.DecodeString(payload[1])
	if err != nil {
		return Global
	}
	var result map[string]interface{}
	if err = json.Unmarshal(decodedStr, &result); err != nil {
		return Global
	}
	issUrl, err := url.Parse(result["iss"].(string))
	if err != nil {
		return Global
	}
	if strings.HasSuffix(issUrl.Host, ".cn") {
		return China
	}
	return Global
}
