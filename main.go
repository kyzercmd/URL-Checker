package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func checkUrlAndSave(urlStr string, wg *sync.WaitGroup) {
	response, err := http.Get(urlStr)
	if err != nil {
		fmt.Println(err)
	}else {
		defer response.Body.Close()
		fmt.Printf("%s -> status code: %d\n", urlStr, response.StatusCode)

		if response.StatusCode == http.StatusOK {
			responseSlice, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Println("Error Reading body")
			}

			u, err := url.Parse(urlStr)
			fileName := u.Host + ".html"
			err = os.WriteFile(fileName, responseSlice, 0644)
			if err != nil {
				log.Fatal(err)
			}else {
				fmt.Println("Page saved to", fileName)
			}
		}
	}
	wg.Done()
}

func main() {
	urls := []string {"https://google.com", "https://youtube.com", "https://githhub.org", "https://spotify.com/kling"}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, value := range urls {
		go checkUrlAndSave(value, &wg)
		
	}

	wg.Wait()
}