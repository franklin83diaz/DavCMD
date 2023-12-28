package webdav

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, username, password, outputPath string) {

	// Create a HTTP client that ignores certificate errors, similar to the -k option in curl
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//add /download to url
	url += "/download"

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Set the 'X-Requested-With' header
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// Set basic auth
	req.SetBasicAuth(username, password)

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check that the request was successful
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	print(resp.StatusCode)

	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Copy the response to the file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		panic(err)
	}

}
