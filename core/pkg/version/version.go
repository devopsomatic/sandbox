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

package version

import (
	"strings"

	"github.com/blang/semver"
)

// The current version of the minikube and localkube
// This is a private field and should be set when compiling with --ldflags="-X github.com/stackfoundation/sandbox/core/pkg/version.version=vX.Y.Z"
const VersionPrefix = "v"

var version = "v0.21.0"

var isoVersion = "v0.23.1"

var isoPath = "minikube/iso"

func GetVersion() string {
	return version
}

func GetIsoVersion() string {
	return isoVersion
}

func GetIsoPath() string {
	return isoPath
}

func GetSemverVersion() (semver.Version, error) {
	return semver.Make(strings.TrimPrefix(GetVersion(), VersionPrefix))
}
