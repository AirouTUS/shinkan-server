## Shinkan server
東京理科大学Web新歓用APIのレポジトリです

SwaggerEditor: <https://editor.swagger.io>  
API仕様書: [cmd/swagger.yaml](./cmd/swagger.yaml)

## Set up

```
$ cp docker-compose.sample.yaml docker-compose.yaml
```
portなどは各自で適当に設定してください
## Usage
### run
```
$ make docker/run
```

### down
```
$ make docker/down
```

### test
```
$ make test
```
