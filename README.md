cmake helper
=============

The *cmake helper* (cmh) simplifies the building of libraries for cmake projects. To
build and install a library, provided as a cmake project, just run cmh in the
same directory where the `CMakeLists.txt` is stored.

cmh will create a temporary build directory, run `make` and `make install` and will
then remove the build directory.

Per default cmh will install into `~/install`.

# Options

* `-r`, `--release`: Set `CMAKE_BUILD_TYPE` to *Release*. Default is *Debug*.
* `-s`, `--static`: Set `BUILD_SHARED_LIBS` to *ON*. Default is *OFF*.
* `-d`, `--dry`: Just print what would happen, but don't do anything.
* `-p`, `--prefix`: Set `CMAKE_PREFIX_PATH` and `CMAKE_INSTALL_PREFIX` to the given
  path. Default is `~/install`.
* `--no-install`: Build the project, but skip the installation step.

An example using all options:

    $ cmh --static --release --dry --prefix="~/builds" --no-install


