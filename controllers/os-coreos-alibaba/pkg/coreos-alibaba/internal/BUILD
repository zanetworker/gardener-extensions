load("@io_bazel_rules_go//go:def.bzl", "go_embed_data", "go_library", "go_test")

go_embed_data(
    name = "go_default_embed_data",
    src = "templates/cloud-init.sh.template",
    string = True,
    var = "CloudInitTemplate",
)

go_library(
    name = "go_default_library",
    srcs = [
        "generator.go",
        "template_generator.go",
        ":go_default_embed_data",
    ],
    data = glob(["templates/**"]),
    importpath = "github.com/gardener/gardener-extensions/controllers/os-coreos-alibaba/pkg/coreos-alibaba/internal",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gobuffalo/packr:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "internal_suite_test.go",
        "template_generator_test.go",
    ],
    data = glob(["testdata/**"]),
    embed = [":go_default_library"],
    deps = [
        "//vendor/github.com/gobuffalo/packr:go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
    ],
)
