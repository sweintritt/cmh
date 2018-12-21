cmake helper
=============

The *cmake helper* (cmh) simplifies the building of libraries for cmake projects. To
build and install a library, provided as a cmake project, just run cmh in the
same directory where the `CMakeLists.txt` is stored.

cmh will create a temporary build directory, run `make` and `make install` and will
then remove the build directory.

Per default cmh will install into `~/install`.

# Options

* `-a`, `--args`: Pass additional arguments to the cmake call.
* `-r`, `--release`: Set `CMAKE_BUILD_TYPE` to *Release*. Default is *Debug*.
* `-s`, `--static`: Build static libraries by setting `BUILD_SHARED_LIBS` to *OFF*.
  Default is *ON*.
* `-d`, `--dry`: Just print what would happen, but don't do anything.
* `-p`, `--prefix`: Set `CMAKE_PREFIX_PATH` and `CMAKE_INSTALL_PREFIX` to the given
  path. Default is `~/install`.
* `--no-install`: Build the project, but skip the installation step.
* `-v`, `--verbose`: Show output of the build.

An example using all options:

    $ cmh --static --release --dry --prefix="~/builds" --no-install
    --args="-DBUILD_TESTS=OFF" --verbose

or using the short options

    $ cmh -s -r -d -p="~/builds" --no-install
    -a="-DBUILD_TESTS=OFF" -v
