"""download mosek"""

def _mosek_repo_impl(repository_ctx):
    url = ""
    sha256 = ""
    platform_str = ""
    dylib_extension = ""
    is_macos = False
    if "amd64" in repository_ctx.os.arch and "linux" in repository_ctx.os.name:
        url = "https://download.mosek.com/stable/10.1.21/mosektoolslinux64x86.tar.bz2"
        sha256 = "f37b7b3806e467c64a02e95b2ab009f6fe8430f25ffc72ed56885f7684dec486"
        platform_str = "linux64x86"
        dylib_extension = "so"
    elif "aarch64" in repository_ctx.os.arch and "mac" in repository_ctx.os.name:
        url = "https://download.mosek.com/stable/10.1.21/mosektoolsosxaarch64.tar.bz2"
        sha256 = "f6e862cab171b7897a6f1ad21c3c0fbdf33dc1310f50c792295ab008321950c7"
        platform_str = "osxaarch64"
        dylib_extension = "dylib"
        is_macos = True
    else:
        fail("doesn't support {} {}", repository_ctx.os.arch, repository_ctx.os.name)

    repository_ctx.download_and_extract(
        stripPrefix = "mosek/10.1",
        url = url,
        sha256 = sha256,
    )
    repository_ctx.file("WORKSPACE.bazel")
    repository_ctx.template(
        "BUILD.bazel",
        repository_ctx.attr._build_tpl,
        substitutions = {
            "%dylib_extension%": dylib_extension,
            "%platform%": platform_str,
        },
    )
    if is_macos:
        install_name_tool = repository_ctx.which("install_name_tool")
        mosek_path = repository_ctx.path("tools/platform/{}/bin/libmosek64.{}".format(platform_str, dylib_extension))
        repository_ctx.execute([install_name_tool, "-id", "@rpath/{}".format(mosek_path.basename), mosek_path])
        tbb_path = repository_ctx.path("tools/platform/{}/bin/libtbb.{}".format(platform_str, dylib_extension))
        repository_ctx.execute([install_name_tool, "-id", "@rpath/{}".format(tbb_path.basename), tbb_path])

mosek_repo = repository_rule(
    doc = "download mosek and set it up",
    implementation = _mosek_repo_impl,
    attrs = {
        "_build_tpl": attr.label(default = ":mosek.BUILD.tpl"),
    },
)
