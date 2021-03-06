/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GCEProviderConfig struct {
	metav1.TypeMeta `json:",inline"`

	Project     string `json:"project"`
	Zone        string `json:"zone"`
	MachineType string `json:"machineType"`
	Image       string `json:"image"`
}

func (config *GCEProviderConfig) String() string {
	return fmt.Sprintf("Project=%s, Zone=%s, MachineType=%s, Image=%s",
		config.Project, config.Zone, config.MachineType, config.Image)
}
