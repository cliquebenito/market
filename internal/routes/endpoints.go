package routes

import (
	"fmt"
	"net/http"
	"project/config"
)

func Connect() string {
	cfg, _ := config.CheckConfig()
	s := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name)
	fmt.Println(s)
	return s

}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ERR")
	//
}
func GetProducts(w http.ResponseWriter, r *http.Request) {
	Connect()
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE PRODJCTS")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE PRODJCTS")
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ADD CATEGORY")
}
func GetCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET CATEGORIES")
}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE CATEGORIES")
}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE CATEGORIES")
}
