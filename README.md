## Shinkan server
東京理科大学Web新歓用APIのレポジトリです

SwaggerEditor: <https://editor.swagger.io>  
API仕様書: [cmd/swagger.yaml](./cmd/swagger.yaml)

## Set up
### 環境変数の設定
```
$ export MYSQL_USER=root \
    MYSQL_PASSWORD=root \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=shinkan \
    SHINKAN_USER_NAME=hoge \
    SHINKAN_PASSWORD=hoge
```

### docker-composeの設定
```
$ cp docker-compose.sample.yaml docker-compose.yaml
```

### MySQLの起動
```
$ docker-compose up -d
```


## APIローカル起動方法
```
$ go run ./cmd/main.go
```

### ビルド方法
作成したAPIを実際にをサーバ上にデプロイする場合は、<br>
ビルドされたバイナリファイルを配置して起動することでデプロイを行います。
#### ローカルビルド
Macの場合
```
$ GOOS=linux GOARCH=amd64 go build -o dojo-api ./cmd/main.go
```
