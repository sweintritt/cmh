# Changelog

## [Unreleased]

### Added

* cmh now automatically uses the available number of cores when calling make

### Fixed

* Error when calling cmh with dry-run. The build directory is not created and the
  application cannot change to that directory. On dry-run the change is now
  skipped.
* Exchanged err.Error with err.Error() in error messages.

## [0.1.0] - 2017-05-25

### Added

* Initial version of cmake helper

