package google

import (
	"github.com/infracost/infracost/api/resources/google"
	"github.com/infracost/infracost/api/schema"
)

func getComputeMachineImageRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_compute_machine_image",
		RFunc: newComputeMachineImage,
	}
}

func newComputeMachineImage(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := d.Get("region").String()

	r := &google.ComputeMachineImage{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
