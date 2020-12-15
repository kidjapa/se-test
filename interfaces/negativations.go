package interfaces

import "se-test/db/models"

type NegativationsInterface interface {
    GetById(int64) (*models.Negativations, error)
    GetByCompanyDocument(string) (*[]models.Negativations, error)
    GetByCustomerDocument(string) (*[]models.Negativations, error)
}