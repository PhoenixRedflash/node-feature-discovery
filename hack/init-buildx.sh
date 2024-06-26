#!/usr/bin/env bash
# Copyright 2022 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit -o nounset -o pipefail

export DOCKER_CLI_EXPERIMENTAL=enabled

# We can skip setup if the current builder already has multi-arch
# AND if it isn't the docker driver, which doesn't work
current_builder="$(docker buildx inspect nfd-builder || true)"
# linux/amd64, linux/arm64, linux/riscv64, linux/ppc64le, linux/s390x, linux/386, linux/arm/v7, linux/arm/v6
if ! grep -Eq "^Driver:\s*docker$"  <<<"${current_builder}" && \
     grep -q "linux/amd64" <<<"${current_builder}" && \
     grep -q "linux/arm64" <<<"${current_builder}"; then
  exit 0
fi

# Ensure qemu is in binfmt_misc
# Docker desktop already has these in versions recent enough to have buildx
# We only need to do this setup on linux hosts
if [ "$(uname)" == 'Linux' ]; then
  # NOTE: this is pinned to a digest for a reason!
  docker run --rm --privileged tonistiigi/binfmt:qemu-v6.1.0@sha256:11128304bc582dc7dbaa35947ff3e52e2610d23cecb410ddfa381a6ce74fa763 --install all
fi

# Ensure we use a builder that can leverage it (the default on linux will not)
docker buildx rm nfd-builder || true
docker buildx create --use --name=nfd-builder                 \
  ${http_proxy:+--driver-opt env.http_proxy="$http_proxy"}    \
  ${HTTP_PROXY:+--driver-opt env.HTTP_PROXY="$HTTP_PROXY"}    \
  ${https_proxy:+--driver-opt env.https_proxy="$https_proxy"} \
  ${HTTPS_PROXY:+--driver-opt env.HTTPS_PROXY="$HTTPS_PROXY"} \
  ${no_proxy:+--driver-opt '"env.no_proxy='$no_proxy'"'}      \
  ${NO_PROXY:+--driver-opt '"env.NO_PROXY='$NO_PROXY'"'}
