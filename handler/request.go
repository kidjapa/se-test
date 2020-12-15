package handler

import (
    "github.com/labstack/echo/v4"
)

type negativationsGetByIdRequest struct {
    Id int64 `json:"id" query:"id" form:"id" validate:"required"`
}

func (n *negativationsGetByIdRequest) bind(c echo.Context) error {
    if err := c.Bind(n); err != nil {
        return err
    }
    if err := c.Validate(n); err != nil {
        return err
    }
    return nil
}

type negativationsGetByCompanyDocumentRequest struct {
    CompanyDocument string `json:"company_document" query:"company_document" form:"company_document" validate:"required"`
}

func (n *negativationsGetByCompanyDocumentRequest) bind(c echo.Context) error {
    if err := c.Bind(n); err != nil {
       return err
    }
    // the bind for url parameters is for the version v5
    n.CompanyDocument = c.Param("company_document")
    return nil
}

type negativationsGetByCustomerDocumentRequest struct {
    CustomerDocument string `json:"customer_document" query:"customer_document" form:"customer_document" validate:"required"`
}

func (n *negativationsGetByCustomerDocumentRequest) bind(c echo.Context) error {
    if err := c.Bind(n); err != nil {
       return err
    }
    // the bind for url parameters is for the version v5
    n.CustomerDocument = c.Param("customer_document")
    return nil
}
