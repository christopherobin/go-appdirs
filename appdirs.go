package appdirs

type AppConf struct {
	Name    string
	Author  string
	Version string
}

func HomeDir() (string, error) {
	return homeDir()
}

// Return the full path to the user-specific directory for this application
//
// Typical user data directories are:
//   Mac OS X:               ~/Library/Application Support/<AppName>
//   Unix:                   ~/.local/share/<AppName>    # or in $XDG_DATA_HOME, if defined
//
// Fox Unix, if $XDG_DATA_HOME is defined it will be used for generating this directory
func (conf AppConf) UserDataDir() (string, error) {
	return conf.userDataDir()
}

// Return full path to the user-shared data dir for this application.
//
// Typical user data directories are:
//   Mac OS X:   /Library/Application Support/<AppName>
//   Unix:       /usr/local/share/<AppName> or $XDG_DATA_DIRS[0]/<AppName>
func (conf AppConf) SiteDataDir() (string, error) {
	return conf.siteDataDir()
}

// Return full path to the user-specific config dir for this application.
//
// Typical user data directories are:
//   Mac OS X:               same as user_data_dir
//   Unix:                   ~/.config/<AppName>     # or in $XDG_CONFIG_HOME, if defined
//
// For Unix, we follow the XDG spec and support $XDG_CONFIG_HOME.
// That means, by deafult "~/.config/<AppName>".
func (conf AppConf) UserConfigDir() (string, error) {
	return conf.userConfigDir()
}

// Return full path to the user-shared data dir for this application.
//
// Typical user data directories are:
//   Mac OS X:   same as site_data_dir
//   Unix:       /etc/xdg/<AppName> or $XDG_CONFIG_DIRS[0]/<AppName>
func (conf AppConf) SiteConfigDir() (string, error) {
	return conf.siteConfigDir()
}

// Return full path to the user-specific cache dir for this application.
//
// Typical user data directories are:
//   Mac OS X:               ~/Library/Caches/<AppName>
//   Unix:                   ~/.cache/<AppName> (XDG default)
func (conf AppConf) UserCacheDir() (string, error) {
	return conf.userCacheDir()
}

// Return full path to the user-specific log dir for this application.
//
// Typical user data directories are:
//   Mac OS X:               ~/Library/Logs/<AppName>
//   Unix:                   ~/.cache/<AppName>/logs  # or under $XDG_CACHE_HOME if defined
func (conf AppConf) UserLogDir() (string, error) {
	return conf.userLogDir()
}

// Shortcut for AppConf{name}.UserDataDir
func UserDataDir(name string) (string, error) {
	return (AppConf{Name: name}).UserDataDir()
}

// Shortcut for AppConf{name}.SiteDataDir
func SiteDataDir(name string) (string, error) {
	return (AppConf{Name: name}).SiteDataDir()
}

// Shortcut for AppConf{name}.UserConfigDir
func UserConfigDir(name string) (string, error) {
	return (AppConf{Name: name}).UserConfigDir()
}

// Shortcut for AppConf{name}.SiteConfigDir
func SiteConfigDir(name string) (string, error) {
	return (AppConf{Name: name}).SiteConfigDir()
}

// Shortcut for AppConf{name}.UserCacheDir
func UserCacheDir(name string) (string, error) {
	return (AppConf{Name: name}).UserCacheDir()
}

// Shortcut for AppConf{name}.UserLogDir
func UserLogDir(name string) (string, error) {
	return (AppConf{Name: name}).UserLogDir()
}
