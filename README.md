# GMSK

`gmsk` is an unofficial wrapper for MOSEK, the conic optimizer from [MOSEK ApS](https://www.mosek.com).

[![Go Reference](https://pkg.go.dev/badge/github.com/fardream/gmsk.svg)](https://pkg.go.dev/github.com/fardream/gmsk)

## Setup MOSEK c api for gmsk

GMSK requires the C api for MOSEK, which can be obtained from [MOSEK's website](https://www.mosek.com).
To actually build the package, a recent enough C toolchain, and the environment should be properly set up.

### Set up on Linux

On linux, this can be achieved by setting `CPATH` environment variable and `LIBRARY_PATH`

```shell
# prepend mosek header folder to CPATH, replace {version} and {arch} with the one you have installed
export CPATH=${MSKHOME}/mosek/{version}/tools/platform/{arch}/h${CPATH:+":${CPATH}"}
# prepend mosek lib folder to LIBRARY_PATH
export LIBRARY_PATH=${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin${LIBRARY_PATH:+":${LIBRARY_PATH}"}
export LD_LIBRARY_PATH=${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin${LD_LIBRARY_PATH:+":${LD_LIBRARY_PATH}"}
```

Alternatively, this can set up for cgo specifically

```shell
export CGO_CFLAGS="-I${MSKHOME}/mosek/{version}/tools/platform/{arch}/h"
export CGO_LDFLAGS="-L${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin"
```

**IF** `LD_LIBRARY_PATH` doesn't include MOSEK's binary folder when running the code, the binary can be set up by adding MOSEK's binary folder to `rpath`, for example, the `CGO_LDFLAGS` can be updated to add `rpath`

```shell
export CGO_LDFLAGS="-L${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin -Wl,-rpath=${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin"
```

### Setup on macOS

Follow the installation instructions on MOSEK's website - remember to run `install.py`. Note unless System Integrity Protect (SIP) is turned off on macOS, the environment variable `LD_LIBRARY_PATH` (and macOS specific `DYLD_LIBRARY_PATH`) will **NOT** work.

When building, the setup is similar to linux

```shell
# prepend mosek header folder to CPATH, replace {version} and {arch} with the one you have installed
export CPATH=${MSKHOME}/mosek/{version}/tools/platform/{arch}/h${CPATH:+":${CPATH}"}
# prepend mosek lib folder to LIBRARY_PATH
export LIBRARY_PATH=${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin${LIBRARY_PATH:+":${LIBRARY_PATH}"}
```

or

```shell
export CGO_CFLAGS="-I${MSKHOME}/mosek/{version}/tools/platform/{arch}/h"
export CGO_LDFLAGS="-L${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin"
```

However, the binary may not be able to find `libmosek` at runtime. Besides default search path, macOS will also look for dynlib in the rpaths of the binary - however, it only does so if the load path of the dynlib starts with `@rpath` - which is controlled by the `LC_ID_DYLIB` of the dynlib (in this case, `libmosek64`).

So there are two options

1. Use `rpath`. First add `rpath` to the linker line

   ```shell
   export CGO_LDFLAGS="-L${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin -Wl,-headerpad,128 -Wl,-rpath,${MSKHOME}/mosek/{version}/tools/platform/{arch}/bin" #
   ```

   Then check `LC_ID_DYNLIB` on the `libmosek64.dynlib`, and update it with `install_name_tool` (which is part of Apple's development suite). For example, for 10.0 version of the library

   ```shell
   install_name_tool -id @rpath/libmosek64.10.0.dynlib path/to/libmosek64.10.0.dynlib
   ```

1. Change `LC_ID_DYNLIB` to the full absolute path of `libmosek64`.
   ```shell
   install_name_tool -id /abspath/to/libmosek64.10.0.dynlib path/to/libmosek64.10.0.dynlib
   ```
