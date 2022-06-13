package hw10programoptimization

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type users [100_000]User

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

func getUsers(r io.Reader) (result users, err error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	lines := bytes.Split(content, []byte{'\n'})
	for i, line := range lines {
		var user User
		if err := user.UnmarshalJSON(line); err != nil {
			return result, err
		}
		result[i] = user
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	suffix := "." + domain

	var pieces []string
	var name string
	var matched bool

	for _, user := range u {
		matched = strings.HasSuffix(user.Email, suffix)
		if matched {
			pieces = strings.SplitN(user.Email, "@", 2)
			name = strings.ToLower(pieces[1])
			num := result[name]
			num++
			result[name] = num
		}
	}

	return result, nil
}
