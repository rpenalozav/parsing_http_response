package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	usercli "parsing_http_response/internal"
	"parsing_http_response/internal/errors"
)

const (
	APIEndpoint string = "https://api.github.com/"
	APIResource string = "users/"
)

type UserRepo struct {
	url string
}

func (u *UserRepo) GetUser(login string) (user *usercli.User, err error) {
	resp, err := http.Get(APIEndpoint + APIResource + login)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", APIEndpoint+APIResource)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", APIEndpoint+APIResource)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into user")
	}
	return
}

func (u UserRepo) AddUser(user *usercli.User) error {
	fmt.Printf("Not implemented")
	panic(1)
}

func NewRemoteRepository() usercli.UserRepo {
	return &UserRepo{url: APIEndpoint}
}
