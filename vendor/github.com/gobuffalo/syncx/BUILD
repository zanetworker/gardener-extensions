load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "byte_map.go",
        "int_map.go",
        "interface_map.go",
        "string_map.go",
        "version.go",
    ],
    importmap = "github.com/gardener/gardener-extensions/vendor/github.com/gobuffalo/syncx",
    importpath = "github.com/gobuffalo/syncx",
    visibility = ["//visibility:public"],
)
