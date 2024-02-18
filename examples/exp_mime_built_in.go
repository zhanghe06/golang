package main

import (
	"fmt"
	"mime"
	"strings"
)


func fixFileExt(filename, ext string) string {
	if ext == "" {
		return filename
	}
	fl := strings.Split(filename, ".")
	fExt := "." + fl[len(fl)-1]

	return strings.TrimSuffix(filename, fExt) + "." + strings.TrimPrefix(strings.ToLower(ext), ".")
}

func fixFileName(filename, contentType string) string {
	extension, _ := mime.ExtensionsByType(contentType)
	el := len(extension)
	if el == 0 {
		return filename
	}
	suffixExt := strings.ToLower(extension[el-1])

	fl := strings.Split(filename, ".")
	fileExt := "." + fl[len(fl)-1]

	return strings.TrimSuffix(filename, fileExt) + suffixExt
}

func guessMimeType(filename string)  {
	fl := strings.Split(filename, ".")
	ext := "." + fl[len(fl)-1]
	fmt.Println(mime.TypeByExtension(ext))
}

func main() {
	contentType := "text/plain"  // [,v .asc .txt]
	//contentType := "image/jpeg" // [.jpe .jpeg .jpg]
	extension, _ := mime.ExtensionsByType(contentType)
	fmt.Println(extension)

	contentType = "application/pdf"

	fmt.Printf("%s\n", fixFileName("Test.pdf", ""))
	fmt.Printf("%s\n", fixFileName("Test.pdf", contentType))
	fmt.Printf("%s\n", fixFileName("Test.Pdf", contentType))
	fmt.Printf("%s\n", fixFileName("Test.PDF", contentType))
	fmt.Printf("%s\n", fixFileName("Test", contentType))
	fmt.Printf("%s\n", fixFileName("Test.", contentType))


	fmt.Printf("%s\n", fixFileExt("123.txt", ""))
	fmt.Printf("%s\n", fixFileExt("123.txt", "pdf"))
	fmt.Printf("%s\n", fixFileExt("123.txt", ".pdf"))
	fmt.Printf("%s\n", fixFileExt("123", "pdf"))
	fmt.Printf("%s\n", fixFileExt("123", ".pdf"))
	fmt.Printf("%s\n", fixFileExt("123.", ".pdf"))


	guessMimeType("test.doc")
	guessMimeType("test.tif")
	guessMimeType("test.TIF")
}
