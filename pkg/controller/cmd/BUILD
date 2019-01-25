load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["namedflagset.go"],
    importpath = "github.com/gardener/gardener-extensions/pkg/controller/cmd",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/spf13/pflag:go_default_library"],
)
