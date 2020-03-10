## aurora02專案api
golang撰寫
> ### 內容
> 
> 呼叫狀態
> 
> 錄音檔下載
> 
> 錄音檔打包下載

### 本地開發
```
yarn dev
or
go run main.go
```
### 打包windows exe
```
yarn build
or
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```