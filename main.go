package main

import (
	"log"

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
		log.Println("Ошибка инициализации:", err)
	}

	window.ShowAndRun()
}
