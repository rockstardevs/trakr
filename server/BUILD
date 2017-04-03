# Build rules for frontend.
load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
  name = "frontend",
  srcs = ["main.go"],
  linkstamp = "main",
  deps = [
    "//models:go_default_library",
    "@com_github_golang_glog//:go_default_library",
    "@com_github_gorilla_mux//:go_default_library",
  ]
)