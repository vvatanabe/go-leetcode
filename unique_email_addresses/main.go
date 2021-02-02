package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	set := make(map[string]struct{})
	for _, email := range emails {
		local2host := strings.Split(email, "@")
		local := local2host[0]
		i := strings.Index(local, "+")
		if i > 0 {
			local = local[:i]
		}
		email = strings.Replace(local, ".", "", -1) + "@" + local2host[1]
		fmt.Println(email)
		_, exists := set[email]
		if !exists {
			set[email] = struct{}{}
		}
	}
	return len(set)
}
