package main

import (
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// AppConfig is a global variable to hold the Wails app configuration
var AppConfig = &options.App{}

// Configure platform-specific settings
func init() {
	// Use native Windows title bar
	AppConfig.Windows = &windows.Options{
		Theme: windows.SystemDefault,
	}
}
