package ui

import (
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/Mr-Cheen1/go-reg/models"
	"github.com/Mr-Cheen1/go-reg/storage"
	"github.com/Mr-Cheen1/go-reg/utils"
)

func ShowEditDialog(parent fyne.Window, product models.Product, list *widget.List, 
	searchEntry *widget.Entry, products *models.Products, storage *storage.ExcelStorage) {
	
	editWindow := fyne.CurrentApp().NewWindow("Редактировать запись")
	editWindow.SetIcon(theme.DocumentIcon())

	nameEntry := widget.NewEntry()
	timeEntry := widget.NewEntry()

	nameEntry.SetText(product.Name)
	timeEntry.SetText(product.TimeCalculation)

	nameEntry.MultiLine = true
	nameEntry.Wrapping = fyne.TextWrapWord

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Наименование", Widget: nameEntry},
			{Text: "Время обработки", Widget: timeEntry},
		},
		OnSubmit: func() {
			name := strings.TrimSpace(nameEntry.Text)
			if name == "" {
				return
			}
			
			timeCalc := timeEntry.Text
			if strings.TrimSpace(timeCalc) == "" {
				timeCalc = "0"
			}
			totalTime := utils.CalculateTime(timeCalc)

			updatedProduct := models.Product{
				ID:              product.ID,
				Name:            name,
				ProcessingTime:  totalTime,
				TimeCalculation: timeCalc,
			}

			products.Update(updatedProduct)
			if err := storage.Save(*products); err != nil {
				log.Println("Ошибка сохранения:", err)
				return
			}

			if searchEntry != nil {
				searchEntry.SetText("")
				if searchEntry.OnChanged != nil {
					searchEntry.OnChanged("")
				}
			}
			if list != nil {
				list.Refresh()
			}
			editWindow.Close()
		},
		OnCancel: func() {
			editWindow.Close()
		},
		SubmitText: "Сохранить",
		CancelText: "Отмена",
	}

	content := container.NewVBox(form)
	editWindow.SetContent(content)
	editWindow.Resize(fyne.NewSize(400, 300))
	editWindow.CenterOnScreen()
	editWindow.Show()
}

func ShowAddDialog(parent fyne.Window, list *widget.List, searchEntry *widget.Entry, 
	products *models.Products, storage *storage.ExcelStorage) {
	
	addWindow := fyne.CurrentApp().NewWindow("Добавить запись")
	addWindow.SetIcon(theme.DocumentIcon())

	nameEntry := widget.NewEntry()
	timeEntry := widget.NewEntry()

	nameEntry.SetPlaceHolder("Введите наименование")
	timeEntry.SetPlaceHolder("Например: 8+2+5")

	nameEntry.MultiLine = true
	nameEntry.Wrapping = fyne.TextWrapWord

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Наименование", Widget: nameEntry},
			{Text: "Время обработки", Widget: timeEntry},
		},
		OnSubmit: func() {
			name := strings.TrimSpace(nameEntry.Text)
			if name == "" {
				return
			}
			
			timeCalc := timeEntry.Text
			if strings.TrimSpace(timeCalc) == "" {
				timeCalc = "0"
			}
			totalTime := utils.CalculateTime(timeCalc)

			product := models.Product{
				ID:              products.GetNextID(),
				Name:            name,
				ProcessingTime:  totalTime,
				TimeCalculation: timeCalc,
			}
			*products = append(*products, product)

			if err := storage.Save(*products); err != nil {
				log.Println("Ошибка сохранения:", err)
				return
			}

			if searchEntry != nil {
				searchEntry.SetText("")
				if searchEntry.OnChanged != nil {
					searchEntry.OnChanged("")
				}
			}
			addWindow.Close()
		},
		OnCancel: func() {
			addWindow.Close()
		},
		SubmitText: "Сохранить",
		CancelText: "Отмена",
	}

	content := container.NewVBox(form)
	addWindow.SetContent(content)
	addWindow.Resize(fyne.NewSize(400, 300))
	addWindow.CenterOnScreen()
	addWindow.Show()
} 