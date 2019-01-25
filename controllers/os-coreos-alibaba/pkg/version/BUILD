load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["version.go"],
    importpath = "github.com/gardener/gardener-extensions/controllers/os-coreos-alibaba/pkg/version",
    visibility = ["//visibility:public"],
    x_defs = {"Version": "{STABLE_VERSION}"},  # keep
)
