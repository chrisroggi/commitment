package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func Commits(repo string) []Commit {
	repo = "bedrocketjmd/ruby-unimatrix-sdk"

	var commitsURL = "https://api.github.com/repos/" + repo + "/commits?per_page=100"
	var commits []Commit

	resp, err := http.Get(commitsURL)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &commits)

	return commits
}

func CommitWordMap() map[string]int {
	var wordMap = make(map[string]int)
	for _, commit := range Commits("test") {
		for _, word := range commit.Words() {
			word = strings.ToLower(word)
			if wordMap[word] == 0 {
				wordMap[word] = 1
			} else {
				wordMap[word]++
			}
		}
	}

	return wordMap
}

func SortedCommitTerms() []Term {
	var terms []Term
	var index int

	for word, count := range CommitWordMap() {
		if index != 0 {
			var i = index - 1

			if count > terms[i].Count {
				if i != 0 {
					for count > terms[i-1].Count {
						i--
						if i == 0 {
							break
						}
					}
				}

				terms = append(terms, Term{})
				copy(terms[i+1:], terms[i:])
				terms[i] = Term{Count: count, Value: word}
			} else {
				terms = append(terms, Term{Count: count, Value: word})
			}
		} else {
			terms = append(terms, Term{Count: count, Value: word})
		}

		index++
	}

	return terms
}
