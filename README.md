# Image Upload API

このAPIは画像を受け取り、サーバーのローカルディレクトリに保存する機能を提供します。画像は`uploads`ディレクトリに一時ファイルとして保存されます。

## 始め方

以下の手順に従ってローカル環境にAPIを設定してください。

### 前提条件

- Go言語がインストールされていること
- ポート8080が空いていること

### インストール

このAPIには外部ライブラリは必要ありませんが、Goがシステムにインストールされている必要があります。Goのインストール方法は[こちら](https://golang.org/doc/install)を参照してください。

### 実行

リポジトリをクローンし、以下のコマンドを実行してサーバーを起動します：

```bash
go run main.go
```

### APIの使用方法

画像をアップロードするには、以下のエンドポイントに対してPOSTリクエストを送信します：

```bash
POST /upload
```

リクエストはmultipart/form-data形式でなければなりません。以下はcurlを使用した例です：

```bash
curl -F "file=@path_to_your_image.jpg" http://localhost:8080/upload
```

成功すると、以下のようなJSONレスポンスが返されます：

```bash
{
  "message": "Successfully Uploaded File: filename.jpg"
}
```

### エラーハンドリング

ファイルが画像でない場合や、サーバーでの処理中にエラーが発生した場合、適切なHTTPステータスコードとエラーメッセージが返されます。

### コントリビューション

このプロジェクトはオープンソースであり、改善のためのプルリクエストを歓迎します。小さなタイポから新機能の提案まで、どんな貢献も価値があります。
