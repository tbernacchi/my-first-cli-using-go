# appdoida

## Requirements 
* CentOS Linux 7.7.1908 
* go version go1.15.3 linux/amd64

## Usage

``` 
git clone https://github.com/tbernacchi/my-first-cli-using-go.git
cd appdoida/
mkdir -p $GOPATH/src/appdoida
cp -pr * $GOPATH/src/appdoida
go mod init github.com/my-first-cli-using-go/appdoida
GOOS=linux GOARCH=amd64 go build -o /bin/appdoida
```
Voilá!

## Author

👤 **Tadeu Bernacchi**

