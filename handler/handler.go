package handler

import (
	"mailsender/repo"
)

type Handler struct {
	confirm *repo.Confirm
}

func InitHandler(c *repo.Confirm)*Handler{
	return &Handler{confirm: c}
}
