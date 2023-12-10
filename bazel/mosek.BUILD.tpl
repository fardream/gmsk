package(default_visibility = ["//visibility:public"])

_PLATFORM_DIR = "tools/platform/%platform%"

_BIN_DIR = "{}/bin".format(_PLATFORM_DIR)

_HDRS_DIR = "{}/h".format(_PLATFORM_DIR)

_DYLIB_EXTENSION = "%dylib_extension%"

cc_library(
    name = "mosek",
    srcs = glob([
        "{}/libmosek64.{}*".format(_BIN_DIR, _DYLIB_EXTENSION),
        "{}/libtbb.{}*".format(_BIN_DIR, _DYLIB_EXTENSION),
    ]),
    hdrs = ["{}/mosek.h".format(_HDRS_DIR)],
    strip_include_prefix = _HDRS_DIR,
)
