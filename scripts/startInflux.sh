docker run -p 8086:8086 \
      -v influxdb:/var/lib/influxdb \
      --net=influxdb \
      --name=influxdb \
      influxdb
