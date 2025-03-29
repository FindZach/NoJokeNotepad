package main

import (
	"context"
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

// App struct
type App struct {
	ctx         context.Context
	currentFile string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	println("App started successfully")
}

// NewFile clears the editor
func (a *App) NewFile() string {
	println("NewFile called")
	a.currentFile = ""
	return ""
}

// OpenFile opens a file and returns its contents
func (a *App) OpenFile() (string, error) {
	println("OpenFile called")
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open File",
		Filters: []runtime.FileFilter{
			{DisplayName: "Text Files (*.txt)", Pattern: "*.txt"},
		},
	})
	if err != nil || filePath == "" {
		return "", err
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	a.currentFile = filePath
	return string(data), nil
}

// SaveFile saves the given content to the current file (if set) or prompts for a new file
func (a *App) SaveFile(content string) (string, error) {
	println("SaveFile called with content:", content)
	if a.currentFile != "" {
		err := os.WriteFile(a.currentFile, []byte(content), 0644)
		if err != nil {
			return "", err
		}
		return a.currentFile, nil
	}
	return a.SaveFileAs(content)
}

// SaveFileAs saves the given content to a new file
func (a *App) SaveFileAs(content string) (string, error) {
	println("SaveFileAs called with content:", content)
	filePath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save File As",
		Filters: []runtime.FileFilter{
			{DisplayName: "Text Files (*.txt)", Pattern: "*.txt"},
		},
	})
	if err != nil || filePath == "" {
		return "", err
	}
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return "", err
	}
	a.currentFile = filePath
	return filePath, nil
}

func main() {
	// Create application with options
	app := NewApp()

	// Create Wails app
	err := wails.Run(&options.App{
		Width:            500,
		Height:           400,
		Title:            "NoJoke Notepad",
		DisableResize:    false,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Frameless:        true,
		Windows: &windows.Options{
			Theme:               windows.SystemDefault,
			DisableWindowIcon:   true,
			WindowIsTranslucent: false,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
