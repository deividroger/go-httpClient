package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/deividroger/go-httpClient/gohttp"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package examples")
	gohttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {

		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")

			if err.Error() != "timeout getting github endpoints" {
				t.Error("invalid error message received")
			}

		}
	})

	t.Run("TestErrorUnmarshalFromBody", func(t *testing.T) {

		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})

		endpoints, err := GetEndpoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		} else {
			if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
				t.Error("invalid error message received")
			}
		}

	})

	t.Run("TestNoError", func(t *testing.T) {

		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `"{current_user_url": "https://api.github.com/user"}`,
		})

		endpoints, err := GetEndpoints()

		if err != nil {
			t.Error(fmt.Sprintf("no error was expected and we got '%s'", err.Error()))
		}

		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		} else {
			if endpoints.CurrentUserUrl != "https://api.github.com/user" {
				t.Error("invalid current user url")
			}
		}
	})
}
