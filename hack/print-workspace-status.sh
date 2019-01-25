#!/bin/bash

REGISTRY="eu.gcr.io/gardener-project"

echo "STABLE_VERSION $(cat VERSION)"
echo "STABLE_REGISTRY ${REGISTRY}"
echo "STABLE_IMAGE_REPOSITORY ${REGISTRY}/gardener-extension-os-coreos"
