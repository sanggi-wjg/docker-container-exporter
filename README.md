# Docker-Container-Exporter
Exporter for docker container state.
Prometheus exporter for docker container state, written in Go.

This exporter only export containers state with name

## Install
docker-container-exporter listen on port 9091 by default.
### docker install
```
# git clone
cd /home 
git clone https://github.com/sanggi-wjg/docker-container-exporter.git

# docker build and run
docker build -t dc-exporter .

docker run -p 9091:9091 -v "/var/run/docker.sock":"/var/run/docker.sock" --privileged=true --name dc-exporter dc-exporter 

# detached
docker run -d -p 9091:9091 -v "/var/run/docker.sock":"/var/run/docker.sock" --privileged=true --name dc-exporter dc-exporter 
```

```sh
root@ubuntu20:/home/docker-container-exporter# docker build -t dc-exporter .
Sending build context to Docker daemon  2.527MB
Step 1/9 : FROM golang:1.16-alpine
 ---> eee5af307da8
Step 2/9 : EXPOSE  9091
 ---> Using cache
 ---> 0a765e1707ee
Step 3/9 : RUN     mkdir /app
 ---> Using cache
 ---> 4328f745c56d
Step 4/9 : WORKDIR /app
 ---> Using cache
 ---> 532bae7c98bf
Step 5/9 : COPY    . ./
 ---> 0f653b6f4651
Step 6/9 : RUN     go mod download
 ---> Running in c2bac11bdd25
Removing intermediate container c2bac11bdd25
 ---> aecdf566d60b
Step 7/9 : RUN     go mod verify
 ---> Running in ec1331cc01ff
all modules verified
Removing intermediate container ec1331cc01ff
 ---> 453ba6f5f3de
Step 8/9 : RUN     go build -o /docker_container_exporter
 ---> Running in f2842a89972c
Removing intermediate container f2842a89972c
 ---> aa58f65d641f
Step 9/9 : CMD     [ "/docker_container_exporter" ]
 ---> Running in f09d0ac008d3
Removing intermediate container f09d0ac008d3
 ---> d74e2c2e02aa
Successfully built d74e2c2e02aa
Successfully tagged dc-exporter:latest

root@ubuntu20:/home/docker-container-exporter# docker run -d -p 9091:9091 -v "/var/run/docker.sock":"/var/run/docker.sock" --privileged=true --name dc-exporter dc-exporter
fef9fe8ec397124820b2a2548773597315d17a34927dbff4f4f14285e8fa0198

root@ubuntu20:/home/docker-container-exporter# docker ps -a
CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS                      PORTS                                                 NAMES
fef9fe8ec397   dc-exporter            "/docker_container_e…"   5 seconds ago    Up 5 seconds                0.0.0.0:9091->9091/tcp, :::9091->9091/tcp             dc-exporter
```


## Metrics
* app_docker_container_create_time{name="container image name"}
* app_docker_container_up{name="container image name"}
* app_docker_container_up_time{name="container image name"} (not completed)

```
root@ubuntu20:/home/docker-container-exporter# curl localhost:9091/metrics

# HELP app_docker_container_create_time docker container create time
# TYPE app_docker_container_create_time counter
app_docker_container_create_time{names="airflow_airflow-scheduler_1"} 1.637574017e+09
app_docker_container_create_time{names="airflow_airflow-triggerer_1"} 1.637574017e+09
app_docker_container_create_time{names="airflow_airflow-webserver_1"} 1.637574017e+09
app_docker_container_create_time{names="airflow_airflow-worker_1"} 1.637574017e+09
app_docker_container_create_time{names="airflow_flower_1"} 1.637574017e+09
app_docker_container_create_time{names="airflow_postgres_1"} 1.637573956e+09
app_docker_container_create_time{names="airflow_redis_1"} 1.637573956e+09
app_docker_container_create_time{names="dc-exporter"} 1.638281884e+09
# HELP app_docker_container_up docker container up
# TYPE app_docker_container_up counter
app_docker_container_up{names="airflow_airflow-scheduler_1"} 1
app_docker_container_up{names="airflow_airflow-triggerer_1"} 1
app_docker_container_up{names="airflow_airflow-webserver_1"} 1
app_docker_container_up{names="airflow_airflow-worker_1"} 1
app_docker_container_up{names="airflow_flower_1"} 1
app_docker_container_up{names="airflow_postgres_1"} 1
app_docker_container_up{names="airflow_redis_1"} 1
app_docker_container_up{names="dc-exporter"} 1
# HELP app_docker_container_up_time docker container uptime
# TYPE app_docker_container_up_time counter
app_docker_container_up_time{names="airflow_airflow-scheduler_1"} 0
app_docker_container_up_time{names="airflow_airflow-triggerer_1"} 0
app_docker_container_up_time{names="airflow_airflow-webserver_1"} 0
app_docker_container_up_time{names="airflow_airflow-worker_1"} 0
app_docker_container_up_time{names="airflow_flower_1"} 0
app_docker_container_up_time{names="airflow_postgres_1"} 0
app_docker_container_up_time{names="airflow_redis_1"} 0
app_docker_container_up_time{names="dc-exporter"} 0
# HELP docker_container_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which docker_container_exporter was built.
# TYPE docker_container_exporter_build_info gauge
docker_container_exporter_build_info{branch="",goversion="go1.16.10",revision="",version=""} 1
```