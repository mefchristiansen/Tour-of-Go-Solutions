// Modify the Crawl function to fetch URLs in parallel without fetching
// the same URL twice.

package main

import (
	"fmt"
	"sync"
)

// Cache struct to store already visisted urls
// This cache will be used to preven the crawler from visiting the same URLs
// This struct is thread safe by use of a mutex lock
type Cache struct {
	visited map[string]bool
	mux  sync.Mutex
}

// Visit a url and set its visited status to true (thread safe)
func (c *Cache) visit (url string) {
	c.mux.Lock()
	c.visited[url] = true
	c.mux.Unlock()
}

// Get a url's visited status (thread safe)
func (c *Cache) get (url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.visited[url]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache Cache,
		wg *sync.WaitGroup) {
	defer wg.Done()
	
	// If maximum depth reached, or url has already been visited, return
	if depth <= 0 || cache.get(url) {
		return
	}
	
	// Cache the current url to prevent future visits
	cache.visit(url)
	
	body, urls, err := fetcher.Fetch(url)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	
	for _, u := range urls {
		wg.Add(1)
		// Execute Crawl goroutines in parallel
		go Crawl(u, depth-1, fetcher, cache, wg)
	}
				
	return
}

func main() {
	// Use wait group to ensure all goroutines are have finished before exiting
	wg := &sync.WaitGroup{}

	wg.Add(1)
	
	go Crawl("https://golang.org/",
		  4,
		  fetcher,
		  Cache{visited:make(map[string]bool)},
		  wg,
	)
	
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}