package google

import (
	"github.com/infracost/infracost/api/resources/google"
	"github.com/infracost/infracost/api/schema"
)

func getMonitoringItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_monitoring_metric_descriptor",
		RFunc: NewMonitoringMetricDescriptor,
	}
}

func NewMonitoringMetricDescriptor(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.MonitoringMetricDescriptor{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
