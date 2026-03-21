//TODO: Rob Pike gave a great talk called Concurrency is not parallelism
// Now here is my current logic for my web health checker app, answer my questions to raise my go level up then optimize the logic:
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"sync"
)

//& collect the domains -> hit that domain (ask for the header for performance)
func main() {

	//~ Data captcher
	var urls []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter URLs (press Enter on an empty line to finish):")
	for scanner.Scan() { //? How I can loop on a scanner.Scan() function while it returns a boolean value, is these even a valid loop
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		urls = append(urls, line)
	}
	// err := scanner.Err()  //? what is the difference between declaring the var out side the if and inside it, which one cleaner
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error while reading your inputs : ", err) //? why here i use Fprintln, and what does that signature mean: func fmt.Fprintln(w io.Writer, a ...any) (n int, err error)
	}

	//~ executing the logic
	client := &http.Client{ //! "&" here and "*" below to prevent a copying the client struct, serve the performance and offering the full controlle
		Timeout: 5 * time.Second, //! timeout prevents our program from hanging forever if the site is unresponsive
	}

	fmt.Println("Starting Website Health Check...")
	fmt.Println("-------------------------------")

	// without goroutines, if we check 10 websites and each takes 2 seconds, the user waits 20 seconds. 
		// for _, url := range urls {
		// 	checkUrl(client, url)
		// }

	//~ but with it, the user only waits ~2 seconds total, because we check all 10 at the exact same time (Concurrent).
		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			go func(u string){ //? why that error when I did not passed the u string to the func : too many arguments in call to (func() literal) have (string) want ()compilerWrongArgCount
				defer wg.Done()
				checkUrl(client, u)
			}(url)
		}
		wg.Wait()
}

func checkUrl(client *http.Client, url string) {
	res, err := client.Head(url) //! here using the Head over the Get to rais the performance
	if err != nil {
		fmt.Printf("[DOWN] %s - Error: %v\n", url, err) 
		return
	}
	defer res.Body.Close()
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		fmt.Printf("[ UP ] %s - Status: %d\n", url, res.StatusCode)
	} else {
		fmt.Printf("[WARN] %s - Status: %d\n", url, res.StatusCode)
	}
	fmt.Printf("[WARN] %s - Status: %d\n", url, res.StatusCode)
}
