package webdav

import (
	"crypto/tls"
	"devcmd/internal/pkg"
	"fmt"
	"net/http"
	"os"
)

func UploadZip(url string, folderPath string, rFilePath string, username string, password string) {

	// Create a zip file
	filePath, err := pkg.ZipFolder(folderPath, rFilePath)
	if err != nil {
		panic(err)
	}
	defer os.Remove(filePath)

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

	//add zip extension
	if rFilePath[len(rFilePath)-4:] != ".zip" {
		rFilePath += ".zip"
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
		fmt.Println("zip file uploaded")
	} else {
		fmt.Println("zip file not uploaded")
	}

}
