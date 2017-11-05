load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "bazelsauce",
    srcs = [
        "main.go",
    ],
    data = glob([
        "data/**",
        "client/tester-match/dist/**"
    ]),
    deps = [
        "//context",
        "//handlers"
    ],
    importpath = "github.com/patternMiner/bazelsauce",
    visibility = ["//visibility:public"],
)
