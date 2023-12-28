package pkg

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipFolder Compresses a specified folder and all its subfolders into a zip file.
func ZipFolder(source, target string) (string, error) {

	zipfile, err := os.CreateTemp("", target+"-davcmd-*.zip")

	if err != nil {
		return "", err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// Anonymous function to handle recursively files and folders
	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Make sure the file path in the zip file maintains the folder structure
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return err
	})

	return zipfile.Name(), err
}

// ZipFile Compresses a specified file into a zip file.
func ZipFile(source, target string) (string, error) {

	zipfile, err := os.CreateTemp("", target+"-davcmd-*.zip")

	if err != nil {
		return "", err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	file, err := os.Open(source)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Get the file information
	info, err := file.Stat()
	if err != nil {
		return "", err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return "", err
	}

	// Make sure the file path in the zip file maintains the folder structure
	header.Name, err = filepath.Rel(filepath.Dir(source), source)
	if err != nil {
		return "", err
	}

	header.Method = zip.Deflate

	writer, err := archive.CreateHeader(header)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		return "", err
	}

	return zipfile.Name(), err
}
