package services

import (
	"github.com/AH-dark/Anchor/pkg/conf"
	"strings"
)

func CheckGithubWhiteList(user string, repo string) bool {
	if len(conf.Config.Proxy.Github.WhiteList) == 0 {
		return true
	}

	for _, v := range conf.Config.Proxy.Github.WhiteList {
		t := strings.Split(v, "/")
		if len(t) != 2 {
			continue
		}

		if t[0] == "*" {
			return true
		}

		if t[0] == user {
			if t[1] == repo {
				return true
			} else if t[1] == "*" {
				return true
			}
		}
	}

	return false
}
