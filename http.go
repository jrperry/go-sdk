package iland

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (c *client) Get(endpoint string) (io.ReadCloser, error) {
	return c.request(endpoint, http.MethodGet, []byte{})
}

func (c *client) Post(endpoint string, body []byte) (io.ReadCloser, error) {
	return c.request(endpoint, http.MethodPost, body)
}

func (c *client) Put(endpoint string, body []byte) (io.ReadCloser, error) {
	return c.request(endpoint, http.MethodPut, body)
}

func (c *client) Delete(endpoint string) (io.ReadCloser, error) {
	return c.request(endpoint, http.MethodDelete, []byte{})
}

func (c *client) getObject(endpoint string, object interface{}) error {
	resp, err := c.Get(endpoint)
	if err != nil {
		return err
	}
	return unmarshalBody(resp, object)
}

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	GrantType    string `json:"grant_type"`
}

type RefreshTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type APIError struct {
	Error         string `json:"error"`
	Message       string `json:"message"`
	DetailMessage string `json:"detail_message"`
}

func (c *client) getToken() error {
	tokenRequest := TokenRequest{c.clientID, c.clientSecret, c.username, c.password, "password"}
	form := url.Values{}
	form.Add("client_id", tokenRequest.ClientID)
	form.Add("client_secret", tokenRequest.ClientSecret)
	form.Add("username", tokenRequest.Username)
	form.Add("password", tokenRequest.Password)
	form.Add("grant_type", tokenRequest.GrantType)
	resp, err := http.Post(accessURL, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Could not retrieve a token.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}
	c.Token = t
	c.setTokenExpiration()
	return nil
}

func (c *client) refreshToken() error {
	tokenRequest := RefreshTokenRequest{c.clientID, c.clientSecret, c.Token.RefreshToken, "refresh_token"}
	form := url.Values{}
	form.Add("client_id", tokenRequest.ClientID)
	form.Add("client_secret", tokenRequest.ClientSecret)
	form.Add("refresh_token", tokenRequest.RefreshToken)
	form.Add("grant_type", tokenRequest.GrantType)
	resp, err := http.Post(refreshURL, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Could not refresh current token.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}
	c.Token = t
	c.setTokenExpiration()
	return nil
}

func (c *client) setTokenExpiration() {
	c.tokenExpiration = time.Now().Add(time.Duration(c.Token.ExpiresIn-60) * time.Second)
}

func (c *client) request(relPath, verb string, payload []byte) (io.ReadCloser, error) {
	if c.Token.AccessToken == "" {
		err := c.getToken()
		if err != nil {
			return nil, err
		}
	}
	err := c.RefreshTokenIfNecessary()
	if err != nil {
		return nil, err
	}
	bytesJSON := bytes.NewBuffer(payload)
	req, err := http.NewRequest(verb, fmt.Sprintf("https://%s%s", apiHostname, relPath), bytesJSON)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token.AccessToken))
	req.Header.Add("Accept", "application/vnd.ilandcloud.api.v1.0+json")
	if verb == http.MethodPut || verb == http.MethodPost {
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 204 {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			return nil, errors.New(string(data))
		}
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	return resp.Body, nil
}

func (c *client) RefreshTokenIfNecessary() error {
	emptyToken := Token{}
	if c == nil || c.Token == emptyToken {
		err := c.getToken()
		if err != nil {
			return fmt.Errorf("Error retrieving iland cloud API token. %s", err.Error())
		}
	}
	if c.isTokenExpired() {
		err := c.refreshToken()
		if err != nil {
			err := c.getToken()
			if err != nil {
				return fmt.Errorf("Error refreshing iland cloud API token. %s", err.Error())
			}
		}
	}
	return nil
}

func (c *client) isTokenExpired() bool {
	if time.Now().After(c.tokenExpiration) {
		return true
	}
	return false
}

func unmarshalBody(body io.ReadCloser, object interface{}) error {
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	data = bytes.TrimPrefix(data, []byte(")]}'"))
	return json.Unmarshal(data, object)
}
