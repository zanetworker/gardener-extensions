load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/gardener/gardener-extensions/controllers/os-coreos-alibaba/cmd/gardener-extension-os-coreos-alibaba",
    visibility = ["//visibility:private"],
    deps = [
        "//controllers/os-coreos-alibaba/cmd/gardener-extension-os-coreos-alibaba/app:go_default_library",
        "//controllers/os-coreos-alibaba/pkg/version:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/runtime/log:go_default_library",
    ],
)

go_binary(
    name = "gardener-extension-os-coreos-alibaba",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
