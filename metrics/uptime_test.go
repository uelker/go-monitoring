package metrics_test

import (
	"github.com/uelker/go-monitoring/metrics"
	"testing"
	"time"
)

func TestValue(t *testing.T) {
	tenMinutesAgo := time.Now().Add(time.Duration(-10) * time.Minute)
	uptime := metrics.NewUptime(tenMinutesAgo)

	actualUptime, actualUnit := uptime.Value()
	expectedUptime, expectedUnit := time.Since(tenMinutesAgo).Milliseconds(), "milliseconds"

	if actualUptime != expectedUptime || actualUnit != expectedUnit {
		t.Errorf("wrong uptime result, got (%d, %s), want (%d, %s)", actualUptime, actualUnit, expectedUptime, expectedUnit)
	}
}
