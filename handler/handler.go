package handler

import "se-test/interfaces"

type Handler struct {
    NegativationsStore interfaces.NegativationsInterface
}

func NewHandler(n interfaces.NegativationsInterface) *Handler {
    return &Handler{
        NegativationsStore: n,
    }
}