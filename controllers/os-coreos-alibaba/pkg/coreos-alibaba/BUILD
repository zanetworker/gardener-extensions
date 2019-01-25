load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["actuator.go"],
    importpath = "github.com/gardener/gardener-extensions/controllers/os-coreos-alibaba/pkg/coreos-alibaba",
    visibility = ["//visibility:public"],
    deps = [
        "//controllers/os-coreos-alibaba/pkg/coreos-alibaba/internal:go_default_library",
        "//controllers/os-coreos-alibaba/pkg/coreos-alibaba/internal/cloudinit:go_default_library",
        "//controllers/os-coreos/pkg/coreos:go_default_library",
        "//pkg/controller:go_default_library",
        "//pkg/controller/operatingsystemconfig:go_default_library",
        "//vendor/github.com/gardener/gardener/pkg/apis/extensions/v1alpha1:go_default_library",
        "//vendor/github.com/go-logr/logr:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/client:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/controller/controllerutil:go_default_library",
    ],
)
