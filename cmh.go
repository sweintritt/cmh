// cmake helper (cmh) helps building and installing libraries
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var version string = "1.1.0"

// settings stores all available options and is passed to any funcion
type settings struct {
	static     bool
	release    bool
	build_dir  string
	prefix_dir string
	source_dir string
	dry_run    bool
	no_install bool
	verbose    bool
	args       string
}

// create a new settings struct with default values
func newSettings() *settings {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("unable to get current directory:", err.Error())
		return nil
	}

	return &settings{
		static:     false,
		release:    false,
		build_dir:  dir + "/" + tmpDir(),
		prefix_dir: "~/install",
		source_dir: dir,
		dry_run:    false,
		no_install: false,
		verbose:    false,
		args:       ""}
}

// return true if the given path exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

// returns the given relative path as absolute path
func absPath(path string) string {
	if strings.Contains(path, "~") {
		usr, _ := user.Current()
		path = strings.Replace(path, "~", usr.HomeDir, 1)
	}

	path, _ = filepath.Abs(path)
	return path
}

// return the cmake build type name
func buildStr(release bool) string {
	if release {
		return "Release"
	} else {
		return "Debug"
	}
}

// converts bool to cmake option value
func optStr(value bool) string {
	if value {
		return "ON"
	} else {
		return "OFF"
	}
}

// creates a new tmp dir name
func tmpDir() string {
	timestamp := time.Now()
	return fmt.Sprintf("tmp-build-%d", timestamp.Unix())
}

// prepares for the build by creating required directories
func prepare(s *settings) bool {
	if s.dry_run {
		fmt.Println("build directory", s.build_dir, "would have been created")
	} else {
		r, _ := exists(s.build_dir)
		if r {
			fmt.Println(s.build_dir, "already exists. This shout not be happening...")
			return false
		} else {
			err := os.MkdirAll(s.build_dir, os.ModePerm)
			if err != nil {
				fmt.Println("unable for create", s.build_dir, ":", err.Error())
				return false
			}
		}
	}

	rc, _ := exists(s.prefix_dir)
	if rc {
		fmt.Println(s.prefix_dir, "does already exist")
	} else {
		if s.dry_run {
			fmt.Println("prefix directory", s.prefix_dir, "would have been created")
		} else {
			err := os.MkdirAll(s.prefix_dir, os.ModePerm)
			if err != nil {
				fmt.Println("unable for create", s.prefix_dir, ":", err.Error())
				return false
			}
		}
	}

	return true
}

// runs the given command with the given arguments
func run(bin string, verbose bool, args ...string) error {
	dir, _ := os.Getwd()
	fmt.Println("calling:", bin, "with '", args, "' in", dir)
	cmd := exec.Command(bin, args...)
	if verbose {
		cmd.Stdout = os.Stdout
	} else {
		cmd.Stdout = nil
	}
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	cmd.Wait()
	return err
}

// runs cmake with the given options
func cmake(s *settings) bool {
	options := make([]string, 6)
	options[0] = fmt.Sprintf("-DCMAKE_BUILD_TYPE=%s", buildStr(s.release))
	options[1] = fmt.Sprintf("-DBUILD_SHARED_LIBS=%s", optStr(!s.static))
	options[2] = fmt.Sprintf("-DCMAKE_INSTALL_PREFIX=%s", s.prefix_dir)
	options[3] = fmt.Sprintf("-DCMAKE_PREFIX_PATH=%s", s.prefix_dir)
	options[4] = s.args
	options[5] = fmt.Sprintf(s.source_dir)

	if s.dry_run {
		fmt.Println("cmake would have been called with this options:", options)
	} else {
		err := run("cmake", s.verbose, options...)
		if err != nil {
			fmt.Println("cmake failed:", err.Error())
			return false
		}
	}

	return true
}

// runs make
func build(s *settings) bool {
	cores := runtime.NumCPU()

	if s.dry_run {
		fmt.Printf("'make -j%d' would have been called in %s\n", cores, s.build_dir)
	} else {
		err := run("make", s.verbose, "-j"+strconv.Itoa(cores))
		if err != nil {
			fmt.Println("build failed:", err.Error())
			return false
		}
	}

	return true
}

// runs make install
func install(s *settings) bool {
	if s.dry_run {
		fmt.Println("'make install' would have been called in", s.build_dir)
	} else {
		err := run("make", s.verbose, "install")
		if err != nil {
			fmt.Println("install failed:", err.Error())
			return false
		}
	}

	return true
}

// removes the build directory
func clean(s *settings) {
	err := os.Chdir(s.source_dir)
	if err != nil {
		fmt.Println(err.Error())
	}

	if s.dry_run {
		fmt.Println(s.build_dir, "would have been removed")
	} else {
		fmt.Println("Removing", s.build_dir)
		os.RemoveAll(s.build_dir)
	}
}

// changes to the given directory, returns true on success
func chdir(dir string) bool {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("unable to change directory:", err.Error())
		return false
	}

	return true
}

// run the complete build process
func work(s *settings) {
    // TODO check for CMakeLists.txt
	result := prepare(s)

	if result && !s.dry_run {
		result = chdir(s.build_dir)
	}

	if result {
		cmake(s)
	}

	if result {
		build(s)
	}

	if result {
		install(s)
	}

	if result && !s.dry_run {
		os.Chdir(s.source_dir)
	}

	exists, _ := exists(s.build_dir)
	if result || exists {
		clean(s)
	}
}

func main() {
	s := newSettings()
	if s == nil {
		return
	}

	flag.StringVar(&s.args, "a", s.args, "Pass additional arguments to the cmake call")
	flag.StringVar(&s.args, "args", s.args, "Pass additional arguments to the cmake call")
	flag.BoolVar(&s.release, "r", s.release, "Set CMAKE_BUILD_TYPE to 'Release'")
	flag.BoolVar(&s.release, "release", s.release, "Set CMAKE_BUILD_TYPE to 'Release'")
	flag.BoolVar(&s.static, "s", s.static, "Set BUILD_SHARED_LIBS to OFF")
	flag.BoolVar(&s.static, "static", s.static, "Set BUILD_SHARED_LIBS to OFF")
	flag.BoolVar(&s.dry_run, "d", s.dry_run, "Dry run. Just show what commands would have been executed.")
	flag.BoolVar(&s.dry_run, "dry", s.dry_run, "Dry run. Just show what commands would have been executed.")
	flag.StringVar(&s.prefix_dir, "p", s.prefix_dir, "Used for CMAKE_PREFIX_PATH and CMAKE_INSTALL_PREFIX. Default is "+s.prefix_dir)
	flag.StringVar(&s.prefix_dir, "prefix", s.prefix_dir, "Used for CMAKE_PREFIX_PATH and CMAKE_INSTALL_PREFIX. Default is "+s.prefix_dir)
	flag.BoolVar(&s.no_install, "no-install", s.no_install, "Do not run 'make install' after build")
	flag.BoolVar(&s.verbose, "v", s.verbose, "Show the output of the build")
	flag.BoolVar(&s.verbose, "verbose", s.verbose, "Show the output of the build")
	flag.Parse()

	s.prefix_dir = absPath(s.prefix_dir)
	work(s)
}
