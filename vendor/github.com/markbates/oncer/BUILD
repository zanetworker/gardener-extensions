load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "deprecate.go",
        "log.go",
        "oncer.go",
    ],
    importmap = "github.com/gardener/gardener-extensions/vendor/github.com/markbates/oncer",
    importpath = "github.com/markbates/oncer",
    visibility = ["//visibility:public"],
)
