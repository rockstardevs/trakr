# Build rules for frontend.
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@com_github_gorilla_mux//:go_default_library",
    ],
)

go_binary(
    name = "server",
    library = ":go_default_library",
    linkstamp = "main",
    visibility = ["//visibility:public"],
)
