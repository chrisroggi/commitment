package main

import "regexp"

type Commit struct {
	Commit struct {
		Message string `json:"message"`
	} `json:"commit"`
}

func (commit *Commit) Words() []string {
	delimeter := regexp.MustCompile(` |\\n`)
	return delimeter.Split(commit.Commit.Message, -1)
}
