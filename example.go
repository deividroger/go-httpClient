package main

import (
	"fmt"
	"io/ioutil"

	"github.com/deividroger/go-httpClient/gohttp"
)

var (
	httpclient = getGitHubClient()
)

func getGitHubClient() gohttp.Client {

	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		DisableTimeouts(true).
		SetConnectionTimeout(10).
		Build()

	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {

	response, err := httpclient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))

}

func createUser(user User) {

	response, err := httpclient.Post("https://api.github.com", nil, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))

}
