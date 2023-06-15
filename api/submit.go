package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create a new folder if folder does not exist
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the folder was created successfully.
	if _, err := os.Stat("uploads"); err == nil {
		fmt.Println("Folder created successfully.")
	} else {
		fmt.Println("Error creating folder:", err)
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(filepath.Join("uploads", fileHeader.Filename)) //(fileHeader.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload successful")
}
