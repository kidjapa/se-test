package negativations

import (
    "errors"
    "se-test/db"
    "se-test/db/models"
)

var (
    ErrorCompanyDocumentEmpty = errors.New("company document is required")
    ErrorCustomerEmpty = errors.New("customer document is required")
)

type Handler struct {
    DbHandler *db.Handler
}

func NewHandler(conn *db.Handler) *Handler {
    return &Handler{
        DbHandler: conn,
    }
}

func(h *Handler) GetById(id int64) (n *models.Negativations, err error) {
    n = &models.Negativations{}
    dbRes := h.DbHandler.Conn.Model(&models.Negativations{}).
        First(n, id)
    if dbRes.Error != nil {
        err = dbRes.Error
    }
    return
}

func(h *Handler) GetByCompanyDocument(companyDocument string) (ns *[]models.Negativations, err error) {
    if companyDocument == "" {
        err = ErrorCompanyDocumentEmpty
        return
    }
    ns = &[]models.Negativations{}
    dbRes := h.DbHandler.Conn.Model(&models.Negativations{}).
        Where("company_document = ?", companyDocument).
        First(ns)
    if dbRes.Error != nil {
        err = dbRes.Error
    }
    return
}

func(h *Handler) GetByCustomerDocument(customerDocument string) (ns *[]models.Negativations, err error) {
    if customerDocument == "" {
        err = ErrorCustomerEmpty
        return
    }
    ns = &[]models.Negativations{}
    dbRes := h.DbHandler.Conn.Model(&models.Negativations{}).
        Where("customer_document = ?", customerDocument).
        Order("id").
        Find(ns)
    if dbRes.Error != nil {
        err = dbRes.Error
    }
    return
}