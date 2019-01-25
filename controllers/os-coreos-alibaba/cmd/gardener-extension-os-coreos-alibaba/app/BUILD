load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["app.go"],
    importpath = "github.com/gardener/gardener-extensions/controllers/os-coreos-alibaba/cmd/gardener-extension-os-coreos-alibaba/app",
    visibility = ["//visibility:public"],
    deps = [
        "//controllers/os-coreos-alibaba/pkg/coreos-alibaba:go_default_library",
        "//pkg/controller:go_default_library",
        "//pkg/controller/operatingsystemconfig:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
    ],
)
