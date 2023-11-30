package plugins

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	// Name is plugin name
	Name = "memory-limit-scheduler-plugin"
)

var _ framework.FilterPlugin = &MemoryLimitPlugin{}

type MemoryLimitPlugin struct {
	handle framework.Handle
}

func New(_ runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &MemoryLimitPlugin{
		handle: handle,
	}, nil
}

func (s *MemoryLimitPlugin) Name() string {
	return Name
}

func (s *MemoryLimitPlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *framework.NodeInfo) *framework.Status {
	klog.V(2).Infof("filter pod: %v", pod.Name)
	limitMemory := node.Node().Status.Allocatable.Memory().Value()
	// node limitMemory > 50GB, 预留 5GB
	if limitMemory > 50*1024*1024*1024 {
		limitMemory = limitMemory - 5*1024*1024*1024
	} else {
		// node limitMemory < 50GB, 预留 10%
		limitMemory = limitMemory - limitMemory/10
	}
	// check pod limits.memory, it must be less than node allocatable memory
	for _, container := range pod.Spec.Containers {
		if container.Resources.Limits == nil {
			continue
		}
		if container.Resources.Limits.Memory().Value() > limitMemory {
			return framework.NewStatus(
				framework.Unschedulable,
				fmt.Sprintf(
					"pod memory limit %d must be less than node allocatable %d",
					container.Resources.Limits.Memory().Value(),
					limitMemory,
				),
			)
		}
	}
	return framework.NewStatus(framework.Success, "")
}
