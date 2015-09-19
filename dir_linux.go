package appdirs

import (
	"os"
	"os/user"
	"path"
	"strings"
)

const (
	PATHSEP = ":"
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
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	dataHome, err := expandDir(os.Getenv("XDG_DATA_HOME"))
	if err != nil {
		return "", err
	}

	if dataHome == "" {
		home, err := HomeDir()
		if err != nil {
			return "", err
		}

		dataHome = path.Join(home, ".local/share")
	}

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
	dataDir := "/usr/local/share"

	if conf.Prefix != "" {
		dataDir = path.Join(conf.Prefix, "share")
	}

	// special case if the prefix is /, use /usr/share
	if conf.Prefix == "/" {
		dataDir = "/usr/share"
	}

	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	if os.Getenv("XDG_DATA_DIRS") != "" {
		var err error
		dataDir, err = expandDir(strings.Split(os.Getenv("XDG_DATA_DIRS"), PATHSEP)[0])
		if err != nil {
			return "", nil
		}
	}

	if conf.Name != "" {
		dataDir = path.Join(dataDir, conf.Name)

		if conf.Version != "" {
			dataDir = path.Join(dataDir, conf.Version)
		}
	}

	return dataDir, nil
}

func (conf AppConf) userConfigDir() (string, error) {
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	configHome, err := expandDir(os.Getenv("XDG_CONFIG_HOME"))
	if err != nil {
		return "", err
	}

	if configHome == "" {
		home, err := HomeDir()
		if err != nil {
			return "", err
		}

		configHome = path.Join(home, ".config")
	}

	// add app information
	if conf.Name != "" {
		configHome = path.Join(configHome, conf.Name)

		if conf.Version != "" {
			configHome = path.Join(configHome, conf.Version)
		}
	}

	return configHome, nil
}

func (conf AppConf) siteConfigDir() (string, error) {
	configDir := "/etc/xdg"

	if conf.Prefix != "" {
		configDir = path.Join(conf.Prefix, "etc")
	}

	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	if os.Getenv("XDG_CONFIG_DIRS") != "" {
		var err error
		configDir, err = expandDir(strings.Split(os.Getenv("XDG_CONFIG_DIRS"), PATHSEP)[0])
		if err != nil {
			return "", nil
		}
	}

	if conf.Name != "" {
		configDir = path.Join(configDir, conf.Name)

		if conf.Version != "" {
			configDir = path.Join(configDir, conf.Version)
		}
	}

	return configDir, nil
}

func (conf AppConf) userCacheDir() (string, error) {
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	cacheDir, err := expandDir(os.Getenv("XDG_CACHE_HOME"))
	if err != nil {
		return "", err
	}

	if cacheDir == "" {
		home, err := HomeDir()
		if err != nil {
			return "", err
		}

		cacheDir = path.Join(home, ".cache")
	}

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
	// first use the XDG_DATA_HOME env variable, otherwise fallback to a safe default
	cacheDir, err := conf.userCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(cacheDir, "logs"), nil
}
