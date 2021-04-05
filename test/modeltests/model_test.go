package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/kartuludus/wallaster-be/api/controllers"
	"github.com/kartuludus/wallaster-be/api/models"
)

var server = controllers.Server{}
var CustomerInstance = models.Customer{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error


	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to postgres database\n")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the postgres database\n")
	}

}

func refreshCustomerTable() error {
	err := server.DB.DropTableIfExists(&models.Customer{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.Customer{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneCustomer() (models.Customer, error) {
	var now = time.Now()
	refreshCustomerTable()

	user := models.Customer{
		Name: "Peter",
		Surname: "Pan",
		Email:    "test@domain.com",
		Birthday: now.AddDate(-30, 11, 26),
		Gender: true,
		Address: "Toomrüütli 1",
	}

	err := server.DB.Model(&models.Customer{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedCustomers() error {
	var now = time.Now()

	users := []models.Customer{
		models.Customer{
			Name: "Steven",
			Surname: "Victor",
			Email:    "steven@domain.com",
			Birthday: now.AddDate(-22, 3, 0),
		},
		models.Customer{
			Name: "Martin",
			Surname: "Luther",
			Email:    "luther@domain.com",
			Birthday: now.AddDate(-59, 2, 0),
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.Customer{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

