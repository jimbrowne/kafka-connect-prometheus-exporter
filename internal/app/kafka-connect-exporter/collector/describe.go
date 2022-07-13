package collector

import "github.com/prometheus/client_golang/prometheus"

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	c.up.Describe(ch)
}
