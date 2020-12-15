package handler

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "se-test/utils/types"
)

// Get a negativation by id
// @Summary Get a Negativations by the id
// @Description Get a Negativations by the id
// @Param id path uint true "Negativation DB id"
// @Tags negativations
// @Security BasicAuth
// @Produce  json
// @Success 201 {object} models.Negativations
// @Failure 422 {object} types.DefaultMessageReturn
// @Failure 404 {object} types.DefaultMessageReturn
// @Failure 500 {object} types.DefaultMessageReturn
// @Router /negativations/{id} [GET]
func (h *Handler) getNegativationstById(c echo.Context) error {
    req := &negativationsGetByIdRequest{}
    if err := req.bind(c); err != nil {
        return c.JSON(http.StatusUnprocessableEntity, types.DefaultMessageReturn{
            Code:    http.StatusUnprocessableEntity,
            Message: "getNegativationstById: " + err.Error(),
        })
    }else{
        if n, err := h.NegativationsStore.GetById(req.Id); err != nil {
            return c.JSON(http.StatusInternalServerError, types.DefaultMessageReturn{
                Code:    http.StatusInternalServerError,
                Message: "getNegativationstById: " + err.Error(),
            })
        }else{
            return c.JSON(http.StatusOK, n)
        }
    }
}

// Get a Negativation by Company Document
// @Summary Get a Negativations by the Company Document
// @Description Get a Negativations by the Company Document
// @Param company_document path uint true "Company Document"
// @Tags negativations
// @Security BasicAuth
// @Produce  json
// @Success 201 {array} models.Negativations
// @Failure 422 {object} types.DefaultMessageReturn
// @Failure 404 {object} types.DefaultMessageReturn
// @Failure 500 {object} types.DefaultMessageReturn
// @Router /negativations/company_document/{company_document} [GET]
func (h *Handler) getNegativationstByCompanyDocument(c echo.Context) error {
    req := &negativationsGetByCompanyDocumentRequest{}
    if err := req.bind(c); err != nil {
        return c.JSON(http.StatusUnprocessableEntity, types.DefaultMessageReturn{
            Code:    http.StatusUnprocessableEntity,
            Message: "getNegativationstByCompanyDocument: " + err.Error(),
        })
    }else{
        if ns, err := h.NegativationsStore.GetByCompanyDocument(req.CompanyDocument); err != nil {
            return c.JSON(http.StatusInternalServerError, types.DefaultMessageReturn{
                Code:    http.StatusInternalServerError,
                Message: "getNegativationstByCompanyDocument: " + err.Error(),
            })
        }else{
            return c.JSON(http.StatusOK, ns)
        }
    }
}

// Get a negativation by Customer Document
// @Summary Get a Negativations by the Customer Document
// @Description Get a Negativations by the Customer Document
// @Param company_document path uint true "Customer Document"
// @Tags negativations
// @Security BasicAuth
// @Produce  json
// @Success 201 {array} models.Negativations
// @Failure 422 {object} types.DefaultMessageReturn
// @Failure 404 {object} types.DefaultMessageReturn
// @Failure 500 {object} types.DefaultMessageReturn
// @Router /negativations/customer_document/{customer_document} [GET]
func (h *Handler) getNegativationstByCustomerDocument(c echo.Context) error {
    req := &negativationsGetByCustomerDocumentRequest{}
    if err := req.bind(c); err != nil {
        return c.JSON(http.StatusUnprocessableEntity, types.DefaultMessageReturn{
            Code:    http.StatusUnprocessableEntity,
            Message: "getNegativationstByCustomerDocument: " + err.Error(),
        })
    }else{
        if ns, err := h.NegativationsStore.GetByCustomerDocument(req.CustomerDocument); err != nil {
            return c.JSON(http.StatusInternalServerError, types.DefaultMessageReturn{
                Code:    http.StatusInternalServerError,
                Message: "getNegativationstByCustomerDocument: " + err.Error(),
            })
        }else{
            return c.JSON(http.StatusOK, ns)
        }
    }
}

// Get all negativations
// @Summary Get all negativations
// @Description Get all negativations
// @Tags negativations
// @Security BasicAuth
// @Produce  json
// @Success 201 {array} models.Negativations
// @Failure 422 {object} types.DefaultMessageReturn
// @Failure 404 {object} types.DefaultMessageReturn
// @Failure 500 {object} types.DefaultMessageReturn
// @Router /negativations/customer_document/{customer_document} [GET]
func (h *Handler) getAllNegativations(c echo.Context) error {
    if ns, err := h.NegativationsStore.GetAll(); err != nil {
        return c.JSON(http.StatusInternalServerError, types.DefaultMessageReturn{
            Code:    http.StatusInternalServerError,
            Message: "getNegativationstById: " + err.Error(),
        })
    }else{
        return c.JSON(http.StatusOK, ns)
    }
}