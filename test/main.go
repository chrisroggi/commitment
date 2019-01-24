package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// type Commit struct {
// 	Commit struct {
// 		Message string `json:"message"`
// 	} `json:"commit"`
// }

// type Output struct {
// 	Word  string
// 	Count int
// }

func main() {
	var repo = "bedrocketjmd/ruby-unimatrix-sdk"
	var commitsURL = "https://api.github.com/repos/" + repo + "/commits?per_page=100"
	var wordMap = make(map[string]int)
	var terms []Term
	var commits []Commit

	delimeter := regexp.MustCompile(` |\n`)
	resp, err := http.Get(commitsURL)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &commits)

	curTime := time.Now()

	for _, commit := range commits {
		var words = delimeter.Split(commit.Commit.Message, -1)

		for _, word := range words {
			if wordMap[word] == 0 {
				wordMap[word] = 1
			} else {
				wordMap[word]++
			}
		}
	}

	var index int

	for word, count := range wordMap {
		if index != 0 {
			var i = index - 1

			if count > output[i].Count {
				if i != 0 {
					for count > output[i-1].Count {
						i--
						if i == 0 {
							break
						}
					}
				}

				output = append(output, Output{})
				copy(output[i+1:], output[i:])
				output[i] = Output{Count: count, Word: word}
			} else {
				output = append(output, Output{Count: count, Word: word})
			}
		} else {
			output = append(output, Output{Count: count, Word: word})
		}

		index++
	}

	fmt.Printf("%v", (time.Now().UnixNano()-curTime.UnixNano())/10000)
	fmt.Printf("ms\n\n")

	fmt.Printf("%v", output)
}
