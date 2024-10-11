package models

type ContactForm struct {
	RecaptchaResponse string `json:"g-recaptcha-response"`
	Comment           string `json:"comment"`
	Name              string `json:"name"`
	Email             string `json:"mail"`
}
