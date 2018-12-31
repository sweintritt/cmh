# Changelog

## [Unreleased]

### Added

* Added command line option to show version.

### Fixed

* Check if the current directory contains a `CMakeLists.txt`.

## [1.0.0] - 2018-12-21

* Nothing changed, added or fixed. I just wanted to call this version 1.0.0.

## [0.2.0] - 2018-07-22

### Added

* cmh now automatically uses the available number of cores when calling make
* The -a,--args option allows to pass additional arguments to the cmake call
* cmh now only shows the build output if -a,--verbose is set

### Fixed

* Error when calling cmh with dry-run. The build directory is not created and the
  application cannot change to that directory. On dry-run the change is now
  skipped.
* Exchanged err.Error with err.Error() in error messages.

## [0.1.0] - 2018-05-25

### Added

* Initial version of cmake helper

