load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "actuator.go",
        "predicate.go",
        "types.go",
    ],
    importpath = "github.com/gardener/gardener-extensions/controllers/os-coreos/pkg/coreos",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/controller:go_default_library",
        "//pkg/controller/operatingsystemconfig:go_default_library",
        "//vendor/github.com/gardener/gardener/pkg/apis/extensions/v1alpha1:go_default_library",
        "//vendor/github.com/go-logr/logr:go_default_library",
        "//vendor/gopkg.in/yaml.v2:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/client:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/controller/controllerutil:go_default_library",
        "//vendor/sigs.k8s.io/controller-runtime/pkg/predicate:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "actuator_test.go",
        "coreos_suite_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
    ],
)
