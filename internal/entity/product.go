package entity

//хранение данных каталога продуктов (на этом этапе можно сделать хранение
//в памяти сервиса или попытаться храниться в БД)

//возможность добавления новых продуктов
//получение информации о продукте
//разделение продуктов по категориям
//добавление, удаление категорий
//возможность получить все продукты из выбранной категории

type Category struct {
	CategoryName string
	Products     []Product
}

type Product struct {
	ID     int
	Name   string
	Weight int
}
type User struct {
	ID       int
	Name     string
	Password string
}
