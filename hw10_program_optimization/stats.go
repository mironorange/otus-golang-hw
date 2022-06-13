package hw10programoptimization

import (
	"bufio"
	"fmt"
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
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	s := bufio.NewScanner(r)

	i := 0
	for s.Scan() {
		line := s.Text()
		user := User{}
		if err = user.UnmarshalJSON([]byte(line)); err != nil {
			return
		}
		result[i] = user
		i++
	}

	if err := s.Err(); err != nil {
		return result, err
	}

	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	suffix := "." + domain
	matched := false

	for _, user := range u {
		matched = strings.HasSuffix(user.Email, suffix)
		if matched {
			name := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			num := result[name]
			num++
			result[name] = num
		}
	}

	return result, nil
}
