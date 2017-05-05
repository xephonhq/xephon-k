# Druid: Query

- http://druid.io/docs/latest/querying/querying.html

There are three types of queries

- time series
- top N
- group by (slowest)

## TODO

- [ ] is thee aggregation moving aggregation based on granularity

## Time Series

- http://druid.io/docs/latest/querying/timeseriesquery.html

````json
{
  "queryType": "timeseries",
  "dataSource": "sample_datasource",
  "granularity": "day",
  "descending": "true",
  "filter": {
    "type": "and",
    "fields": [
      { "type": "selector", "dimension": "sample_dimension1", "value": "sample_value1" },
      { "type": "or",
        "fields": [
          { "type": "selector", "dimension": "sample_dimension2", "value": "sample_value2" },
          { "type": "selector", "dimension": "sample_dimension3", "value": "sample_value3" }
        ]
      }
    ]
  },
  "aggregations": [
    { "type": "longSum", "name": "sample_name1", "fieldName": "sample_fieldName1" },
    { "type": "doubleSum", "name": "sample_name2", "fieldName": "sample_fieldName2" }
  ],
  "postAggregations": [
    { "type": "arithmetic",
      "name": "sample_divide",
      "fn": "/",
      "fields": [
        { "type": "fieldAccess", "name": "postAgg__sample_name1", "fieldName": "sample_name1" },
        { "type": "fieldAccess", "name": "postAgg__sample_name2", "fieldName": "sample_name2" }
      ]
    }
  ],
  "intervals": [ "2012-01-01T00:00:00.000/2012-01-03T00:00:00.000" ]
}
````

- `intervals` the time ranges to run the query over
  - just like `start_time` and `end time`
- `granularity` the window of downsampling,
  - http://druid.io/docs/latest/querying/granularities.html
  - NOTE: it will generate `0` for all the time that does not have value!!!
- Timeseries queries normally fill empty interior time buckets with zeroe
  - You can disable all zero-filling with the context flag "skipEmptyBuckets"
- `filter` http://druid.io/docs/latest/querying/filters.html
  - **This is just AST**
- `aggreation`
  - `sum`
  - `avg`
  - `first`
  - `last`

## Query Filter  

- http://druid.io/docs/latest/querying/filters.html

This is `where d1 = v1 and (d2 = v2 or d3 = v3)`

````json
"filter": {
  "type": "and",
  "fields": [
    { "type": "selector", "dimension": "sample_dimension1", "value": "sample_value1" },
    { "type": "or",
      "fields": [
        { "type": "selector", "dimension": "sample_dimension2", "value": "sample_value2" },
        { "type": "selector", "dimension": "sample_dimension3", "value": "sample_value3" }
      ]
    }
  ]
}
````

- in `where outlow in ('Good', 'Bad', 'Ugly')`

````json
{
    "type": "in",
    "dimension": "outlaw",
    "values": ["Good", "Bad", "Ugly"]
}
````

- it allows column compare like `where d1 = d2`, I don't think I can support that, and if it ever used in RDBMS
