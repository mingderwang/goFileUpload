package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("--- File Upload Endpoint Hit ---")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        log.Fatal("Error Retrieving the File")
        log.Fatal(err)
        return
    }
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    dir, err := ioutil.TempDir("", "temp-images")
	  if err != nil {
		    log.Fatal(err)
	  }
    tempFile, err := ioutil.TempFile(dir, "upload-*.png")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("--- Save file to fileName: ---")
    fmt.Println(tempFile.Name())


    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    defer file.Close()
    defer tempFile.Close()
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "上載成功\n")
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":8081", nil)
}

func main() {
    fmt.Println("a maximum upload of 10 MB files...")
    setupRoutes()
}
