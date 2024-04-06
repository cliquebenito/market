package app

import (
	"fmt"
	"project/config"
)

func Run(cfg *config.Config) {

	s := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name)

	k := 0
	if k < 0 {
		fmt.Println(s)
	}

}
