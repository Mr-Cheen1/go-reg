package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Mr-Cheen1/go-reg/models"
	"github.com/Mr-Cheen1/go-reg/storage"
)

type MainWindow struct {
	window        fyne.Window
	products      *models.Products
	searchResults *models.Products
	storage       *storage.ExcelStorage
	productList   *widget.List
	searchEntry   *widget.Entry
}

func NewMainWindow(window fyne.Window, storage *storage.ExcelStorage) *MainWindow {
	return &MainWindow{
		window:  window,
		storage: storage,
	}
}

func (w *MainWindow) Initialize() error {
	var err error
	products, err := w.storage.Load()
	if err != nil {
		log.Println("Ошибка загрузки данных:", err)
		return err
	}

	w.products = &products
	searchResults := make(models.Products, len(products))
	copy(searchResults, products)
	w.searchResults = &searchResults

	w.createUI()
	return nil
}

func (w *MainWindow) createUI() {
	w.searchEntry = widget.NewEntry()
	w.searchEntry.SetPlaceHolder("Поиск...")

	w.productList = widget.NewList(
		func() int {
			return len(*w.searchResults)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("template"),
				widget.NewButton("Редактировать", nil),
				widget.NewButton("Удалить", nil),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			container := o.(*fyne.Container)
			label := container.Objects[0].(*widget.Label)
			editBtn := container.Objects[1].(*widget.Button)
			deleteBtn := container.Objects[2].(*widget.Button)

			product := (*w.searchResults)[i]

			label.SetText(fmt.Sprintf("%s - %.2f часов", product.Name, product.ProcessingTime))

			editBtn.OnTapped = func() {
				ShowEditDialog(w.window, product, w.productList, w.searchEntry, w.products, w.storage)
			}

			deleteBtn.OnTapped = func() {
				w.products.Delete(product.ID)
				if err := w.storage.Save(*w.products); err != nil {
					log.Println("Ошибка сохранения:", err)
					return
				}
				if w.searchEntry != nil {
					w.searchEntry.SetText("")
					if w.searchEntry.OnChanged != nil {
						w.searchEntry.OnChanged("")
					}
				}
				w.updateListState("")
			}
		},
	)

	w.searchEntry.OnChanged = func(query string) {
		w.updateListState(query)
	}

	addButton := widget.NewButton("Добавить запись", func() {
		ShowAddDialog(w.window, w.productList, w.searchEntry, w.products, w.storage)
	})

	content := container.NewBorder(
		container.NewVBox(
			w.searchEntry,
			container.NewHBox(addButton),
		),
		nil, nil, nil,
		w.productList,
	)

	w.window.SetContent(content)
	w.window.Resize(fyne.NewSize(800, 600))
}

func (w *MainWindow) updateListState(query string) {
	if query != "" {
		*w.searchResults = w.products.Search(query)
	} else {
		*w.searchResults = make(models.Products, len(*w.products))
		copy(*w.searchResults, *w.products)
	}

	if w.productList != nil {
		w.productList.Refresh()
	}
}
