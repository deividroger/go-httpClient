package examples

import "fmt"

type Endpoints struct {
	CurrentUser       string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RespositoryUrl    string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {

	response, err := httpClient.Get("https://api.github.com", nil)

	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))

	var endpoints Endpoints

	if err := response.UnMarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repositories URL: %s", endpoints.RespositoryUrl))
	fmt.Println(fmt.Sprintf("Authorizations URL: %s", endpoints.AuthorizationsUrl))
	fmt.Println(fmt.Sprintf("CurrentUser: %s", endpoints.CurrentUser))

	return &endpoints, nil
}
