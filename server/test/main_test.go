package test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

const (
	ip   string = "172.17.0.2"
	port string = "8000"
)

func Test(t *testing.T) {
	t.Run(" ::/users", func(t *testing.T) {

		resp, err := http.Get(fmt.Sprintf("http://%s:%s/users", ip, port))
		if err != nil {
			t.Fatal(err)
		}

		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run(" ::/names", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("http://%s:%s/names", ip, port))
		if err != nil {
			t.Fatal(err)
		}

		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run(" ::/login", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("http://%s:%s/login?name=Oh%%20Hanie", ip, port))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run(" ::/loginuser", func(t *testing.T) {
		resp, err := http.Post(fmt.Sprintf("http://%s:%s/loginuser", ip, port),
			"application/json",
			strings.NewReader(`{"name":"Oh Hanie"}`))

		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run(" ::/loggedin", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("http://%s:%s/loggedin?id=1", ip, port))

		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)

		if err != nil {
			t.Fatal(err)
		}
		/*
			var data map[string]interface{}
			err = json.Unmarshal(body, &data)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, data["data"], "user is logged in")
		*/
		assert.Equal(t, 200, resp.StatusCode)
	})
}
