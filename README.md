# go-sample-api
勉強用リポジトリです。

# コンテナ操作
## 起動
```bash
make up
```

## 停止
```bash
make stop
```

## 破棄
```bash
make down
```

## 起動中確認
```bash
make ps
```

`make ps`で以下コンテナが立ち上がっていることを確認
```bash
CONTAINER ID   IMAGE                          COMMAND                  CREATED             STATUS         PORTS                                        NAMES
df76eb5b5196   minio/mc                       "/bin/sh -c 'until (…"   3 seconds ago       Up 1 second                                                 go-sample-create-bucket
cc0fd1e443d3   swaggerapi/swagger-ui:latest   "/docker-entrypoint.…"   3 seconds ago       Up 2 seconds   80/tcp, 0.0.0.0:8080->8080/tcp               go-sample-openapi
16b086be4e94   minio/minio                    "bash -c '/opt/bin/m…"   3 seconds ago       Up 2 seconds   9000/tcp, 0.0.0.0:9001-9002->9001-9002/tcp   go-sample-minio
685d3deb18c6   redis:6.2-buster               "docker-entrypoint.s…"   3 seconds ago       Up 2 seconds   0.0.0.0:6379->6379/tcp                       go-sample-redis
eb674de517ed   go-sample_mysql                "docker-entrypoint.s…"   30 seconds ago      Up 2 seconds   0.0.0.0:3306->3306/tcp, 33060/tcp            go-sample-mysql
43160562761a   go-sample_api                  "bash"                   2 minutes ago       Up 2 seconds   0.0.0.0:8000->8000/tcp                       go-sample-api
fc6e46d1e48f   go-sample_nginx                "/docker-entrypoint.…"   About an hour ago   Up 2 seconds   0.0.0.0:80->80/tcp                           go-sample-nginx

```
