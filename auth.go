package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthenticationResponse struct {
	RequestType string `json:"requestType"`
	Error       string `json:"error"`
	Token       string `json:"token"`
}

// UserInfo ...
type UserInfo struct {
	Username string
	Token    string
}

var _userInfo UserInfo

func SetUserInfo(userInfo UserInfo) {
	_userInfo = userInfo
}

func GetUserInfo() UserInfo {
	return _userInfo
}

func UserAuthenticated() bool {
	return _userInfo.Token != ""
}

func Authenticate(username string, password string) (string, error) {
	fmt.Printf("\nAuthenticate -> %s-%s", username, password)

	BaseURL := "http://localhost:8080"
	req := struct {
		RequestType string `json:"type"`
		Username    string `json:"username"`
		Password    string `json:"password"`
	}{Username: username, Password: password, RequestType: "authenticate"}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("imposssible marshalling error: %s", err.Error())
	}

	r := bytes.NewReader(reqBytes)
	resp, err := http.Post(BaseURL, "application/json", r)
	if err != nil {
		return "", fmt.Errorf("request execution failed: %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request execution failed: status code: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %d", resp.StatusCode)
	}

	response := AuthenticationResponse{}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if response.Error != "" {
		return "", fmt.Errorf("authentication failed: %s", response.Error)
	}
	return response.Token, nil
}
