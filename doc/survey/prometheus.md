# Prometheus

## In memory representation

`common/model/value.go`

````go
// SamplePair pairs a SampleValue with a Timestamp.
type SamplePair struct {
	Timestamp Time
	Value     SampleValue
}

// Sample is a sample pair associated with a metric.
type Sample struct {
	Metric    Metric      `json:"metric"`
	Value     SampleValue `json:"value"`
	Timestamp Time        `json:"timestamp"`
}

// Samples is a sortable Sample slice. It implements sort.Interface.
type Samples []*Sample
````

`common/model/metric.go`

````go
type Metric LabelSet
````

## Aggregation

`promql/engine.go`

````go
// vector is basically only an alias for model.Samples, but the
// contract is that in a Vector, all Samples have the same timestamp.
type vector []*sample
````

`promql/functions.go`

````go
func aggrOverTime(ev *evaluator, args Expressions, aggrFn func([]model.SamplePair) model.SampleValue) model.Value {
	mat := ev.evalMatrix(args[0])
	resultVector := vector{}

	for _, el := range mat {
		if len(el.Values) == 0 {
			continue
		}

		el.Metric.Del(model.MetricNameLabel)
		resultVector = append(resultVector, &sample{
			Metric:    el.Metric,
			Value:     aggrFn(el.Values),
			Timestamp: ev.Timestamp,
		})
	}
	return resultVector
}

// === avg_over_time(matrix model.ValMatrix) Vector ===
func funcAvgOverTime(ev *evaluator, args Expressions) model.Value {
	return aggrOverTime(ev, args, func(values []model.SamplePair) model.SampleValue {
		var sum model.SampleValue
		for _, v := range values {
			sum += v.Value
		}
		return sum / model.SampleValue(len(values))
	})
}

// === count_over_time(matrix model.ValMatrix) Vector ===
func funcCountOverTime(ev *evaluator, args Expressions) model.Value {
	return aggrOverTime(ev, args, func(values []model.SamplePair) model.SampleValue {
		return model.SampleValue(len(values))
	})
}
````

I don't like how they implement aggregation,
it has too many function calls and can't be optimized much by compiler,
it's kind of the volcano type, but this is really overkill for time series, and using generator would make things much faster.
