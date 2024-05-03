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
	// リクエストからファイルを解析
	r.ParseMultipartForm(10 << 20) // 最大10MBのファイルまで扱う

	// フォームから "file" というキーでファイルを取得
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// MIMEタイプを確認
	if !strings.HasPrefix(handler.Header.Get("Content-Type"), "image/") {
		http.Error(w, "Only image uploads are allowed", http.StatusBadRequest)
		return
	}

	// テンポラリディレクトリにファイルを保存
	tempFile, err := os.CreateTemp("uploads", "upload-*.png")
	if err != nil {
		fmt.Println("Error Creating the File")
		fmt.Println(err)
		http.Error(w, "Failed to create a temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// ファイルの内容を読み取って保存
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error Reading the File")
		fmt.Println(err)
		http.Error(w, "Failed to read the file content", http.StatusInternalServerError)
		return
	}
	tempFile.Write(fileBytes)

	// 成功のレスポンスを返す
	fmt.Fprintf(w, "Successfully Uploaded File: %s\n", handler.Filename)
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
	http.HandleFunc("/upload", uploadFile)

	// サーバの起動
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
