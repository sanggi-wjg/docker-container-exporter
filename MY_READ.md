### Go mod
```
go mod init github.com/sanggi-wjg/docker-continer-exporter

# 추가사항
go mod vendor
# 모듈 다운로드
go mod download
# verify
go mod verify
```

### Ref
```
# Make Dockerfile 
https://docs.docker.com/language/golang/build-images/

# install go on ubuntu 20.04
https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04

# Go Docker client api
https://pkg.go.dev/github.com/docker/docker/client#section-readme

# Custom Exporter
https://reachmnadeem.wordpress.com/2021/01/06/writing-prometheus-exporter-using-go/
https://developpaper.com/implementation-of-prometheus-custom-exporter/
```
