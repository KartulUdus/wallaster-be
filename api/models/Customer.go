package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)


type Customer struct {
	ID uint			`gorm:"primary_key;auto_increment" json:"id"`
	Name string			`gorm:"size:100;not null" json:"name"`
	Surname string		`gorm:"size:100;not null" json:"surname"`
	Birthday time.Time	`json:"birthday"`
	Gender bool		`json:"gender"`
	Email string	`gorm:"size:100;not null;unique" json:"email"`
	Address string `gorm:"size:200" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Customer) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Surname = html.EscapeString(strings.TrimSpace(u.Surname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Customer) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Surname == "" {
			return errors.New("Required Surname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "update":
		if u.ID == 0 {
			return errors.New("Required ID")
		}
		return nil

	default:
		return nil
	}
}

func (u *Customer) FindAllCustomers(db *gorm.DB) (*[]Customer, error) {
	var err error
	users := []Customer{}
	err = db.Debug().Model(&Customer{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]Customer{}, err
	}
	return &users, err
}

func (u *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Customer{}, err
	}
	return u, nil
}

func (u *Customer) FindCustomerByID(db *gorm.DB, id uint32) (*Customer, error) {
	var err error
	err = db.Debug().Model(Customer{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &Customer{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Customer{}, errors.New("Customer Not Found")
	}
	return u, err
}

func (u *Customer) UpdateACustomer(db *gorm.DB, id int) (*Customer, error) {

	db = db.Debug().Model(&Customer{}).Where("id = ?", id).Take(&Customer{}).UpdateColumns(
		map[string]interface{}{
			"name":  u.Name,
			"surname":  u.Surname,
			"email":     u.Email,
			"birthday": u.Birthday,
			"gender": u.Gender,
			"address": u.Address,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &Customer{}, db.Error
	}
	err := db.Debug().Model(&Customer{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &Customer{}, err
	}
	return u, nil
}

func (u *Customer) DeleteACustomer(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Customer{}).Where("id = ?", id).Take(&Customer{}).Delete(&Customer{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}