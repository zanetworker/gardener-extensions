load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["utils.go"],
    importpath = "github.com/gardener/gardener-extensions/pkg/controller",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/controller/error:go_default_library",
        "//vendor/github.com/gardener/gardener/pkg/apis/extensions/v1alpha1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/client:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/reconcile:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/runtime/signals:go_default_library",
    ],
)
