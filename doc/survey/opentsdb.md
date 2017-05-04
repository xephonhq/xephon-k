# OpenTSDB

- GitHub: {{ url }}

## Protocol

- Thrift

It uses a [shim](https://github.com/OpenTSDB/asynccassandra) to connect Cassandra using Thrift protocol in a HBase way.
Detail can be found [here](http://opentsdb.net/docs/build/html/user_guide/backends/cassandra.html)

## Schema

### Metrics Schema

### Meta Schema

### Extra Schema

## Query

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