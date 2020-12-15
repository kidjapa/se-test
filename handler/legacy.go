package handler

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "se-test/utils/types"
)

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
// @Router /legacy/ [GET]
func (h *Handler) getLegacyAllNegativations(c echo.Context) error {
    ln, err := h.LegacyStore.GetAllNegativations()
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, types.DefaultMessageReturn{
            Code:    http.StatusUnprocessableEntity,
            Message: "getLegacyAllNegativations: " + err.Error(),
        })
    }
    return c.JSON(http.StatusOK, ln)
}

// Get a negativation by Customer Document
// @Summary Get a Negativations by the Customer Document
// @Description Get a Negativations by the Customer Document
// @Param company_document path uint true "Customer Document"
// @Tags negativations
// @Security BasicAuth
// @Produce  json
// @Success 201 {object} types.DefaultMessageReturn
// @Failure 422 {object} types.DefaultMessageReturn
// @Failure 404 {object} types.DefaultMessageReturn
// @Failure 500 {object} types.DefaultMessageReturn
// @Router /legacy/start_import [GET]
func (h *Handler) startLegacyImport(c echo.Context) error {
    ln, err := h.LegacyStore.GetAllNegativations()
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, types.DefaultMessageReturn{
            Code:    http.StatusUnprocessableEntity,
            Message: "startLegacyImport: " + err.Error(),
        })
    }

    err = h.LegacyStore.SaveNegativations(ln)
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, types.DefaultMessageReturn{
            Code:    http.StatusUnprocessableEntity,
            Message: "startLegacyImport: " + err.Error(),
        })
    }

    return c.JSON(http.StatusOK, types.DefaultMessageReturn{
        Code:    http.StatusOK,
        Message: "Import performed successfully!",
    })
}