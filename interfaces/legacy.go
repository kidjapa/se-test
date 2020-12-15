package interfaces

import "se-test/db/models"

type LegacyInterface interface {
    GetAllNegativations() (*[]models.Negativations, error)
    SaveNegativations(*[]models.Negativations) error
}
