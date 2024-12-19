package tests

import (
	"os"
	"testing"

	"github.com/Mr-Cheen1/go-reg/models"
	"github.com/Mr-Cheen1/go-reg/storage"
)

func TestExcelStorage(t *testing.T) {
	// Используем временный файл для тестов
	const testFile = "test_database.xlsx"

	// Удаляем тестовый файл после завершения тестов
	defer os.Remove(testFile)

	// Создаем тестовые данные
	testProducts := models.Products{
		{ID: 1, Name: "Тест 1", ProcessingTime: 1.5, TimeCalculation: "1.5"},
		{ID: 2, Name: "Тест 2", ProcessingTime: 3.0, TimeCalculation: "1.5 + 1.5"},
		{ID: 3, Name: "Тест 3", ProcessingTime: 5.0, TimeCalculation: "2 + 3"},
	}

	t.Run("Сохранение и загрузка данных", func(t *testing.T) {
		storage := storage.NewExcelStorage().WithFilename(testFile)
		defer storage.Close()

		// Сохраняем тестовые данные
		err := storage.Save(testProducts)
		if err != nil {
			t.Errorf("Ошибка при сохранении: %v", err)
		}

		// Загружаем данные
		loadedProducts, err := storage.Load()
		if err != nil {
			t.Errorf("Ошибка при загрузке: %v", err)
		}

		// Проверяем количество записей
		if len(loadedProducts) != len(testProducts) {
			t.Errorf("Загружено %d записей, ожидалось %d", len(loadedProducts), len(testProducts))
		}

		// Проверяем каждую запись
		for i, expected := range testProducts {
			loaded := loadedProducts[i]
			if loaded.ID != expected.ID {
				t.Errorf("ID записи %d: получено %d, ожидалось %d", i, loaded.ID, expected.ID)
			}
			if loaded.Name != expected.Name {
				t.Errorf("Name записи %d: получено %s, ожидалось %s", i, loaded.Name, expected.Name)
			}
			if loaded.ProcessingTime != expected.ProcessingTime {
				t.Errorf("ProcessingTime записи %d: получено %f, ожидалось %f",
					i, loaded.ProcessingTime, expected.ProcessingTime)
			}
			if loaded.TimeCalculation != expected.TimeCalculation {
				t.Errorf("TimeCalculation записи %d: получено %s, ожидалось %s",
					i, loaded.TimeCalculation, expected.TimeCalculation)
			}
		}
	})

	t.Run("Создание нового файла", func(t *testing.T) {
		storage := storage.NewExcelStorage().WithFilename(testFile)
		defer storage.Close()

		// Удаляем файл если существует
		os.Remove(testFile)

		// Пробуем загрузить несуществующий файл
		products, err := storage.Load()
		if err != nil {
			t.Errorf("Ошибка при создании нового файла: %v", err)
		}

		// Проверяем что получили пустой список
		if len(products) != 0 {
			t.Errorf("Новый файл должен быть пустым, получено %d записей", len(products))
		}

		// Проверяем что файл создан
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			t.Error("Файл не был создан")
		}
	})

	t.Run("Обновление данных", func(t *testing.T) {
		storage := storage.NewExcelStorage().WithFilename(testFile)
		defer storage.Close()

		// Сначала сохраняем исходные данные
		err := storage.Save(testProducts)
		if err != nil {
			t.Errorf("Ошибка при начальном сохранении: %v", err)
		}

		// Изменяем данные
		updatedProducts := make(models.Products, len(testProducts))
		copy(updatedProducts, testProducts)
		updatedProducts[0].Name = "Обновленный тест 1"
		updatedProducts[1].ProcessingTime = 4.0
		updatedProducts[1].TimeCalculation = "2 + 2"

		// Сохраняем обновленные данные
		err = storage.Save(updatedProducts)
		if err != nil {
			t.Errorf("Ошибка при обновлении: %v", err)
		}

		// Загружаем и проверяем
		loadedProducts, err := storage.Load()
		if err != nil {
			t.Errorf("Ошибка при загрузке обновленных данных: %v", err)
		}

		// Проверяем обновленные данные
		if loadedProducts[0].Name != "Обновленный тест 1" {
			t.Errorf("Имя не обновилось: получено %s, ожидалось 'Обновленный тест 1'",
				loadedProducts[0].Name)
		}
		if loadedProducts[1].ProcessingTime != 4.0 {
			t.Errorf("Время обработки не обновилось: получено %f, ожидалось 4.0",
				loadedProducts[1].ProcessingTime)
		}
		if loadedProducts[1].TimeCalculation != "2 + 2" {
			t.Errorf("Расчет времени не обновился: получено %s, ожидалось '2 + 2'",
				loadedProducts[1].TimeCalculation)
		}
	})
}
