load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "handlers",
    srcs = [
        "common.go",
        "countries_handler.go",
        "devices_handler.go",
        "info_handler.go",
        "match_handler.go"
    ],
    deps = ["//context"],
    importpath = "github.com/patternMiner/bazelsauce/handlers",
    visibility = ["//visibility:public"],
)
