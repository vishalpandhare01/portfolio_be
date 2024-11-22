package migration

import (
	"fmt"
	"log"

	"github.com/vishalpandhare01/portfolio_be/initializer"
	"github.com/vishalpandhare01/portfolio_be/internal/model"
)

func SetUpMigration() {
	err := initializer.DB.AutoMigrate(model.UserModel{}, model.UserProfile{})
	if err != nil {
		log.Fatal("Error in migration", err)
	}
	fmt.Println(("🚀 Migration run successfull !!!"))
}
