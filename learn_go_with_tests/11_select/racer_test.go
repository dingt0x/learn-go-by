package racer

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

// func SlowFunc(w http.ResponseWriter, r *http.Request) {
//     time.Sleep(time.Millisecond * 10)
//     w.WriteHeader(http.StatusOK)
// }
//
// func TestHttp(t *testing.T) {
//     Server := httptest.NewServer(http.HandlerFunc(SlowFunc))
//     fmt.Printf("%+v", Server)
// }
func TestRacer(t *testing.T) {
    slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(time.Millisecond * 10)
        w.WriteHeader(http.StatusOK)
    }))
    fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))

    slowURL := slowServer.URL
    fastURL := fastServer.URL

    want := fastURL
    got := Racer(slowURL, fastURL)

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }

}
