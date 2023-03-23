# Docker-Container-Exporter
Exporter for docker container state.
Prometheus exporter for docker container state, written in Go.

This exporter only export containers state with name


## Install
docker-container-exporter listen on port 9091 by default.


## Docker Container Metrics
* docker_container_created_at{"container_id", "container_name", "image_name"}
* docker_container_is_up{"container_id", "container_name", "image_name"}

```
root@ubuntu20:/home/docker-container-exporter# curl localhost:9091/metrics

# HELP docker_container_created_at docker container created time
# TYPE docker_container_created_at counter
docker_container_created_at{container_id="27f577673e3ea31483f6c11cf47c6ec0daa5626d623d23e181e3d0c6306e0304",container_name="container-php-8.2",image_name="container-php-8.2"} 1.679536525e+09
docker_container_created_at{container_id="2e866e86b152444cd5bd34886c311928474a662fa9c2d2dc9ac6b65510aab560",container_name="container-gcc-4.9",image_name="container-gcc-4.9"} 1.679536525e+09
docker_container_created_at{container_id="459ac3faefb33c16dce5f1884ea4e6f92bf908a1ea938e395472853bc1770f49",container_name="kafka-3",image_name="confluentinc/cp-kafka:7.3.0"} 1.675522467e+09
docker_container_created_at{container_id="4b2abfb5516ddde74e4514f6efa6040d0ec8e7bcf204652e543c8318cb531473",container_name="container-python-2.7",image_name="container-python-2.7"} 1.679536525e+09
docker_container_created_at{container_id="60d6a43258bf069addd8182fb982fbc2338faaf2ac73a5305114549855a0c32f",container_name="kafka-2",image_name="confluentinc/cp-kafka:7.3.0"} 1.675522467e+09
docker_container_created_at{container_id="62dd028133ee7871e01968a38f9cf9c31b77b62146880bef7ce1e9917a6f20cb",container_name="mysql_1",image_name="mysql:8.0.31"} 1.669964384e+09
docker_container_created_at{container_id="64b99a875e07ca66726b91d71d45c475c5c327ae6a5aa4dc96a7863f90c210a6",container_name="zookeeper",image_name="confluentinc/cp-zookeeper:7.3.0"} 1.675522467e+09
docker_container_created_at{container_id="b8bfcb3862b2af1e564024a1cb3e7d8646a4f5b1e353aaebbcc8d0d8a252b9be",container_name="container-python-3.8",image_name="container-python-3.8"} 1.679536525e+09
docker_container_created_at{container_id="e3e14157c4873f15a3922c631d8980f7350ead80a4a6efe340084438971a0fda",container_name="container-php-7.4",image_name="container-php-7.4"} 1.679536525e+09
docker_container_created_at{container_id="f4ce51d9d8ed09534d359e269ccac9ef6233d63892b429587a427ca220609ff3",container_name="kafka-ui",image_name="provectuslabs/kafka-ui"} 1.675522467e+09
docker_container_created_at{container_id="f519990a6cbe49ac61368b29aa189ed3968dcd7f5ee0a09e26068613a77199c4",container_name="kafka-1",image_name="confluentinc/cp-kafka:7.3.0"} 1.675522467e+09
# HELP docker_container_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which docker_container_exporter was built.
# TYPE docker_container_exporter_build_info gauge
docker_container_exporter_build_info{branch="",goversion="go1.19.3",revision="",version=""} 1
# HELP docker_container_is_up docker container is up
# TYPE docker_container_is_up counter
docker_container_is_up{container_id="27f577673e3ea31483f6c11cf47c6ec0daa5626d623d23e181e3d0c6306e0304",container_name="container-php-8.2",image_name="container-php-8.2"} 1
docker_container_is_up{container_id="2e866e86b152444cd5bd34886c311928474a662fa9c2d2dc9ac6b65510aab560",container_name="container-gcc-4.9",image_name="container-gcc-4.9"} 1
docker_container_is_up{container_id="459ac3faefb33c16dce5f1884ea4e6f92bf908a1ea938e395472853bc1770f49",container_name="kafka-3",image_name="confluentinc/cp-kafka:7.3.0"} 0
docker_container_is_up{container_id="4b2abfb5516ddde74e4514f6efa6040d0ec8e7bcf204652e543c8318cb531473",container_name="container-python-2.7",image_name="container-python-2.7"} 1
docker_container_is_up{container_id="60d6a43258bf069addd8182fb982fbc2338faaf2ac73a5305114549855a0c32f",container_name="kafka-2",image_name="confluentinc/cp-kafka:7.3.0"} 0
docker_container_is_up{container_id="62dd028133ee7871e01968a38f9cf9c31b77b62146880bef7ce1e9917a6f20cb",container_name="mysql_1",image_name="mysql:8.0.31"} 1
docker_container_is_up{container_id="64b99a875e07ca66726b91d71d45c475c5c327ae6a5aa4dc96a7863f90c210a6",container_name="zookeeper",image_name="confluentinc/cp-zookeeper:7.3.0"} 0
docker_container_is_up{container_id="b8bfcb3862b2af1e564024a1cb3e7d8646a4f5b1e353aaebbcc8d0d8a252b9be",container_name="container-python-3.8",image_name="container-python-3.8"} 1
docker_container_is_up{container_id="e3e14157c4873f15a3922c631d8980f7350ead80a4a6efe340084438971a0fda",container_name="container-php-7.4",image_name="container-php-7.4"} 1
docker_container_is_up{container_id="f4ce51d9d8ed09534d359e269ccac9ef6233d63892b429587a427ca220609ff3",container_name="kafka-ui",image_name="provectuslabs/kafka-ui"} 1
docker_container_is_up{container_id="f519990a6cbe49ac61368b29aa189ed3968dcd7f5ee0a09e26068613a77199c4",container_name="kafka-1",image_name="confluentinc/cp-kafka:7.3.0"} 0
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 8
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.19.3"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.080488e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 2.080488e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4259
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 570
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 4.233104e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 2.080488e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.58752e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.309568e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 14852
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 4.554752e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 7.897088e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 15422
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 12000
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 63648
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 65280
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 845805
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 491520
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 491520
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.3552656e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 8
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```


## Install
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


