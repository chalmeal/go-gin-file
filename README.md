**go-gin-file**

Golang + Gin + MiniO

## はじめに
S3を想定したファイルダウンロード/アップロード実装です。MiniOを利用し、ローカル環境上でのストレージ利用を可能とします。

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
## ドキュメント
### データ定義
| テーブル名                                                                         | 概要                                                 |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------- |
| [files](https://github.com/chalmeal/go-gin-file/blob/master/.doc/data/mails.md)    | ファイルテーブル。アップロードされたファイルの情報。 |
| [users](https://github.com/chalmeal/go-gin-file/blob/master/.doc/data/historys.md) | ユーザーテーブル。サービスの利用者情報。             |

### 仕様
| 書名                                                                                    | 概要                                      |
| --------------------------------------------------------------------------------------- | ----------------------------------------- |
| [ファイル利用](https://github.com/chalmeal/go-gin-file/blob/master/.doc/method/file.md) | ファイルアップロード/ダウンロードについて |

## セットアップ

### DB
* DBの環境は以下を想定します。
  * MySQL
  * GORM
* create schemaのみ行う必要があります。[DDL](.db/setup/ddl-create-chema.sql)
* [app.ini](config/app.ini)に対してDB接続情報を定義してください。
* テーブルはGORMが提供するAutoMigrateを利用します。
  * 各テーブルは初回API実行時に生成されます。
  * 本プロジェクトで定義されている詳細なデータ定義に関しては、各ドメインのdocを参照してください。
    * ファイル：[files](https://github.com/chalmeal/go-gin-file/blob/master/.doc/data/files.md)
    * ユーザー：[users](https://github.com/chalmeal/go-gin-file/blob/master/.doc/data/users.md)

### Docker
MinIOの利用にDockerを利用しています。Dockerの導入を行ってください。

## アプリケーションスタート
アプリケーションのスタートはデバッガを推奨しています。

Run and Debugの`Run go-echo-file`から実行してください。