package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	transport "github.com/bozd4g/go-http-client"
)

var ServerError = errors.New("server error")
var ValidationError = errors.New("validation error")
var UserNotFound = errors.New("user not found")

type Client interface {
	Create(user User) error
	CheckPassword(username string, password string) (bool, error)
}

type client struct {
	transport transport.Client
}

func NewFromEnv() *client {
	url := os.Getenv("SERVICE_USER_URL")
	return &client{transport: transport.New(url)}
}

func (c *client) Create(user User) error {
	request, err := c.transport.PostWith("/user", user)
	if err != nil {
		return err
	}

	resp, err := c.transport.Do(request)
	if err != nil {
		return err
	}

	return c.getError(resp.Get())
}

func (c *client) CheckPassword(username string, password string) (bool, error) {
	endpoint := fmt.Sprintf("/user/%s/checkPassword?password=%s", username, password)

	request, err := c.transport.Get(endpoint)
	if err != nil {
		return false, err
	}

	resp, err := c.transport.Do(request)
	if err != nil {
		return false, err
	}

	respStruct := resp.Get()
	if respStruct.StatusCode == 404 {
		return false, UserNotFound
	}

	err = c.getError(respStruct)
	if err != nil {
		return false, err
	}

	var result CheckPassword
	err = json.Unmarshal(respStruct.Body, &result)
	if err != nil {
		return false, err
	}

	return result.Result.Valid, nil
}

func (c *client) getError(resp transport.ResponseStruct) error {
	if resp.StatusCode != 200 {
		if resp.StatusCode >= 400 && resp.StatusCode < 400 {
			return ValidationError
		}

		return ServerError
	}

	return nil
}
