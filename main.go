package main

import (
	"os"
	"path/filepath"

	"github.com/asdine/storm/v3"
	"github.com/jschweizer78/fusion-ng/pkg/service"
	"github.com/leaanthony/mewn"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails"
)

// example of BIND from golang to JS
func basic() string {
	return "World!"
}

var dataFileName = "fusion.db"

// DB is golbal Database
var DB *storm.DB

func init() {
	// Open and/or create app data folder
	home := userHomeDir()
	dir := filepath.Join(home, appDataRoot, appDataDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {

		rootDir := filepath.Join(home, appDataRoot)
		if _, err := os.Stat(rootDir); os.IsNotExist(err) {

			err := os.Mkdir(rootDir, 0700)
			if err != nil {
				panic("can't create root appdata folder")
			}
			err = os.Mkdir(dir, 0700)
			if err != nil {
				panic("can't create appdata folder")
			}
		} else {
			err := os.Mkdir(dir, 0700)
			if err != nil {
				panic("found root data, failed creating subfolder")
			}
		}
	}

	// Opening and/or create data base file
	dataBaseFullPath := filepath.Join(dir, dataFileName)
	db, err := storm.Open(dataBaseFullPath)
	if err != nil {
		panic("can't open database file")
	}
	DB = db

	// Opening config file
	viper.SetConfigName("config")
	viper.AddConfigPath(dir)
}

func main() {
	// All good people close file locks
	defer DB.Close()
	srv := service.NewSrvAll(DB)

	js := mewn.String("./frontend/dist/my-app/main-es2015.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Fusion Editor NG",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(srv)
	app.Bind(basic)
	app.Run()
}
