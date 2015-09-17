# AppDirs

This is a somewhat straight port of [`ActiveState/appdirs`](https://github.com/ActiveState/appdirs)
from Python to Go, at least in behavior. The current difference is that no support for Windows is
planned yet.

## Usage

First install the library

```
go get github.com/christopherobin/go-appdirs
```

Then to use it:

```go
package main

import (
	"fmt"
	"github.com/christopherobin/go-appdirs"
)

var appConf := appdirs.AppConf{Name:"my-app"}

func main() {
	// get the user data dir
	fmt.Println(appConf.UserDataDir())
	// get the user config dir
	fmt.Println(appConf.UserConfigDir())
	// get the user cache dir
	fmt.Println(appConf.UserCacheDir())
	// get the user log dir
	fmt.Println(appConf.UserLogDir())
}

```

## API

```go
type AppConf struct {
	Name    string
	Version string
}

// Return the full path to the user-specific directory for this application
//
// Typical user data directories are:
//   Mac OS X:               ~/Library/Application Support/<AppName>
//   Unix:                   ~/.local/share/<AppName>    # or in $XDG_DATA_HOME, if defined
//
// Fox Unix, if $XDG_DATA_HOME is defined it will be used for generating this directory
func (conf AppConf) UserDataDir() (string, error) {}

// Return full path to the user-shared data dir for this application.
//
// Typical user data directories are:
//   Mac OS X:   /Library/Application Support/<AppName>
//   Unix:       /usr/local/share/<AppName> or $XDG_DATA_DIRS[0]/<AppName>
func (conf AppConf) SiteDataDir() (string, error) {}

// Return full path to the user-specific config dir for this application.
//
// Typical user data directories are:
//   Mac OS X:               same as user_data_dir
//   Unix:                   ~/.config/<AppName>     # or in $XDG_CONFIG_HOME, if defined
//
// For Unix, we follow the XDG spec and support $XDG_CONFIG_HOME.
// That means, by deafult "~/.config/<AppName>".
func (conf AppConf) UserConfigDir() (string, error) {}

// Return full path to the user-shared data dir for this application.
//
// Typical user data directories are:
//   Mac OS X:   same as site_data_dir
//   Unix:       /etc/xdg/<AppName> or $XDG_CONFIG_DIRS[0]/<AppName>
func (conf AppConf) SiteConfigDir() (string, error) {}

// Return full path to the user-specific cache dir for this application.
//
// Typical user data directories are:
//   Mac OS X:               ~/Library/Caches/<AppName>
//   Unix:                   ~/.cache/<AppName> (XDG default)
func (conf AppConf) UserCacheDir() (string, error) {}

// Return full path to the user-specific log dir for this application.
//
// Typical user data directories are:
//   Mac OS X:               ~/Library/Logs/<AppName>
//   Unix:                   ~/.cache/<AppName>/logs  # or under $XDG_CACHE_HOME if defined
func (conf AppConf) UserLogDir() (string, error) {}
```

## TODO

* OSX implementation