## Start with docker
```
# git clone
cd /home 
git clone https://github.com/sanggi-wjg/docker-container-exporter.git

# docker build and run
docker build -t dc-exporter .

docker run -p 9091:9091 --privileged=true --name dc-exporter dc-exporter -v "/var/run/docker.sock":"/var/run/docker.sock"

# detached
docker run -d -p 9091:9091 --privileged=true --name dc-exporter dc-exporter -v "/var/run/docker.sock":"/var/run/docker.sock"
```
