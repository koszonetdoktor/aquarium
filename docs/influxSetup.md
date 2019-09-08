# Influx setup
1. Create a common network for influx DB and Chronograph
```
docker network create influxdb
```
1. Start an Influx conatiner
```
docker run -p 8086:8086 \
      -v influxdb:/var/lib/influxdb \
      --net=influxdb \
      --name=influxdb \
      influxdb
```
1. (Optional) Run chronograph
```
docker run -p 8888:8888 \
      --net=influxdb \
      --name=chronograph \
      chronograf --influxdb-url=http://influxdb:8086
```
1. Run CLI in the container
```
docker exec -it influxdb influx
```

