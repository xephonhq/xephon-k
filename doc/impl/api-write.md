# API Write format

Following are the format used by popular TSDBs when they do write

- KairosDB use `[1359788400000, 10]` to avoid the overhead of key in data points, don't know if golang can
handle this properly, may need to add custom unmarshal handler
- KairosDB seems to be allowing mixing int and float number, because it store value using blob
  - [ ] we may need to split int and float into different table, I don't know if cassandra allow mixed float and int,
  and there seems to be little reason for changing from int to float for one series from time to time.
- Prometheus has protobuf and text format
  - text https://prometheus.io/docs/instrumenting/exposition_formats/
  - https://github.com/prometheus/prometheus/blob/master/storage/remote/remote.proto
- InfluxDB use line protocol
  - https://docs.influxdata.com/influxdb/v1.2//tools/api/#write
- OpenTSDB use straight forward json, (which is really verbose ....)
  - http://opentsdb.net/docs/build/html/api_http/put.html
- Heroic use JSON

## JSON KairosDB & OpenTSDB

KairosDB

- http://kairosdb.github.io/docs/build/html/restapi/AddDataPoints.html
- serialize <time, value> as array instead of object 

````json
[
  {
      "name": "archive_file_tracked",
      "datapoints": [[1359788400000, 123], [1359788300000, 13.2], [1359788410000, 23.1]],
      "tags": {
          "host": "server1",
          "data_center": "DC1"
      },
      "ttl": 300
  },
  {
      "name": "impedance",
      "type": "complex-number",
      "datapoints": [
          [
              1359788400000,
              {
                  "real": 2.3,
                  "imaginary": 3.4
              }
          ]
      ],
      "tags": {
          "host": "server1",
          "data_center": "DC1"
      }
  },
  {
      "name": "archive_file_search",
      "timestamp": 1359786400000,
      "value": 321,
      "tags": {
          "host": "server2"
      }
  }
]
````

````json
{
  "errors": [
    "Connect to 10.92.4.1:4242 timed out"
  ]
}
````

Heoric 

- https://spotify.github.io/heroic/#!/docs/api/post-write
- its JSON format for <time, value> is same as KairosDB
- it only allows to write one series at a time

````json
{
  "series": {"key": "foo", "tags": {"site": "lon", "host": "www.example.com"}},
  "data": {"type": "points", "data": [[1300000000000, 42.0], [1300001000000, 84.0]]}
}
````

````json
{
  "ok": true
}
````

OpenTSDB

- http://opentsdb.net/docs/build/html/api_http/put.html
- one object per point, several points from same series are not put under one object as an array

````json
[
    {
        "metric": "sys.cpu.nice",
        "timestamp": 1346846400,
        "value": 18,
        "tags": {
           "host": "web01",
           "dc": "lga"
        }
    },
    {
        "metric": "sys.cpu.nice",
        "timestamp": 1346846400,
        "value": 9,
        "tags": {
           "host": "web02",
           "dc": "lga"
        }
    }
]
````

````json
{
    "errors": [
        {
            "datapoint": {
                "metric": "sys.cpu.nice",
                "timestamp": 1365465600,
                "value": "NaN",
                "tags": {
                    "host": "web01"
                }
            },
            "error": "Unable to parse value to a number"
        }
    ],
    "failed": 1,
    "success": 0
}
````