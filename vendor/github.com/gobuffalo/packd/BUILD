load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "file.go",
        "file_info.go",
        "interfaces.go",
        "memory_box.go",
        "skip_walker.go",
        "version.go",
    ],
    importmap = "github.com/gardener/gardener-extensions/vendor/github.com/gobuffalo/packd",
    importpath = "github.com/gobuffalo/packd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gobuffalo/syncx:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)
