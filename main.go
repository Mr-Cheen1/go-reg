// Package main provides the entry point for the database editor application
package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/Mr-Cheen1/go-reg/storage"
	"github.com/Mr-Cheen1/go-reg/ui"
)

func main() {
	myApp := app.NewWithID("com.mr-cheen1.go-reg")
	myApp.Preferences().SetString("theme", "dark")
	
	window := myApp.NewWindow("Редактор базы данных")

	excelStorage := storage.NewExcelStorage()
	defer excelStorage.Close()

	mainWindow := ui.NewMainWindow(window, excelStorage)
	if err := mainWindow.Initialize(); err != nil {
		log.Printf("Ошибка инициализации: %v\n", err)
		os.Exit(1)
	}

	window.ShowAndRun()
}
