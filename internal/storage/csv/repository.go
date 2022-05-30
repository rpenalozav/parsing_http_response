package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"parsing_http_response/internal"
)

type UserRepo struct {
}

func NewCsvRepository() internal.UserRepo {
	return &UserRepo{}
}

func (p *UserRepo) AddUser(user *internal.User) error {
	f, err := os.Create("planet." + user.Login + ".csv")
	defer f.Close()

	csvWriter := csv.NewWriter(f)

	csvWriter.Write(user.ToArray())
	csvWriter.Flush()

	return err
}

func (p *UserRepo) GetUser(login string) (user *internal.User, err error) {
	fmt.Printf("Not implemented")
	panic(1)
}
