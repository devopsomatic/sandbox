/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"github.com/docker/machine/drivers/hyperv"
	"github.com/docker/machine/libmachine/drivers"
	cfg "github.com/stackfoundation/sandbox/core/pkg/minikube/config"
	"github.com/stackfoundation/sandbox/core/pkg/minikube/constants"
)

func createHypervHost(config MachineConfig) drivers.Driver {
	d := hyperv.NewDriver(cfg.GetMachineName(), constants.GetMinipath())
	d.Boot2DockerURL = config.Downloader.GetISOFileURI(config.MinikubeISO)
	d.VSwitch = config.HypervVirtualSwitch
	d.MemSize = config.Memory
	d.CPU = config.CPUs
	d.DiskSize = int(config.DiskSize)
	d.SSHUser = "docker"
	return d
}
