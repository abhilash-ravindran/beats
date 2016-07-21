package cpu

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
	//"fmt"
	"github.com/elastic/beats/metricbeat/module/docker/services/stats"
	"github.com/elastic/beats/metricbeat/module/docker/services/config"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	if err := mb.Registry.AddMetricSet("docker", "cpu", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the MetricSet
// As a minimum it must inherit the mb.BaseMetricSet fields, but can be extended with
// additional entries. These variables can be used to persist data or configuration between
// multiple fetch calls.
type MetricSet struct {
	mb.BaseMetricSet
	//ds *statistics.DockerStatistics
}

// New create a new instance of the MetricSet
// Part of new is also setting up the configuration by processing additional
// configuration entries if needed.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	config := config.GetDefaultConf()
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}
	stats.CreateDS(config)
	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
// It returns the event which is then forward to the output. In case of an error, a
// descriptive error must be returned.
func (m *MetricSet) Fetch() ([]common.MapStr, error) {

	events := stats.GetCPUStats()
	return eventsMapping(events), nil
}
