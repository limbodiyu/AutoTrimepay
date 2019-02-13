package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var email = "!!!change to your email!!!"
var password = "!!!change to your password!!!"

func main() {
	urlHome := "https://api.trimepay.com/"
	client := &http.Client{}
	rand.Seed(time.Now().UnixNano())

	letterRunes := []rune(`abcdefghijklmnopqrstuvwxyz`)
	csrf := make([]rune, 32)
	for csrfIndex := range csrf {
		csrf[csrfIndex] = letterRunes[rand.Intn(len(letterRunes))]
	}
	requestBody := url.Values{}
	requestBody.Set("email", email)
	requestBody.Set("password", password)
	request, _ := http.NewRequest(
		"POST",
		urlHome+"passport/auth/login?CSRF="+string(csrf),
		strings.NewReader(requestBody.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, _ := client.Do(request)
	cookies := response.Cookies()
	defer response.Body.Close()

	request, _ = http.NewRequest(
		"GET",
		urlHome+"merchant/app/dashboard?CSRF="+string(csrf),
		nil)
	for _, cookieIndex := range cookies {
		request.AddCookie(cookieIndex)
	}
	response, _ = client.Do(request)
	responseBody, _ := ioutil.ReadAll(response.Body)
	var responseBodyMap map[string]interface{}
	json.Unmarshal(responseBody, &responseBodyMap)
	balance := responseBodyMap["data"].(map[string]interface{})["merchant"].(map[string]interface{})["balance"].(float64)
	defer response.Body.Close()

	if balance <= 0 {
		return
	}
	requestBody = url.Values{}
	requestBody.Set("withdrawMethod", "1")
	requestBody.Set("totalFee", strconv.FormatFloat(balance, 'f', 0, 64))
	request, _ = http.NewRequest(
		"POST",
		urlHome+"merchant/withdraw/create?CSRF="+string(csrf),
		strings.NewReader(requestBody.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, cookieIndex := range cookies {
		request.AddCookie(cookieIndex)
	}
	response, _ = client.Do(request)
	defer response.Body.Close()
}
