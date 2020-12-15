package handler

import "se-test/interfaces"

type Handler struct {
    NegativationsStore interfaces.NegativationsInterface
    LegacyStore        interfaces.LegacyInterface
}

func NewHandler(n interfaces.NegativationsInterface, l interfaces.LegacyInterface) *Handler {
    return &Handler{
        NegativationsStore: n,
        LegacyStore: l,
    }
}
