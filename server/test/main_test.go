package test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test(t *testing.T) {
	t.Run(" ::/users", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8000/users")
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
		resp, err := http.Get("http://localhost:8000/names")
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
		resp, err := http.Get("http://localhost:8000/login?name=Oh%20Hanie")
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
		resp, err := http.Post("http://localhost:8000/loginuser",
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
		resp, err := http.Get("http://localhost:8000/loggedin?id=1")

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
