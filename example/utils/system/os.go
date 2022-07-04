package system

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
)

//IsWindows check if current os is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux check if current os is linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac check if current os is macos
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsEnv gets the value of the environment variable named by the key.
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv sets the value of the environment variable named by the key.
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv remove a single environment variable.
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// CompareOsEnv gets env named by the key and compare it with comparedEnv
func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}

// ExecCommand use shell /bin/bash -c to execute command
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err = cmd.Run()

	if err != nil {
		stderr = string(errOut.Bytes())
	}
	stdout = string(out.Bytes())

	return
}