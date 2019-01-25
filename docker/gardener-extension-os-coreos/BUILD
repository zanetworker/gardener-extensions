load("@io_bazel_rules_docker//container:container.bzl", "container_bundle", "container_image")

container_image(
    name = "gardener-extension-os-coreos-internal",
    files = [
        "//controllers/os-coreos/cmd/gardener-extension-os-coreos",
    ],
    stamp = True,
)

container_bundle(
    name = "gardener-extension-os-coreos",
    images = {"{STABLE_IMAGE_REPOSITORY}:{STABLE_VERSION}": "gardener-extension-os-coreos-internal"},
)
