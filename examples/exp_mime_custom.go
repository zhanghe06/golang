package main

import (
	"fmt"
	"path"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types

var mimeExt2TypeMap = map[string]string{
	".txt":  "text/plain",
	".log":  "text/plain",
	".sh":  "text/plain",
	".bat":  "text/plain",
	".csv":  "text/csv",
	".css":  "text/css",
	".ics":  "text/calendar",
	".htm":  "text/html",
	".html": "text/html",
	".pdf":  "application/pdf",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".gif":  "image/gif",
	".bmp":  "image/bmp",
	".svg":  "image/svg+xml",
	".ico":  "image/vnd.microsoft.icon",
	".tif":  "image/tiff",
	".tiff":  "image/tiff",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".xls":  "application/vnd.ms-excel",
	".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".ppt":  "application/vnd.ms-powerpoint",
	".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".rar":  "application/vnd.rar",
	".tar":  "application/x-tar",
	".zip":  "application/zip",
	".gz":   "application/gzip",
	".7z":   "application/x-7z-compressed",
	".jar":  "application/java-archive",
	".bin":  "application/octet-stream",
	".mp3":  "audio/mpeg",
	".wav": "audio/wav",
	".mp4":  "video/mp4",
	".mpeg": "video/mpeg",
}

var mimeType2ExtMap = map[string]string{
	"text/plain":               ".txt",
	"text/csv":                 ".csv",
	"text/css":                 ".css",
	"text/calendar":            ".ics",
	"text/html":                ".html",
	"application/pdf":          ".pdf",
	"image/png":                ".png",
	"image/jpeg":               ".jpg",
	"image/gif":                ".gif",
	"image/bmp":                ".bmp",
	"image/svg+xml":            ".svg",
	"image/vnd.microsoft.icon": ".ico",
	"image/tiff": ".tif",
	"application/msword":       ".doc",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",
	"application/vnd.ms-excel": ".xls",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         ".xlsx",
	"application/vnd.ms-powerpoint":                                             ".ppt",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
	"application/vnd.rar":         ".rar",
	"application/x-tar":           ".tar",
	"application/zip":             ".zip",
	"application/gzip":            ".gz",
	"application/x-7z-compressed": ".7z",
	"application/java-archive":    ".jar",
	"application/octet-stream":    ".bin",
	"audio/mpeg":                  ".mp3",
	"audio/wav":                  ".wav",
	"video/mp4":                   ".mp4",
	"video/mpeg":                  ".mpeg",
}

func getContentType(filename string) string {
	ext := path.Ext(filename)
	if contentType, ok := mimeExt2TypeMap[ext]; ok {
		return contentType
	} else {
		fmt.Println("key does not exist")
		return ""
	}
}

func getExtByType(contentType string) string {
	if ext, ok := mimeType2ExtMap[contentType]; ok {
		return ext
	} else {
		fmt.Println("key does not exist")
		return ""
	}
}

func main() {
	fmt.Println(getContentType("abc.doc"))
	fmt.Println(getExtByType("application/msword"))
}
