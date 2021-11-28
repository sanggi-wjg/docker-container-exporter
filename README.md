## Start with docker
```
cd /home 
git clone https://github.com/sanggi-wjg/docker-container-client.git

docker build -t dc-client .
docker run -d -p 9091:9091 --privileged=true --name dc-client dc-client
```



### Go mod
```
go mod init github.com/sanggi-wjg/docker-continer-client        

# 추가사항
go mod vendor
# 모듈 다운로드
go mod download
# verify
go mod verify
```

### Ref
```
# Go Docker client api
https://pkg.go.dev/github.com/docker/docker/client#section-readme

# Make Dockerfile 
https://docs.docker.com/language/golang/build-images/

# install go on ubuntu 20.04
https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04
```
