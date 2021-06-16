package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Maps do not like it whenever more than one thing tries to write to
	// them at once. This is a race condition, a bug that occurs when the
	// output of our software is dependent on the timing and sequence of
	// events that we have no control over. Test with `go test -race`
	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// Channels are a Go data structure that can both receive and send values.
	// These operations allow communication between different processes.
	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
