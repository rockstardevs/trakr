git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.1",
)
load("@io_bazel_rules_go//go:def.bzl",
     "go_repositories",
     "new_go_repository")

go_repositories()

new_go_repository(
    name = "com_github_golang_glog",
    commit = "23def4e6c14b4da8ac2ed8007337bc5eb5007998",
    importpath = "github.com/golang/glog",
)

new_go_repository(
    name = "com_github_gorilla_mux",
    tag = "v1.3.0",
    importpath = "github.com/gorilla/mux",
)