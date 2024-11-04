package test

import (
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
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
}

func TestGetUsersName(t *testing.T) {
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
}

func TestLoginbyparam(t *testing.T) {
	resp, err := http.Get("http://localhost:8000/login2?name=Oh%20Hanie")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}

func TestLoginbybody(t *testing.T) {
	resp, err := http.Post("http://localhost:8000/login",
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
}
