## TypeScriptとReact/Next.jsでつくるC2Cアプリケーション

![サンプルアプリ](https://private-user-images.githubusercontent.com/135807730/327265639-9cb3f71a-d407-4fad-869a-89d575bc2a06.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MTQ2MDQ0ODIsIm5iZiI6MTcxNDYwNDE4MiwicGF0aCI6Ii8xMzU4MDc3MzAvMzI3MjY1NjM5LTljYjNmNzFhLWQ0MDctNGZhZC04NjlhLTg5ZDU3NWJjMmEwNi5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjQwNTAxJTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI0MDUwMVQyMjU2MjJaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT05MWZlMTc3MmJkODgyOTkwZWNhNWYwODllMzIxYjdmYWM0NTNmMmYyNTRmNjQ4YmJlN2RlMmY0YTVhMTFmMzU2JlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZhY3Rvcl9pZD0wJmtleV9pZD0wJnJlcG9faWQ9MCJ9.qJCN16q1-p-7896Gu8C73Y_FgNjUjqhTBDgalsEPmA0)

## 環境

- Node.js: 16.14.0
- Next.js: 14.2.3
- React: 18.2.0

## インストール

```bash
npm install
```

## 環境変数の設定

.envを開く

```
API_BASE_URL=<バックエンドAPIへのベースURLの設定>
NEXT_PUBLIC_API_BASE_PATH=/api/proxy
```

## 開発サーバー起動

JSON Serverの設定と起動


こちらのリポジトリとは別のディレクトリで作業する

```
cd https://github.com/gihyo-book/ts-nextbook-json
cd ts-nextbook-json
npm ci
npm start
```

開発サーバーを起動し、http://localhost:3000/ にアクセス

```
npm run dev
```

## Storybook起動

Storybookを起動し、http://localhost:6006/ にアクセス

```
npm run storybook
```

## テスト実行

ユニットテスト実行

```
npm run test
```

## ソースコードリンター・フォーマッター

ソースコードをリンターでチェック

```
npm run lint
```

ソースコードをフォーマッターで整形

```
npm run format
```

## ディレクトリ構成

簡単にディレクトリ構成を以下に説明します。

```
├── .editorconfig
├── .env <-- 環境変数
├── .env.production <-- 本番用環境変数
├── .eslintrc.json <-- ESLint設定ファイル
├── README.md
├── jest.config.js <-- Jestの設定ファイル
├── jest.setup.js <-- テストファイルが実行される前に走る
├── next-env.d.ts
├── next.config.js <-- Next.js設定ファイル
├── package-lock.json
├── package.json
├── public
├── src
│   ├── components <-- Presentational Components
│   ├── containers <-- Container Compoments
│   ├── contexts <-- React Context
│   ├── pages <-- Next.jsのページ
│   ├── services <-- Web API Client
│   ├── themes <-- styled-componentsのテーマ
│   ├── types <-- 型定義
│   └── utils <-- 汎用関数
└── tsconfig.json
```

## 参考
「[TypeScriptとReact/Next.jsでつくる実践Webアプリケーション](https://gihyo.jp/book/2022/978-4-297-12916-3)」のサンプルアプリ
