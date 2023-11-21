package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for i := 0; i < len(urls); i++ {
		req, err := http.NewRequest("GET", urls[i], nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create a new `http.Client` object.
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Response status code:", resp.StatusCode)
			return
		}

		filename := "image.jpg"
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		f.Close()

		fmt.Println("Image saved to", filename)

	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrently(urls []string) {
	var wg sync.WaitGroup

	downloadImage := func(url string) {
		defer wg.Done()

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Response status code:", resp.StatusCode)
			return
		}

		filename := "image.jpg"
		f, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Image saved to", filename)
	}

	for _, url := range urls {
		wg.Add(1)
		go downloadImage(url)
	}

	wg.Wait()
}

func main() {

	urls := []string{
		"https://unsplash.com/photos/a-view-of-the-top-of-a-mountain-in-the-clouds-wdQ7DUGJmk8",
		"https://unsplash.com/photos/a-man-in-a-kimono-walking-down-a-dark-alley-mHS1sTT2ybQ",
		"https://unsplash.com/photos/a-very-tall-building-with-lots-of-windows-d1h2rPLWfBg ",
		"https://unsplash.com/photos/a-leaf-is-laying-on-the-ground-in-the-middle-of-the-road-DIi8kbZZmKQ",
		"https://unsplash.com/photos/a-cup-of-hot-chocolate-with-marshmallows-in-it-s0Cnecr8W4U",
		// Add more image URLs
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrently(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// TODO: Implement download logic
	return nil
}
