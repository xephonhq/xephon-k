# Prometheus: Translating between monitoring languages

Compare Graphite, PromQL, InfluxQL from Prometheus guys https://www.robustperception.io/translating-between-monitoring-languages/,

## Related

- Turning complete https://www.robustperception.io/conways-life-in-prometheus/

## Review

Start and end

- Graphite: from, until
- PromQL: start, end
- InfluxQL: where clause on time

**Graphite have color and stacked**

Syntax

- Graphite: nested prefix function
- PromQL: C expression like
- InfluxQL: SQL-like

Math

Adding 1 to all values across a set of time series

- Graphite: `offset(seriesname.*, 1)`
- InfluxQL: `SELECT 1 + "value" FROM "seriesname"`
  - NOTE: InfluxQL need to specify value because it support multiple fields
- PromQL: `seriesname + 1`

Arithmetic operations

- Graphite: invert, scale, pow
- PromQL: `+, -, *, /, %`
- InfluxQL: `+, -, *, /`

Functions

- Graphite: `logarithm(seriesname.*, 1)`
- InfluxQL:
  - [ ] TODO: the blog say it does not support function
  - But I found https://docs.influxdata.com/influxdb/v1.2/query_language/functions/ , e... ceiling point to an issue
  - And there is a issue https://github.com/influxdata/influxdb/issues/5930, seems most of these functions are not implemented
  - Custom script support calling all go functions https://docs.influxdata.com/kapacitor/v1.2/tick/expr/#math-functions
- PromQL:
  - https://prometheus.io/docs/querying/functions/
  - day_of_month etc.

Select

- Graphite: `seriesname.filter`
- InfluxQL: `SELECT "value" FROM "seriesnamne" WHERE tag="filter"`
- PromQL: `seriesnamne{labelname="filter"}`

Working Across Series

- Graphite: `averageSeries(seriesname.*)`
- InfluxQL: `SELECT MEAN("value") FROM "seriesname"`
- PromQL: `avg(seriesname)`
  - https://prometheus.io/docs/querying/operators/#aggregation-operators
- [ ] TODO: it seems to be the group by, like I have tags like application, region, instances, now I want to group by region, and use average
for all the instances in the same region

Skipped

- [ ] Range, Number of Unique Values, Mode?
- [ ] Histogram https://prometheus.io/docs/practices/histograms/

Working Across Time

- [ ] moving average

Prediction

- Graphite: Holt-Winters
  - http://graphite.readthedocs.io/en/latest/functions.html#graphite.render.functions.holtWintersForecast
- InfluxQL: Holt-Winters
  - https://docs.influxdata.com/influxdb/v1.2/query_language/functions/#holt-winters
- PromQL:
  - https://prometheus.io/docs/querying/functions/#holt_winters
  - https://prometheus.io/docs/querying/functions/#predict_linear

Binary Operators

math between two different series

- Graphite: `reduceSeries(my.*, "diffSeries", 1, "a", "b")`
- InfluxQL fields: `SELECT "a" - "b" FROM "table"`
- InfluxQL measurements: Not supported
  - TICK script does https://docs.influxdata.com/kapacitor/v1.2/nodes/join_node/
- PromQL: `my_a - my_b`
