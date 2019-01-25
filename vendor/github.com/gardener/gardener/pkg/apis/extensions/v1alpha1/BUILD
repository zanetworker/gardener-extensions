load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "register.go",
        "types_defaults.go",
        "types_operatingsystemconfig.go",
        "zz_generated.deepcopy.go",
    ],
    importmap = "github.com/gardener/gardener-extensions/vendor/github.com/gardener/gardener/pkg/apis/extensions/v1alpha1",
    importpath = "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gardener/gardener/pkg/apis/extensions:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
    ],
)
