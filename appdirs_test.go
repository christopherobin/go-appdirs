package appdirs_test

import (
	"fmt"
	"github.com/christopherobin/go-appdirs"
)

func ExampleUserDataDir() {
	app := appdirs.AppConf{
		Name: "hello-app",
	}
	fmt.Println(app.UserDataDir())
}

func ExampleSiteDataDir() {
	app := appdirs.AppConf{
		Name: "hello-app",
	}
	fmt.Println(app.SiteDataDir())
}

func ExampleUserConfigDir() {
	app := appdirs.AppConf{
		Name: "hello-app",
	}
	fmt.Println(app.UserConfigDir())
}

func ExampleSiteConfigDir() {
	app := appdirs.AppConf{
		Name: "hello-app",
	}
	fmt.Println(app.SiteConfigDir())
}

func ExampleUserCacheDir() {
	app := appdirs.AppConf{
		Name: "hello-app",
	}
	fmt.Println(app.UserCacheDir())
}

func ExampleUserLogDir() {
	app := appdirs.AppConf{
		Name: "hello-app",
	}
	fmt.Println(app.UserLogDir())
}
