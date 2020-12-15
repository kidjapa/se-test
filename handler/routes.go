package handler

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "se-test/utils"
)

func (h *Handler) Register(v1 *echo.Group) {

    basicAuthMiddleware := middleware.BasicAuth(utils.BasicAuthValidator)

    // Negativations Group router
    gNegativations := v1.Group("/negativations", basicAuthMiddleware)

    gNegativations.GET("/:id", h.getNegativationstById)
    gNegativations.GET("/company_document/:company_document", h.getNegativationstByCompanyDocument)
    gNegativations.GET("/customer_document/:customer_document", h.getNegativationstByCustomerDocument)

}