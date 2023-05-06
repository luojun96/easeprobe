package http

import (
	"testing"

	"github.com/megaease/easeprobe/global"
	"github.com/megaease/easeprobe/metric"
	"github.com/stretchr/testify/assert"
)

func TestNewMetrics(t *testing.T) {
	name, subsystem, icon := "test", "test-subsystem", "icon"
	expMetrics := &metrics{
		StatusCode: metric.NewCounter(name, subsystem, name, "status_code",
			"HTTP Status Code", []string{"name", "status"}),
		ContentLen: metric.NewGauge(name, subsystem, name, "content_len",
			"HTTP Content Length", []string{"name", "status"}),
		DNSDuration: metric.NewGauge(name, subsystem, name, "dns_duration",
			"DNS Duration", []string{"name", "status"}),
		ConnectDuration: metric.NewGauge(name, subsystem, name, "connect_duration",
			"TCP Connection Duration", []string{"name", "status"}),
		TLSDuration: metric.NewGauge(name, subsystem, name, "tls_duration",
			"TLS Duration", []string{"name", "status"}),
		SendDuration: metric.NewGauge(name, subsystem, name, "send_duration",
			"Send Duration", []string{"name", "status"}),
		WaitDuration: metric.NewGauge(name, subsystem, name, "wait_duration",
			"Wait Duration", []string{"name", "status"}),
		TransferDuration: metric.NewGauge(name, subsystem, name, "transfer_duration",
			"Transfer Duration", []string{"name", "status"}),
		TotalDuration: metric.NewGauge(name, subsystem, name, "total_duration",
			"Total Duration", []string{"name", "status"}),
	}
	global.InitEaseProbe(name, icon)
	m := newMetrics(subsystem, name)

	assert.Equal(t, expMetrics, m)
}
