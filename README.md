# go-test
Testing Go programming language

## Generate tls certificate
Mac:
```
go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost
```
Windows:
```
go run %GOROOT%/src/crypto/tls/generate_cert.go --host=localhost
```