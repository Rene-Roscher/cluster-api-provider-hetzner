/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta1

import "fmt"

const (
	// ResourceLifecycleOwned is the value we use when tagging resources to indicate
	// that the resource is considered owned and managed by the cluster,
	// and in particular that the lifecycle is tied to the lifecycle of the cluster.
	ResourceLifecycleOwned = ResourceLifecycle("owned")

	// ResourceLifecycleShared is the value we use when tagging resources to indicate
	// that the resource is shared between multiple clusters, and should not be destroyed
	// if the cluster is destroyed.
	ResourceLifecycleShared = ResourceLifecycle("shared")

	// NameKubernetesHetznerCloudProviderPrefix is the tag name used by the cloud provider to logically
	// separate independent cluster resources. We use it to identify which resources we expect
	// to be permissive about state changes.
	// logically independent clusters running in the same AZ.
	// The tag key = NameKubernetesHetznerCloudProviderPrefix + clusterID
	// The tag value is an ownership value.
	NameKubernetesHetznerCloudProviderPrefix = "caph"

	// NameHetznerProviderPrefix is the tag prefix we use to differentiate
	// cluster-api-provider-hetzner owned components from other tooling that
	// uses NameKubernetesClusterPrefix
	// NameHetznerProviderPrefix = "sigs.k8s.io/cluster-api-provider-hetzner/".
	NameHetznerProviderPrefix = "caph-"
	// NameHetznerProviderOwned is the tag name we use to differentiate
	// cluster-api-provider-hetzner owned components from other tooling that
	// uses NameKubernetesClusterPrefix.
	NameHetznerProviderOwned = NameHetznerProviderPrefix + "cluster-"

	// MachineNameTagKey tags related MachineNameTag.
	MachineNameTagKey = "machine." + NameHetznerProviderPrefix + "name"

	// ServerRecycleLabelKey marks an HCloud server as recyclable. When ServerRecycling is enabled for an
	// HCloudMachine, a server with this label (value "true") that matches the HCloudMachine's type and location
	// can be claimed and rebuilt instead of creating a new server. On deletion the server keeps this
	// label and is returned to the recyclable set instead of being deleted.
	ServerRecycleLabelKey = NameHetznerProviderPrefix + "recycle"

	// ServerRecycleAvailableLabelKey marks a recyclable server that no HCloudMachine uses or claims (value
	// "true"). A claim removes it. Returning the server when its HCloudMachine is deleted, releasing a
	// failed claim and reaping an abandoned one set it again.
	ServerRecycleAvailableLabelKey = NameHetznerProviderPrefix + "recycle-available"

	// ServerRecycleClaimantLabelKey holds the name of the HCloudMachine that is currently claiming a
	// recyclable server (value = HCloudMachine name). It is set while a server is being reserved and rebuilt,
	// after the available marker is dropped but before the machine identity is applied. Because this key
	// is not the one the server lookup matches on, a reserved-but-not-yet-rebuilt server is never
	// adopted as a provisioned HCloudMachine, which keeps the claim idempotent if the rebuild fails.
	ServerRecycleClaimantLabelKey = NameHetznerProviderPrefix + "recycle-claimant"

	// ServerRecycleClaimedAtLabelKey holds the Unix timestamp (seconds) at which the claim recorded in
	// ServerRecycleClaimantLabelKey was written. The claimant continues its claim until it completes or
	// fails, also across controller restarts. If the claimant goes away first, the claim would otherwise
	// keep the server out of the pool for good, because a claimed server carries no available marker and
	// therefore matches no lookup. This timestamp lets a recycling HCloudMachine that is looking for a
	// server recognise such an abandoned claim and return the server to the pool. It also tells which
	// rebuild of the server belongs to the current claim.
	ServerRecycleClaimedAtLabelKey = NameHetznerProviderPrefix + "recycle-claimed-at"
)

// ClusterHetznerCloudProviderTagKey generates the key for resources associated a cluster's HCloud cloud provider.
func ClusterHetznerCloudProviderTagKey(name string) string {
	return fmt.Sprintf("%s%s", NameKubernetesHetznerCloudProviderPrefix, name)
}
