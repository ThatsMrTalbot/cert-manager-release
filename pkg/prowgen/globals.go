/*
Copyright 2022 The cert-manager Authors.

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

package prowgen

const (
	// CommonTestImage defines the common base image used across many prow jobs
	CommonTestImage = "eu.gcr.io/jetstack-build-infra-images/bazelbuild:20220629-ee75d11-4.2.1"

	// AlertEmailAddress is the address to which testgrid alerts should be sent
	AlertEmailAddress = "cert-manager-dev-alerts@googlegroups.com"
)
