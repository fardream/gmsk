"""module extension"""

load(":mosek_repo.bzl", "mosek_repo")

def _download_mosek_impl(_ctx):
    mosek_repo(name = "mosek")

download_mosek = module_extension(implementation = _download_mosek_impl, os_dependent = True, arch_dependent = True)
