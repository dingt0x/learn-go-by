package concurrency

type WebsiteChecker func(string) bool
type result struct {
    string
    bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
    res := make(map[string]bool)
    resChannel := make(chan result)

    for _, url := range urls {
        go func() {
            resChannel <- result{url, wc(url)}
        }()
    }
    for i := 0; i < len(urls); i++ {
        r := <-resChannel
        res[r.string] = r.bool
    }

    return res
}
