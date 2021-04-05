package modeltests

import (
	"log"
	"strconv"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kartuludus/wallaster-be/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllCustomers(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedCustomers()
	if err != nil {
		log.Fatal(err)
	}

	Customers, err := CustomerInstance.FindAllCustomers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Customers: %v\n", err)
		return
	}
	assert.Equal(t, len(*Customers), 2)
}

func TestSaveCustomer(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}
	newCustomer := models.Customer{
		ID:       1,
		Email:    "test@gmail.com",
		Name: "test",
		Surname: "tset",
	}
	savedCustomer, err := newCustomer.SaveCustomer(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Customers: %v\n", err)
		return
	}
	assert.Equal(t, newCustomer.ID, savedCustomer.ID)
	assert.Equal(t, newCustomer.Email, savedCustomer.Email)
	assert.Equal(t, newCustomer.Name, savedCustomer.Name)
}

func TestGetCustomerByID(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}

	Customer, err := seedOneCustomer()
	if err != nil {
		log.Fatalf("cannot seed Customers table: %v", err)
	}
	foundCustomer, err := CustomerInstance.FindCustomerByID(server.DB, strconv.Itoa(Customer.ID))
	if err != nil {
		t.Errorf("this is the error getting one Customer: %v\n", err)
		return
	}
	assert.Equal(t, foundCustomer.ID, Customer.ID)
	assert.Equal(t, foundCustomer.Email, Customer.Email)
	assert.Equal(t, foundCustomer.Name, Customer.Name)
}

func TestUpdateACustomer(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}

	Customer, err := seedOneCustomer()
	if err != nil {
		log.Fatalf("Cannot seed Customer: %v\n", err)
	}

	Customer.Name = "modiUpdate"
	Customer.Email= "modiupdate@example.com"

	updatedCustomer, err := Customer.UpdateACustomer(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the Customer: %v\n", err)
		return
	}
	assert.Equal(t, updatedCustomer.ID, Customer.ID)
	assert.Equal(t, updatedCustomer.Email, Customer.Email)
	assert.Equal(t, updatedCustomer.Name, Customer.Name)
}

func TestDeleteACustomer(t *testing.T) {

	err := refreshCustomerTable()
	if err != nil {
		log.Fatal(err)
	}

	Customer, err := seedOneCustomer()

	if err != nil {
		log.Fatalf("Cannot seed Customer: %v\n", err)
	}

	isDeleted, err := CustomerInstance.DeleteACustomer(server.DB, strconv.Itoa(Customer.ID))
	if err != nil {
		t.Errorf("this is the error updating the Customer: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}