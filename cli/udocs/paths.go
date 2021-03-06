package udocs

import (
	"os"
	"path/filepath"
	"runtime"
)

func RootPath() string {
	return udocsRootDir()
}

func ArchivePath() string {
	return filepath.Join(udocsRootDir(), "/var/archive")
}

func BuildPath() string {
	return filepath.Join(udocsRootDir(), "/var/build")
}

func DeployPath() string {
	return filepath.Join(udocsRootDir(), "/var/deploy")
}

func SearchPath() string {
	return filepath.Join(udocsRootDir(), "/var/deploy/search")
}

func ConfPath() string {
	return filepath.Join(udocsRootDir(), "udocs.conf")
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		if home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH"); home == "" {
			return os.Getenv("USERPROFILE")
		}
	}
	return os.Getenv("HOME")
}

func udocsRootDir() string {
	return filepath.Join(userHomeDir(), ".udocs")
}
