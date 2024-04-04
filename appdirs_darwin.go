package appdirs

import (
	"path/filepath"
)

func userDataDir(name, author, version string, roaming bool) (path string) {
	path = filepath.Join(homeDir(), "Library", "Application Support")

	if name != "" {
		path = filepath.Join(path, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func siteDataDir(name, author, version string) (path string) {
	path = "/Library/Application Support"

	if name != "" {
		path = filepath.Join(path, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func userConfigDir(name, author, version string, roaming bool) (path string) {
	return UserDataDir(name, author, version, roaming)
}

func siteConfigDir(name, author, version string) (path string) {
	return SiteDataDir(name, author, version)
}

func userCacheDir(name, author, version string, opinion bool) (path string) {
	path = filepath.Join(homeDir(), "Library", "Caches")

	if name != "" {
		path = filepath.Join(path, name)
	}

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}

func userLogDir(name, author, version string, opinion bool) (path string) {
	path = filepath.Join(homeDir(), "Library", "Logs", name)

	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}

	return path
}
