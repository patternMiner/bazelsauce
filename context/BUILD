load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "context",
    srcs = [
        "common.go",
        "context.go",
        "data_fetch_event.go",
        "match_testers.go"
    ],
    deps = ["@patternMiner_async//:async"],
    importpath = "github.com/patternMiner/bazelsauce/context",
    visibility = ["//visibility:public"],
)
