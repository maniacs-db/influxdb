package influxql_test

import (
	"testing"

	"github.com/influxdata/influxdb/influxql"
)

func TestHoltWinters_AusTourists(t *testing.T) {
	hw := influxql.NewFloatHoltWintersReducer(10, 4, false)
	// Dataset from http://www.inside-r.org/packages/cran/fpp/docs/austourists
	austourists := []influxql.FloatPoint{
		{Time: 1, Value: 30.052513},
		{Time: 2, Value: 19.148496},
		{Time: 3, Value: 25.317692},
		{Time: 4, Value: 27.591437},
		{Time: 5, Value: 32.076456},
		{Time: 6, Value: 23.487961},
		{Time: 7, Value: 28.47594},
		{Time: 8, Value: 35.123753},
		{Time: 9, Value: 36.838485},
		{Time: 10, Value: 25.007017},
		{Time: 11, Value: 30.72223},
		{Time: 12, Value: 28.693759},
		{Time: 13, Value: 36.640986},
		{Time: 14, Value: 23.824609},
		{Time: 15, Value: 29.311683},
		{Time: 16, Value: 31.770309},
		{Time: 17, Value: 35.177877},
		{Time: 18, Value: 19.775244},
		{Time: 19, Value: 29.60175},
		{Time: 20, Value: 34.538842},
		{Time: 21, Value: 41.273599},
		{Time: 22, Value: 26.655862},
		{Time: 23, Value: 28.279859},
		{Time: 24, Value: 35.191153},
		{Time: 25, Value: 41.727458},
		{Time: 26, Value: 24.04185},
		{Time: 27, Value: 32.328103},
		{Time: 28, Value: 37.328708},
		{Time: 29, Value: 46.213153},
		{Time: 30, Value: 29.346326},
		{Time: 31, Value: 36.48291},
		{Time: 32, Value: 42.977719},
		{Time: 33, Value: 48.901525},
		{Time: 34, Value: 31.180221},
		{Time: 35, Value: 37.717881},
		{Time: 36, Value: 40.420211},
		{Time: 37, Value: 51.206863},
		{Time: 38, Value: 31.887228},
		{Time: 39, Value: 40.978263},
		{Time: 40, Value: 43.772491},
		{Time: 41, Value: 55.558567},
		{Time: 42, Value: 33.850915},
		{Time: 43, Value: 42.076383},
		{Time: 44, Value: 45.642292},
		{Time: 45, Value: 59.76678},
		{Time: 46, Value: 35.191877},
		{Time: 47, Value: 44.319737},
		{Time: 48, Value: 47.913736},
	}

	for _, p := range austourists {
		hw.AggregateFloat(&p)
	}

	points := hw.Emit()
	t.Log(points)

}

func TestHoltWinters_USPopulation(t *testing.T) {
	series := []influxql.FloatPoint{
		{Time: 1, Value: 3.93},
		{Time: 2, Value: 5.31},
		{Time: 3, Value: 7.24},
		{Time: 4, Value: 9.64},
		{Time: 5, Value: 12.90},
		{Time: 6, Value: 17.10},
		{Time: 7, Value: 23.20},
		{Time: 8, Value: 31.40},
		{Time: 9, Value: 39.80},
		{Time: 10, Value: 50.20},
		{Time: 11, Value: 62.90},
		{Time: 12, Value: 76.00},
		{Time: 13, Value: 92.00},
		{Time: 14, Value: 105.70},
		{Time: 15, Value: 122.80},
		{Time: 16, Value: 131.70},
		{Time: 17, Value: 151.30},
		{Time: 18, Value: 179.30},
		{Time: 19, Value: 203.20},
	}
	hw := influxql.NewFloatHoltWintersReducer(10, 0, true)
	for _, p := range series {
		hw.AggregateFloat(&p)
	}

	points := hw.Emit()
	t.Log(points)
}
