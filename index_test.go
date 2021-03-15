package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Turns out you NEED to capitalize Test for this to get read
func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		fmt.Println("got, want", got, want)

		if got != want {
			t.Error("got, want", got, want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		fmt.Println("got, want", got, want)

		if got != want {
			t.Error("got, want", got, want)
		}
	})
}

// func testGETScores(t *testing.T) {
// 	t.Run("returns the score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.)
// 	})
// }
