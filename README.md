# contextWorkshop
Go context 演示，透過 context 機制，在服務和客戶端實作，有效降低 TCP `TIME_WAIT` 連線。

# 說明
*  `srv.go` 是服務端
*  `client.go` 是客戶端
*  `netstat.sh` 是一個觀察腳本，每秒更新一次，觀察服務端所有連線。

# 使用方法

```sh
sh netstat.sh
```

```sh
go run srv.go
```
```sh
go run client.go
```