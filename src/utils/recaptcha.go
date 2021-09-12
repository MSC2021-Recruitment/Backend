package utils

import (
	"MSC2021/src/global"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//RecaptchaResponse is ...
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

//RecaptchaRequest is ...
type RecaptchaRequest struct {
	Response string `json:"response" binding:"required"`
}

const recaptchaServerName = "https://www.recaptcha.net/recaptcha/api/siteverify"

//VerifyReCaptcha is ...
func VerifyReCaptcha(token string) bool {
	secretKey := global.CONFIG.ReCaptchaToken

	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {secretKey}, "response": {token}})
	if err != nil {
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			global.LOGGER.Sugar().Warnf("Verify ReCaptcha: read body failed: " + err.Error())
			return
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var responseData RecaptchaResponse
	if err := json.Unmarshal(body, &responseData); err != nil {
		return false
	}

	return true
}
