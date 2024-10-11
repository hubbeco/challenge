package services

import (
	"bytes"
	"challenge/config"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

var useRecaptchaMocks = false

func SetRecaptchaMocksEnabled(enabled bool) {
	useRecaptchaMocks = enabled
}

func ValidateRecaptcha(recaptchaResponse string) bool {

	if useRecaptchaMocks {
		return ValidateRecaptchaMock(recaptchaResponse)
	}

	secretKey := config.GetRecaptchaKey()
	recaptchaURL := config.GetRecaptchaURL()

	data := url.Values{}
	data.Set("secret", secretKey)
	data.Set("response", recaptchaResponse)

	resp, err := http.Post(recaptchaURL, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var result struct {
		Success bool `json:"success"`
	}

	err = json.Unmarshal(body, &result)
	return err == nil && result.Success
}
