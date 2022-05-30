package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	usercli "parsing_http_response/internal"
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
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &user)
	return
}

func (u UserRepo) AddUser(user *usercli.User) error {
	fmt.Printf("Not implemented")
	panic(1)
}

func NewRemoteRepository() usercli.UserRepo {
	return &UserRepo{url: APIEndpoint}
}
