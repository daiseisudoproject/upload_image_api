package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// リクエストからファイルを解析
	r.ParseMultipartForm(10 << 20) // 最大10MBのファイルまで扱う
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error Retrieving the File: %v", err)
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if !strings.HasPrefix(handler.Header.Get("Content-Type"), "image/") {
		http.Error(w, "Only image uploads are allowed", http.StatusBadRequest)
		return
	}

	tempFile, err := os.CreateTemp("uploads", "upload-*.png")
	if err != nil {
		log.Printf("Error Creating the File: %v", err)
		http.Error(w, "Failed to create a temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error Reading the File: %v", err)
		http.Error(w, "Failed to read the file content", http.StatusInternalServerError)
		return
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File: %s\n", handler.Filename)
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	// uploadsディレクトリが存在するか確認し、なければ作成
	if _, err := os.ReadDir("uploads"); err != nil {
		fmt.Println("Uploads directory is missing: ", err)
		if err := os.Mkdir("uploads", 0755); err != nil {
			log.Fatal("Failed to create uploads directory: ", err)
		}
	}

	// ハンドラの設定
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		uploadFile(w, r)
	})

	// サーバの起動
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
