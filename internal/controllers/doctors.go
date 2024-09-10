package controllers

import (
	"mediversepreip/internal/models"

	"github.com/sev-2/raiden"
)

type BooksController struct {
	raiden.ControllerBase
	Http  string `path:"/doctors" type:"rest"`
	Model models.Doctors
}
