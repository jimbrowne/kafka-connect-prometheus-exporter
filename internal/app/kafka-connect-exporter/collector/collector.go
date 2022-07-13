package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type collector struct {
	URI                      string
	up                       prometheus.Gauge
	connectorsCount          prometheus.Gauge
	isConnectorRunning       *prometheus.Desc
	areConnectorTasksRunning *prometheus.Desc
}

type connectors []string

func NewCollector(uri, nameSpace string) Collector {
	log.Infoln("Collecting data from:", uri)

	return &collector{
		URI: uri,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: nameSpace,
			Name:      "up",
			Help:      "was the last scrape of kafka connect successful?",
		}),
		connectorsCount: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: nameSpace,
			Subsystem: "connectors",
			Name:      "count",
			Help:      "number of deployed connectors",
		}),
		isConnectorRunning: prometheus.NewDesc(
			prometheus.BuildFQName(nameSpace, "connector", "state_running"),
			"is the connector running?",
			[]string{"connector", "consumer_group", "state", "worker_id"},
			nil,
		),
		areConnectorTasksRunning: prometheus.NewDesc(
			prometheus.BuildFQName(nameSpace, "connector", "tasks_state"),
			"the state of tasks. 0-failed, 1-running, 2-unassigned, 3-paused",
			[]string{"connector", "consumer_group", "state", "worker_id", "id"},
			nil,
		),
	}
}
