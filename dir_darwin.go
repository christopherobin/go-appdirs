package appdirs

import (
	"os"
	"os/user"
	"path"
	"strings"
)

func homeDir() (string, error) {
	// first check for the $HOME var
	if os.Getenv("HOME") != "" {
		return os.Getenv("HOME"), nil
	}

	// otherwise retrieve from the user
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func expandDir(dir string) (string, error) {
	if !strings.HasPrefix(dir, "~/") {
		return dir, nil
	}

	home, err := homeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, dir[2:]), nil
}

func (conf AppConf) userDataDir() (string, error) {
	home, err := HomeDir()
	if err != nil {
		return "", err
	}

	dataHome = path.Join(home, "Library/Application Support")

	// add app information
	if conf.Name != "" {
		dataHome = path.Join(dataHome, conf.Name)

		if conf.Version != "" {
			dataHome = path.Join(dataHome, conf.Version)
		}
	}

	return dataHome, nil
}

func (conf AppConf) siteDataDir() (string, error) {
	dataDir := "/Library/Application Support"

	if conf.Name != "" {
		dataDir = path.Join(dataDir, conf.Name)

		if conf.Version != "" {
			dataDir = path.Join(dataDir, conf.Version)
		}
	}

	return dataDir, nil
}

func (conf AppConf) userConfigDir() (string, error) {
	return conf.userDataDir()
}

func (conf AppConf) siteConfigDir() (string, error) {
	return conf.siteDataDir()
}

func (conf AppConf) userCacheDir() (string, error) {
	home, err := HomeDir()
	if err != nil {
		return "", err
	}

	cacheDir = path.Join(home, "Library/Caches")

	// add app information
	if conf.Name != "" {
		cacheDir = path.Join(cacheDir, conf.Name)

		if conf.Version != "" {
			cacheDir = path.Join(cacheDir, conf.Version)
		}
	}

	return cacheDir, nil
}

func (conf AppConf) userLogDir() (string, error) {
	home, err := HomeDir()
	if err != nil {
		return "", err
	}

	logDir = path.Join(home, "Library/Logs")

	// add app information
	if conf.Name != "" {
		logDir = path.Join(logDir, conf.Name)

		if conf.Version != "" {
			logDir = path.Join(logDir, conf.Version)
		}
	}

	return logDir, nil
}
