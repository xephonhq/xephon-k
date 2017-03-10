# API Read format

## JSON KairosDB & Heoric & OpenTSDB

Some common characteristic of them

- in one query, ther could be multiple series, but there can be only one time range,
  - [ ] I think prometheus and influxdb supports a feature called offset
  - https://prometheus.io/docs/querying/basics/#offset-modifier
  - https://github.com/influxdata/influxdb/issues/1709
  - NOTE: offset is actually quite inefficient, should change start time
- filtering
- down sampling
- [ ] operation across series
  - [ ] how can you substract if there are missmatch, do align, drop these points?

KairosDB

- http://kairosdb.github.io/docs/build/html/restapi/QueryMetrics.html

````json
{
   "start_absolute": 1357023600000,
   "end_relative": {
       "value": "5",
       "unit": "days"
   },
   "time_zone": "Asia/Kabul",
   "metrics": [
       {
           "tags": {
               "host": ["foo", "foo2"],
               "customer": ["bar"]
           },
           "name": "abc.123",
           "limit": 10000,
           "aggregators": [
               {
                   "name": "sum",
                   "sampling": {
                       "value": 10,
                       "unit": "minutes"
                   }
               }
           ]
       },
       {
           "tags": {
               "host": ["foo", "foo2"],
               "customer": ["bar"]
           },
           "name": "xyz.123",
           "aggregators": [
               {
                   "name": "avg",
                   "sampling": {
                       "value": 10,
                       "unit": "minutes"
                   }
               }
           ]
       }
   ]
}
````

Heoric

- https://spotify.github.io/heroic/#!/docs/api/post-query-metrics

````
$key = "hello kitty" and host = foo.example.com
["and", ["$key", "hello kitty"], ["=", "host", "foo.example.com"]]
````

- [ ] TODO: I think following example is wrong, it should be `$key` instead of `key`
  - https://spotify.github.io/heroic/#!/docs/query_language#special-variables

````json
{
  "range": {"type": "relative", "unit": "HOURS", "value": 2},
  "filter": ["and", ["key", "foo"], ["=", "foo", "bar"], ["+", "role"]],
  "groupBy": ["site"]
}
````

````json
{
  "errors": [
    {
      "type": "node",
      "nodeId": "abcd-efgh",
      "nodeUri": "http://example.com",
      "tags": {"site": "lon"},
      "error": "Connection refused",
      "internal": true
    },
    {
      "type": "series",
      "tags": {"site": "lon"},
      "error": "Aggregation too heavy, too many rows from the database would have to be fetched to satisfy the request!",
      "internal": true
    }
  ],
  "result": [
    {
      "hash": "deadbeef",
      "tags": {"foo": "bar"},
      "values": [[1300000000000, 42.0]]
    },
    {
      "hash": "beefdead",
      "tags": {"foo": "baz"},
      "values": [[1300000000000, 42.0]]
    }
  ],
  "range": {
    "end": 1469816790000,
    "start": 1469809590000
  },
  "statistics": {}
}
````

OpenTSDB

- http://opentsdb.net/docs/build/html/api_http/query/index.html

````json
{
    "start": 1356998400,
    "end": 1356998460,
    "queries": [
        {
            "aggregator": "sum",
            "metric": "sys.cpu.0",
            "rate": "true",
            "filters": [
                {
                   "type":"wildcard",
                   "tagk":"host",
                   "filter":"*",
                   "groupBy":true
                },
                {
                   "type":"literal_or",
                   "tagk":"dc",
                   "filter":"lga|lga1|lga2",
                   "groupBy":false
                },
            ]
        },
        {
            "aggregator": "sum",
            "tsuids": [
                "000001000002000042",
                "000001000002000043"
            ]
        }
    ]
}
````
Example Multi-Set Response

````json
[
    {
        "metric": "tsd.hbase.puts",
        "tags": {
            "host": "tsdb-1.mysite.com"
        },
        "aggregatedTags": [],
        "dps": {
            "1365966001": 3758788892,
            "1365966061": 3758804070,
            "1365974281": 3778141673
        }
    },
    {
        "metric": "tsd.hbase.puts",
        "tags": {
            "host": "tsdb-2.mysite.com"
        },
        "aggregatedTags": [],
        "dps": {
            "1365966001": 3902179270,
            "1365966062": 3902197769,
            "1365974281": 3922266478
        }
    }
]
````
