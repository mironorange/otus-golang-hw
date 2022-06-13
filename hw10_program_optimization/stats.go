package hw10programoptimization

import (
	"bufio"
	"io"
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

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return getUsersAndCountDomains(r, domain)
}

func getUsersAndCountDomains(r io.Reader, domain string) (DomainStat, error) {
	s := bufio.NewScanner(r)

	result := make(DomainStat)
	suffix := "." + domain

	var pieces []string
	var name string
	var matched bool
	var user User
	for s.Scan() {
		line := s.Bytes()
		user = User{}
		if err := user.UnmarshalJSON(line); err != nil {
			return result, err
		}
		matched = strings.HasSuffix(user.Email, suffix)
		if matched {
			pieces = strings.SplitN(user.Email, "@", 2)
			name = strings.ToLower(pieces[1])
			num := result[name]
			num++
			result[name] = num
		}
	}

	if err := s.Err(); err != nil {
		return result, err
	}

	return result, nil
}
