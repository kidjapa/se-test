package negativations

import (
    "gopkg.in/go-playground/assert.v1"
    "gorm.io/gorm"
    "se-test/tests"
    "testing"
)

func TestHandler_GetByCompanyDocument(t *testing.T) {

    dbHandler := tests.NewTestDBConnectionsHandlerLocal()

    cdl := map[string]struct {
        CompanyNameExpected      string
        CustomerDocumentExpected string
    }{
        "59291534000167": {CompanyNameExpected: "ABC S.A.", CustomerDocumentExpected: "51537476467"},
        "77723018000146": {CompanyNameExpected: "123 S.A.", CustomerDocumentExpected: "51537476467"},
        "4843574000182":  {CompanyNameExpected: "DBZ S.A.", CustomerDocumentExpected: "26658236674"},
        "23993551000107": {CompanyNameExpected: "XPTO S.A.", CustomerDocumentExpected: "62824334010"},
    }

    // new negativations controller handler
    h := NewHandler(dbHandler)

    for cd, exp := range cdl {
        ns, err := h.GetByCompanyDocument(cd)
        if err != nil {
            t.Error(err)
        } else {
            assert.Equal(t, 1, len(*ns))
            for _, n := range *ns {
                assert.Equal(t, exp.CompanyNameExpected, n.CompanyName)
                assert.Equal(t, exp.CustomerDocumentExpected, n.CustomerDocument)
            }
        }
    }

    cdlError := map[string]struct {
        ErrorExpected error
    }{
        "asdf": {ErrorExpected: gorm.ErrRecordNotFound},
    }

    for cd, exp := range cdlError {
        _, err := h.GetByCompanyDocument(cd)
        assert.IsEqual(exp.ErrorExpected, err)
    }

}

func TestHandler_GetById(t *testing.T) {
    dbHandler := tests.NewTestDBConnectionsHandlerLocal()

    cdl := map[int64]struct {
        CompanyNameExpected      string
        CustomerDocumentExpected string
    }{
        1: {CompanyNameExpected: "ABC S.A.", CustomerDocumentExpected: "51537476467"},
        2: {CompanyNameExpected: "123 S.A.", CustomerDocumentExpected: "51537476467"},
        3: {CompanyNameExpected: "DBZ S.A.", CustomerDocumentExpected: "26658236674"},
        4: {CompanyNameExpected: "XPTO S.A.", CustomerDocumentExpected: "62824334010"},
    }

    // new negativations controller handler
    h := NewHandler(dbHandler)

    for cd, exp := range cdl {
        n, err := h.GetById(cd)
        if err != nil {
            t.Error(err)
        } else {
            assert.Equal(t, exp.CompanyNameExpected, n.CompanyName)
            assert.Equal(t, exp.CustomerDocumentExpected, n.CustomerDocument)
        }
    }

    cdlError := map[string]struct {
        ErrorExpected error
    }{
        "asdf": {ErrorExpected: gorm.ErrInvalidData},
    }

    for cd, exp := range cdlError {
        _, err := h.GetByCompanyDocument(cd)
        assert.IsEqual(exp.ErrorExpected, err)
    }
}

func TestHandler_GetByCustomerDocument(t *testing.T) {
    dbHandler := tests.NewTestDBConnectionsHandlerLocal()

    cdl := map[string]struct {
        LenExpected int
        Expecteds   []struct {
            CompanyNameExpected      string
            CustomerDocumentExpected string
        }
    }{
        "51537476467": {LenExpected: 2, Expecteds: []struct {
            CompanyNameExpected      string
            CustomerDocumentExpected string
        }{{CompanyNameExpected: "ABC S.A.", CustomerDocumentExpected: "51537476467"},
            {CompanyNameExpected: "123 S.A.", CustomerDocumentExpected: "51537476467"}}},
        "26658236674": {LenExpected: 1, Expecteds: []struct {
            CompanyNameExpected      string
            CustomerDocumentExpected string
        }{{CompanyNameExpected: "DBZ S.A.", CustomerDocumentExpected: "26658236674"}}},
        "62824334010": {LenExpected: 1, Expecteds: []struct {
            CompanyNameExpected      string
            CustomerDocumentExpected string
        }{{CompanyNameExpected: "XPTO S.A.", CustomerDocumentExpected: "62824334010"}}},
    }

    // new negativations controller handler
    h := NewHandler(dbHandler)

    for cd, exp := range cdl {
        ns, err := h.GetByCustomerDocument(cd)
        if err != nil {
            t.Error(err)
        } else {
            assert.Equal(t, exp.LenExpected, len(*ns))
            for _, e := range exp.Expecteds {
                customerDocExists := false
                customerNameExists := false
                for _, n := range *ns {
                    if e.CustomerDocumentExpected == n.CustomerDocument {
                        customerDocExists = true
                    }
                    if e.CompanyNameExpected == n.CompanyName {
                        customerNameExists = true
                    }
                }
                if !customerDocExists || !customerNameExists {
                    t.Log("Customer document "+ e.CustomerDocumentExpected + " and Company Name " + e.CompanyNameExpected + " are expected, put not appears in results")
                }
                assert.Equal(t, true, customerDocExists && customerNameExists)
            }
        }
    }

    cdlError := map[string]struct {
        ErrorExpected error
    }{
        "asdf": {ErrorExpected: gorm.ErrRecordNotFound},
    }

    for cd, exp := range cdlError {
        _, err := h.GetByCustomerDocument(cd)
        assert.IsEqual(exp.ErrorExpected, err)
    }
}
