# limiting-server
請實作一個 server，每個 IP 每分鐘僅能接受 60 個 requests

## 題目
* 每個 IP 每分鐘僅能接受 60 個 requests
* 在 response 顯示目前的 request 量,超過限制的話則顯示 “Error”,例如在一分鐘內第
30 個 request 則顯示 30,第 61 個 request 則顯示 Error
* 可以使用任意資料庫,也可以自行設計 in-memory 資料結構,並在文件中說明理由

## 可以使用任意資料庫,也可以自行設計 in-memory 資料結構,並在文件中說明理由
此作業使用server.Limiter類別的實例儲存與運算資料。
如需考慮到server重啟時要載入原先已接受的request數量，
或需將request紀錄備查。則須考慮使用資料庫存放。

## 本機端測試方式
1. 先確認port 1234沒有被防火牆阻攔，或者沒有被其他程序佔用。
2. 取得本專案。 `git clone https://github.com/shodawu/limiting-server.git`
3. 進入專案資料夾。`cd ~/limiting-server`
4. 閱讀專案執行參數說明 `go run main.go -h` 
5. 於不同的兩個命令視窗執行client、server
```
go run main.go -s=s
go run main.go -s=c
```