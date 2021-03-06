package kube

import (
	"context"
	"sync"

	"github.com/stackfoundation/sandbox/core/pkg/workflows/properties"
	workflowsv1 "github.com/stackfoundation/sandbox/core/pkg/workflows/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/pkg/api/v1"
)

const serviceNameKey = "svc"

// PodListener Listener which listens for pod events
type PodListener interface {
	Container(containerID string)
	Ready()
	Done(failed bool, message string)
}

// PodCreationSpec Specification for creating a pod
type PodCreationSpec struct {
	Cleanup          *sync.WaitGroup
	Command          []string
	Context          context.Context
	Environment      *properties.Properties
	Health           *workflowsv1.HealthCheck
	Image            string
	Name             string
	LogPrefix        string
	Ports            []workflowsv1.Port
	Readiness        *workflowsv1.HealthCheck
	Listener         PodListener
	VariableReceiver func(string, string)
	Volumes          []workflowsv1.Volume
	WorkflowReceiver func(string)
}

type podContext struct {
	creationSpec  *PodCreationSpec
	podsClient    corev1.PodInterface
	pod           *v1.Pod
	services      []*v1.Service
	serviceClient corev1.ServiceInterface
	podClosed     chan bool
}
