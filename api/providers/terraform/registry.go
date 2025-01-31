package terraform

import (
	"strings"
	"sync"

	"github.com/infracost/infracost/api/schema"

	"github.com/infracost/infracost/api/providers/terraform/aws"
	"github.com/infracost/infracost/api/providers/terraform/azure"
	"github.com/infracost/infracost/api/providers/terraform/google"
)

type ResourceRegistryMap map[string]*schema.RegistryItem

var (
	resourceRegistryMap ResourceRegistryMap
	once                sync.Once
)

func GetResourceRegistryMap() *ResourceRegistryMap {
	once.Do(func() {
		resourceRegistryMap = make(ResourceRegistryMap)

		// Merge all resource registries
		for _, registryItem := range aws.ResourceRegistry {
			resourceRegistryMap[registryItem.Name] = registryItem
			resourceRegistryMap[registryItem.Name].DefaultRefIDFunc = aws.GetDefaultRefIDFunc
		}
		for _, registryItem := range createFreeResources(aws.FreeResources, aws.GetDefaultRefIDFunc) {
			resourceRegistryMap[registryItem.Name] = registryItem
		}

		for _, registryItem := range azure.ResourceRegistry {
			resourceRegistryMap[registryItem.Name] = registryItem
			resourceRegistryMap[registryItem.Name].DefaultRefIDFunc = azure.GetDefaultRefIDFunc
		}
		for _, registryItem := range createFreeResources(azure.FreeResources, azure.GetDefaultRefIDFunc) {
			resourceRegistryMap[registryItem.Name] = registryItem
		}

		for _, registryItem := range google.ResourceRegistry {
			resourceRegistryMap[registryItem.Name] = registryItem
			resourceRegistryMap[registryItem.Name].DefaultRefIDFunc = google.GetDefaultRefIDFunc
		}
		for _, registryItem := range createFreeResources(google.FreeResources, google.GetDefaultRefIDFunc) {
			resourceRegistryMap[registryItem.Name] = registryItem
		}
	})

	return &resourceRegistryMap
}

func (r *ResourceRegistryMap) GetReferenceAttributes(resourceDataType string) []string {
	var refAttrs []string
	item, ok := (*r)[resourceDataType]
	if ok {
		refAttrs = item.ReferenceAttributes
	}
	return refAttrs
}

func (r *ResourceRegistryMap) GetCustomRefIDFunc(resourceDataType string) schema.ReferenceIDFunc {
	item, ok := (*r)[resourceDataType]
	if ok {
		return item.CustomRefIDFunc
	}
	return nil
}

func (r *ResourceRegistryMap) GetDefaultRefIDFunc(resourceDataType string) schema.ReferenceIDFunc {
	item, ok := (*r)[resourceDataType]
	if ok {
		return item.DefaultRefIDFunc
	}
	return nil
}

func GetUsageOnlyResources() []string {
	r := []string{}
	r = append(r, aws.UsageOnlyResources...)
	r = append(r, azure.UsageOnlyResources...)
	r = append(r, google.UsageOnlyResources...)
	return r
}

func HasSupportedProvider(rType string) bool {
	return strings.HasPrefix(rType, "aws_") || strings.HasPrefix(rType, "google_") || strings.HasPrefix(rType, "azurerm_")
}

func createFreeResources(l []string, defaultRefsFunc schema.ReferenceIDFunc) []*schema.RegistryItem {
	freeResources := make([]*schema.RegistryItem, 0)
	for _, resourceName := range l {
		freeResources = append(freeResources, &schema.RegistryItem{
			Name:             resourceName,
			NoPrice:          true,
			Notes:            []string{"Free resource."},
			DefaultRefIDFunc: defaultRefsFunc,
		})
	}
	return freeResources
}
