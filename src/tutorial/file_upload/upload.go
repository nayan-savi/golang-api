package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func readFile() (string, error) {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	}
	return string(data), err
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	//fmt.Fprintf(w, "<!DOCTYPE html> <html lang=\"en\"> <head> <meta charset=\"UTF-8\" />     <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />     <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\" />     <title>Document</title>   </head>   <body>     <form       enctype=\"multipart/form-data\"       action=\"http://localhost:8080/upload\"       method=\"post\"     >       <input type=\"file\" name=\"myFile\" />       <input type=\"submit\" value=\"upload\" />     </form>   </body> </html>")
	da, _ := readFile()
	fmt.Fprintf(w, da)
    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":8080", nil)
}

func main() {
    fmt.Println("Hello World")
    setupRoutes()
}