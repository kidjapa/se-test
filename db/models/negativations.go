package models

import (
    "time"
)

// Negativations [...]
type Negativations struct {
    ID               int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"id"`
    CompanyDocument  string    `gorm:"column:company_document;type:varchar(55);not null" json:"company_document"`
    CompanyName      string    `gorm:"column:company_name;type:varchar(255);not null" json:"company_name"`
    CustomerDocument string    `gorm:"column:customer_document;type:varchar(55);not null" json:"customer_document"`
    Value            float64   `gorm:"column:value;type:decimal(15,3);not null" json:"value"`
    Contract         string    `gorm:"column:contract;type:varchar(55);not null" json:"contract"`
    DebtDate         time.Time `gorm:"column:debtDate;type:datetime" json:"debt_date"`
    InclusionDate    time.Time `gorm:"column:inclusion_date;type:datetime" json:"inclusion_date"`
}

func (Negativations) TableName() string {
    return "negativations"
}
