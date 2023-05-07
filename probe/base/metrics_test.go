package base

import (
	"testing"

	"github.com/megaease/easeprobe/global"
	"github.com/megaease/easeprobe/metric"
	"github.com/stretchr/testify/assert"
)

func TestNewMetrics(t *testing.T) {
	name, subsystem, icon := "test", "test-subsystem", "icon"
	expMetrics := &metrics{
		TotalCnt: metric.NewGauge(name, subsystem, name, "total",
			"Total Probed Counts", []string{"name", "status"}),
		TotalTime: metric.NewGauge(name, subsystem, name, "total_time",
			"Total Time(Seconds) of Status", []string{"name", "status"}),
		Duration: metric.NewGauge(name, subsystem, name, "duration",
			"Probe Duration", []string{"name", "status"}),
		Status: metric.NewGauge(name, subsystem, name, "status",
			"Probe Status", []string{"name"}),
		SLA: metric.NewGauge(name, subsystem, name, "sla",
			"Probe SLA", []string{"name"}),
	}
	global.InitEaseProbe(name, icon)
	m := newMetrics(subsystem, name)

	assert.Equal(t, expMetrics, m)
}
