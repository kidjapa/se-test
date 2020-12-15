package models

import (
    "time"
)

// Negativations [...]
type Negativations struct {
    ID               int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"id,omitempty"`
    CompanyDocument  string    `gorm:"column:company_document;type:varchar(55);not null" json:"companyDocument"`
    CompanyName      string    `gorm:"column:company_name;type:varchar(255);not null" json:"companyName"`
    CustomerDocument string    `gorm:"column:customer_document;type:varchar(55);not null" json:"customerDocument"`
    Value            float64   `gorm:"column:value;type:decimal(15,3);not null" json:"value"`
    Contract         string    `gorm:"column:contract;type:varchar(55);not null" json:"contract"`
    DebtDate         time.Time `gorm:"column:debtDate;type:datetime" json:"debtDate"`
    InclusionDate    time.Time `gorm:"column:inclusion_date;type:datetime" json:"inclusionDate"`
}

func (Negativations) TableName() string {
    return "negativations"
}
