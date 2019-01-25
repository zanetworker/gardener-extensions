load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "box.go",
        "env.go",
        "file.go",
        "packr.go",
        "version.go",
        "walk.go",
    ],
    importmap = "github.com/gardener/gardener-extensions/vendor/github.com/gobuffalo/packr",
    importpath = "github.com/gobuffalo/packr",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gobuffalo/envy:go_default_library",
        "//vendor/github.com/gobuffalo/packd:go_default_library",
        "//vendor/github.com/markbates/oncer:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)
