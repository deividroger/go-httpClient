package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/deividroger/go-httpClient/gohttp"
)

var (
	httpclient = getGitHubClient()
)

func getGitHubClient() gohttp.HttpClient {
	client := gohttp.New()
	commonHeaders := make(http.Header)
	//commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)
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
