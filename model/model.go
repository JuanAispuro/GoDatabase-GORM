package model

import (
	"github.com/jinzhu/gorm"
)

// Model product
type Product struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100); not null"`
	Observations *string `gorm:"type:varchar(100)"`
	Price        int     `gorm:"not null"`
	InvoiceItems []InvoiceItem
}

// Modelo de invoiceheader
type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm:"type:varchar(100); not null"`
	InvoiceItems []InvoiceItem
}

// Modelo de invoiceitem
type InvoiceItem struct {
	gorm.Model
	InvoiceheaderID uint
	ProductID       uint
}
