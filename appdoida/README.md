# How to appdoida
## CentOS Linux 7.7.1908 

``` 
git clone https://github.com/tbernacchi/my-first-cli-using-go.git
cd appdoida/
mkdir -p $GOPATH/src/appdoida
cp -pr * $GOPATH/src/appdoida
go mod init github.com/my-first-cli-using-go/appdoida
GOOS=linux GOARCH=amd64 go build -o /bin/appdoida
```

