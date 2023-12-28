package webdav

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func Upload(url string, filePath string, rFilePath string, username string, password string) {

	// Open the file you want to upload
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Create a HTTP client that ignores certificate errors, similar to the -k option in curl
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Add '/' to url if not exist
	if url[len(url)-1] != '/' {
		url += "/"
	}
	//Create request PUT
	req, err := http.NewRequest("PUT", url+rFilePath, file)
	if err != nil {
		fmt.Println("Error to create resquet:", err)
		return
	}

	// set header 'X-Requested-With'
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// set basic auth
	req.SetBasicAuth(username, password)

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		fmt.Println("file uploaded")
	} else {
		fmt.Println("file not uploaded")
	}

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error al leer la respuesta:", err)
	// 	return
	// }
	// fmt.Println("Respuesta del servidor:", string(body))
}
