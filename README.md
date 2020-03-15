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

## Usage
### run
```
$ make docker/run
```

### seed
```
$ make docker/mysql/seed
```

### down
```
$ make docker/down
```
