package seed

import (
	"log"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/kartuludus/wallaster-be/api/models"
)

var now = time.Now()

var Customers = []models.Customer{
	models.Customer{
		Name: "Steven",
		Surname: "Victor",
		Email:    "steven@gmail.com",
		Birthday: now.AddDate(-19, 0, 0),
	},
	models.Customer{
		Name: "Martin",
		Surname: "Luther",
		Email:    "luther@gmail.com",
		Birthday: now.AddDate(-59, 0, 0),
	},
	models.Customer{
		Name: "Martin2",
		Surname: "Luther",
		Email:    "luther2@gmail.com",
		Birthday: now.AddDate(-59, 0, 0),
	},
	models.Customer{
		Name: "Martin3",
		Surname: "Luther",
		Email:    "luther3@gmail.com",
		Birthday: now.AddDate(-59, 0, 0),
	},
	models.Customer{
		Name: "Martin4",
		Surname: "Luther",
		Email:    "luther4@gmail.com",
		Birthday: now.AddDate(-59, 0, 0),
	},
	models.Customer{
		Name: "Martin5",
		Surname: "Luther",
		Email:    "luther5@gmail.com",
		Birthday: now.AddDate(-59, 0, 0),
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Customer{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Customer{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}


	for i, _ := range Customers {
		err = db.Debug().Model(&models.Customer{}).Create(&Customers[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Customers table: %v", err)
		}
	}
}