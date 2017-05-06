# API Specification Draft

Only support restful API for early version, maybe add a line protocol

## Info

For getting basic info of current tsdb, so the client library can make some choice

- `/info`

````json
{
    "version": "0.1.0",
    "features" : {
        "tag" : false,
        "shard" : false
    },
    "backend": "cassandra"
}
````

- `/info/version`

````json
{
    "version": "0.1.0"
}
````

- `/info/<backend-name>`

````json
{
    "name": "cassandra",
    "version" : "3.0.3",
    "schema" : "bucket"
}
````

## Insert

- `/w` (I know this is not restful, but why bother)
  - `/write`
- [ ] KairosDB use `[1359788400000, 10]` to avoid the overhead of key in data points, don't know if golang can
handle this properly, may need to add custom unmarshal handler
- [ ] KairosDB seems to be allowing mixing int and float number, bec[1359788400000, 10]ause it store value using blob
  - [ ] we may need to split int and float into different table, I don't know if cassandra allow mixed float and int,
  and there seems to be little reason for changing from int to float for one series from time to time.
- Prometheus has protobuf and text format
  - text https://prometheus.io/docs/instrumenting/exposition_formats/
  - https://github.com/prometheus/prometheus/blob/master/storage/remote/remote.proto
- InfluxDB use line protocol
  - https://docs.influxdata.com/influxdb/v1.2//tools/api/#write
- OpenTSDB use straight forward json, (which is really verbose ....)
  - http://opentsdb.net/docs/build/html/api_http/put.html
   
````json
[
    {
        "name": "cpu.load",
        "tags": {
            "region": "us-east",
            "os": "ubuntu"
        },
        "int_points": [
            [1359788400000, 10],
            [1359788300000, 13]
        ]
    },
    {
        "name": "mem.usage",
        "tags": {
            "region": "us-east",
            "os": "ubuntu"
        },
        "float_points": [
            [1359788400000, 10.3],
            [1359788300000, 13.2]
        ]
    }
]
````

## Read

- `/q` (I know this is not restful too)
- [ ] KairosDB support query multiple series at same time, this could be useful, i.e. I want to query mem.total and mem.usage at same time
- [ ] do we need to support limit when start and end time is provided, if so, which direction does the limit apply to
- [ ] how to handle when one metric name actually have multiple series due to tags
  - just return multiple series if they didn't specify group by (aggregation across series)
  - also aggregation across series requires they have matched timestamp, what about holes and offset

This can be treated as 

- [ ] TODO: how to say the exact match in a SQLish way?
- `where __name__ == cpu.usage AND region == en-us AND (connected_to == en-us OR machine_type in (switch, router) )`
- global query criteria that will be applied to all queries that does not specify them
  - start_time
  - end_time
  - aggregator

````json
{
    "use_cache": false,
    "start_time": 1357023600000,
    "end_time": 1357077800000,
    "aggregator": {
        "type": "avg",
        "window": "60s"
    },
    "queries" : [
        {
            "name": "cpu.idle",
            "tags": {"machine":"machine-01","os":"ubuntu"},
            "match_policy": "exact",
            "start_time": 1493363958000,
            "end_time": 1494363958000
        },
        {
            "name": "cpu.usage",
            "match_policy": "filter",
            "filter": {
                "type": "and",
                "l": {
                    "type": "tag_match",
                    "key": "region",
                    "value": "en-us"
                },
                "r": {
                    "type": "or",
                    "l": {
                        "type": "tag_match",
                        "key": "connected_to",
                        "value": "en-us"
                    },
                    "r": {
                        "type": "in",
                        "key": "machine_type",
                        "values": ["switch", "router"]
                    }
                }
            }
        }
    ]
}
````

response

- include a copy of the query
- [ ] same json problem as above

````json
{
    "query" : {
        "use_cache": false,
        "start_time": 1357023600000,
        "end_time": 1357077800000,
        "metrics" : [
            {
                "name": "cpu.idle"
            },
            {
                "name": "cpu.usage"
            }
        ]
    },
    "metrics" : [
        {
            "name": "cpu.idle",
            "values": [
                [1364968800000, 10],
                [1366351200000, 20]
            ]
        }
    ]
}
````
