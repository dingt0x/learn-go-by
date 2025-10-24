package concurrency

import (
    "reflect"
    "testing"
    "time"
)

func mockWebsiteChecker(url string) bool {
    return url != "waat://furhurterwe.geds"
}

func TestCheckWebsite(t *testing.T) {

    websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
    }

    want := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }

    got := CheckWebsite(mockWebsiteChecker, websites)

    if !reflect.DeepEqual(want, got) {
        t.Fatalf("wanted %v, got %v", want, got)
    }
}

func slowStubWebChecker(_ string) bool {
    time.Sleep(time.Millisecond * 20)
    return true
}

func BenchmarkCheckWebsite(b *testing.B) {
    urls := make([]string, 10)

    for i := 0; i < len(urls); i++ {
        urls[i] = "a url"
    }
    for i := 0; i < b.N; i++ {
        CheckWebsite(slowStubWebChecker, urls)
    }
}
