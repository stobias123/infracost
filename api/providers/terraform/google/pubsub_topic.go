package google

import (
	"github.com/infracost/infracost/api/resources/google"
	"github.com/infracost/infracost/api/schema"
)

func getPubSubTopicRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_pubsub_topic",
		RFunc: NewPubSubTopic,
	}
}

func NewPubSubTopic(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.PubSubTopic{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
