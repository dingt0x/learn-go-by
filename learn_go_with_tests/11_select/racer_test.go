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

func makeDelayedServer(delay time.Duration) *httptest.Server {
    slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))

    return slowServer

}
func TestRacer(t *testing.T) {
    t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {

        slowServer := makeDelayedServer(time.Millisecond * 20)
        defer slowServer.Close()
        fastServer := makeDelayedServer(time.Millisecond * 1)
        defer fastServer.Close()

        slowURL := slowServer.URL
        fastURL := fastServer.URL

        want := fastURL
        got, _ := Racer(slowURL, fastURL)

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })
    t.Run("return an error if a server doesn't respond with the specified time", func(t *testing.T) {
        server := makeDelayedServer(20 * time.Millisecond)
        defer server.Close()

        _, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)
        if err == nil {
            t.Error("except an error but didn't get one.")
        }

    })

}
