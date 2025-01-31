package aws

import (
	"github.com/awslabs/goformation/v4/cloudformation/dynamodb"
	"github.com/infracost/infracost/api/resources/aws"
	"github.com/infracost/infracost/api/schema"
	log "github.com/sirupsen/logrus"
)

func GetDynamoDBTableRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "AWS::DynamoDB::Table",
		Notes: []string{
			"DAX is not yet supported.",
		},
		RFunc: NewDynamoDBTable,
	}
}

func NewDynamoDBTable(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	cfr, ok := d.CFResource.(*dynamodb.Table)
	if !ok {
		log.Warnf("Skipping resource %s as it did not have the expected type (got %T)", d.Address, d.CFResource)
		return nil
	}

	region := "us-east-1" // TODO figure out how to set region
	billingMode := cfr.BillingMode
	var readCapacity int64
	if cfr.ProvisionedThroughput != nil {
		readCapacity = cfr.ProvisionedThroughput.ReadCapacityUnits
	}
	var writeCapacity int64
	if cfr.ProvisionedThroughput != nil {
		writeCapacity = cfr.ProvisionedThroughput.WriteCapacityUnits
	}

	a := &aws.DynamoDBTable{
		Address:        d.Address,
		Region:         region,
		BillingMode:    billingMode,
		WriteCapacity:  &writeCapacity,
		ReadCapacity:   &readCapacity,
		ReplicaRegions: []string{}, // Global Tables are defined using AWS::DynamoDB::GlobalTable
	}
	a.PopulateUsage(u)

	resource := a.BuildResource()
	resource.Tags = mapTags(cfr.Tags)

	return resource
}
