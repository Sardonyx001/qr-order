# QR Code Self Order System

## 概要

## TODO

## プロジェクト構成

.
├── README.md
├── backend 　　　　　　　　　　// API サーバ
│ ├── .env                  // 環境変数
│ ├── .env-example
│ ├── config　　　　　　　　　　// Echoサーバの各種config
│ │ ├── config.go
│ │ ├── db.go
│ │ └── http.go
│ ├── go.mod
│ ├── go.sum
│ ├── main.go                // APIサーバエントリポイント
│ └── start-db.sh 　　　　　　　// ローカル開発でのDBを構築するためのスクリプト
└── frontend 　　　　　　　　　　// Web アプリサーバ
  ├── README.md
  ├── app
  │ ├── globals.css
  │ ├── layout.tsx
  │ └── page.tsx
  ├── next-env.d.ts
  ├── next.config.mjs
  ├── package.json
  ├── pnpm-lock.yaml
  ├── postcss.config.js
  ├── public
  ├── tailwind.config.ts
  └── tsconfig.json
