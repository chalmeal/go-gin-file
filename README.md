**go-gin-file**

Golang + Gin + S3(MinIO)

## はじめに
AWS(S3)を想定したファイルアップロード/ダウンロード実装です。MinIOを利用し、ローカル環境上でのストレージ利用を可能とします。

## 構成
```
├── main.go
├── docker-compose.yml
├── common
|    ├── aws.go
|    ├── db.go
|    ├── response.go
|    └── utils.go
├── config
|    ├── app.ini
|    └── config.go
├── app
|     ├── models
|     |     ├── files.go
|     |     └── users.go
|     └── api
|　　　　　　├── controllers.go 
|　　　　　　├── file_controllers.go 
|　　　　　　└── user_controllers.go
└── routes
      └── api.go
```
