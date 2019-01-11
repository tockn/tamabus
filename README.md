# tamabus

## 構成
DBはmysql5.6。dockerで動かす
APIサーバーはGo製

## 環境構築
mysqlコンテナ立ち上げて、dep使ってGoのライブラリ依存関係を解決します。
```
$ make docker_up
$ make deps
```

起動
```
$ make run
```

