package compute_instance

import (
	"github.com/kaytu-io/kaytu/pkg/plugin/proto/src/golang"
)

type ComputeInstanceItem struct {
	Name                string
	Id                  string
	MachineType         string
	Region              string
	OptimizationLoading bool
	Preferences         []*golang.PreferenceItem
	Skipped             bool
	LazyLoadingEnabled  bool
	SkipReason          string
	// Wastage             kaytu.EC2InstanceWastageResponse
}

func (i ComputeInstanceItem) ToOptimizationItem() *golang.OptimizationItem {
	oi := &golang.OptimizationItem{
		Id:                 i.Id,
		Name:               i.Name,
		ResourceType:       i.MachineType,
		Region:             i.Region,
		Devices:            nil,
		Preferences:        i.Preferences,
		Description:        "description placeholder",
		Loading:            i.OptimizationLoading,
		Skipped:            i.Skipped,
		SkipReason:         i.SkipReason,
		LazyLoadingEnabled: i.LazyLoadingEnabled,
	}

	// if i.Instance.PlatformDetails != nil {
	// 	oi.Platform = *i.Instance.PlatformDetails
	// }
	if oi.Name == "" {
		oi.Name = string(i.Name)
	}

	return oi
}