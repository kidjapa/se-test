package legacy

import (
    "encoding/json"
    "fmt"
    "github.com/thoas/go-funk"
    "se-test/config"
    "se-test/db"
    "se-test/db/models"
    "se-test/utils"
    "strings"
    "time"
)

const (
    chunkInsertSize = 500
)

type Handler struct {
    DbHandler *db.Handler
    Config    config.Configuration
}

func NewHandler(db *db.Handler) *Handler {
    return &Handler{
        Config:    db.Config,
        DbHandler: db,
    }
}

func (h *Handler) GetAllNegativations() (n *[]models.Negativations, err error) {

    httpRequestHelper := &utils.HttpRequestHelper{}

    url := httpRequestHelper.GetUrl(h.Config.LegacyAccess.Url+"/negativations", nil)
    netClient := httpRequestHelper.GetNetClient(time.Second * 30)
    req := httpRequestHelper.GetRequestBodyApi("GET", url, nil, utils.Header{
        "Content-type": {"application/json"},
    })

    bodyData, err := httpRequestHelper.Do(req, netClient)
    if err != nil {
        return
    }

    n = &[]models.Negativations{}
    err = json.Unmarshal(bodyData, n)
    if err != nil {
        return
    }
    return
}

func (h *Handler) SaveNegativations(ns *[]models.Negativations) (err error) {

    // Chunk size by 500 inserts at time
    chunkList := funk.Chunk(*ns, chunkInsertSize)

    tx := h.DbHandler.Conn.Begin()

    for _, chunk := range chunkList.([][]models.Negativations) {
        var valueStrings []string
        var valueArgs []interface{}
        for _, neg := range chunk {
            valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?)")
            valueArgs = append(valueArgs, neg.CompanyDocument)  // Company Document
            valueArgs = append(valueArgs, neg.CompanyName)      // Company Name
            valueArgs = append(valueArgs, neg.CustomerDocument) // Customer Document
            valueArgs = append(valueArgs, neg.Value)            // Value
            valueArgs = append(valueArgs, neg.Contract)         // Contract
            valueArgs = append(valueArgs, neg.DebtDate)         // DebtDate
            valueArgs = append(valueArgs, neg.InclusionDate)    // Inclusion Date
        }

        smt := fmt.Sprintf("INSERT INTO se_test.negativations (company_document, company_name, customer_document, value, contract, debtDate, inclusion_date) VALUES %s",
            strings.Join(valueStrings, ","))
        err = tx.Exec(smt, valueArgs...).Error
        if err != nil {
            tx.Rollback()
            fmt.Println("Error when inserting new contracts (SaveNegativations): " + err.Error())
            return
        }
    }

    err = tx.Commit().Error
    if err != nil {
        fmt.Println("Error when inserting new contracts (SaveNegativations): " + err.Error())
        return
    }
    return
}
