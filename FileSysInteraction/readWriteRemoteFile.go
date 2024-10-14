package main

import (
	"fmt"
	"http"
	"io"
	"os"
)

func main() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://myserver.mydomain/myfile", nil)

	if err != nil {
    	return err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	cancel()
	if err != nil {
    	return err
	}

	// resp contains an io.ReadCloser that we can read as a file.
	// Let's use io.ReadAll() to read the entire content to data.
	data, err := io.ReadAll(resp.Body)

	fmt.Println(string(data))

	flags := os.O_CREATE|os.O_WRONLY|os.O_TRUNC
	f, err := os.OpenFile("path/to/file", flags, 0644)

	if err != nil {
    	return err
	}
	defer f.Close()

	if bytes, err := io.Copy(f, resp.Body); err != nil {
    	return err
	} else {
		fmt.Println("Copied bytes", string(bytes))
	}

}

 