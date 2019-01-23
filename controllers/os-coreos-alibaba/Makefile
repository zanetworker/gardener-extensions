# Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

WORKDIR          := $(shell pwd)
REPO_ROOT        := $(WORKDIR)/../..
VERSION          := $(shell cat $(REPO_ROOT)/VERSION)
LD_FLAGS         := "-w -X github.com/gardener/gardener-extensions/controllers/os-coreos-alibaba/pkg/version.Version=$(VERSION)"

.PHONY: start
start:
	@LEADER_ELECTION_NAMESPACE=garden go run \
	-ldflags $(LD_FLAGS) \
		cmd/gardener-extension-os-coreos-alibaba/*.go \
		--kubeconfig=dev/kubeconfig

.PHONY: install
install:
	@GOBIN=${WORKDIR}/bin go install \
		-ldflags $(LD_FLAGS) \
		./cmd/...

.PHONY: examples
examples:
	@${REPO_ROOT}/hack/generate-controllerregistration os-coreos-alibaba OperatingSystemConfig coreos ${WORKDIR}/charts/os-coreos-alibaba ${WORKDIR}/example/controllerregistration.yaml
