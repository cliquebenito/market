package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"os"
	"project/config"
)

func Run(cfg *config.Config) {

	s := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name)

	//	db, err := sql.Open("postgres", s)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	defer db.Close()
	//	// Выполняем запрос к базе данных
	//	rows, err := db.Query("SELECT * FROM client")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer rows.Close()
	//
	//	// Итерируемся по результатам запроса и выводим содержимое
	//	for rows.Next() {
	//		var id int
	//		var name string
	//		err = rows.Scan(&id, &name)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Println(id, name)
	//	}
	//}
	//conn, err := pgx.Connect(context.Background(), s)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close(context.Background())
	//
	//var greeting string
	//err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(greeting)
	dbpool, err := pgxpool.New(context.Background(), s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
