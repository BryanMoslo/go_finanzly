package grifts

import (
	"github.com/BryanMoslo/go_finanzly/models"
	"github.com/markbates/grift/grift"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

var _ = grift.Add("db:seed", func(c *grift.Context) error {
	models.DB.Transaction(func(tx *pop.Connection) error {
		board := &models.Board{Name: "Tablero Ejemplo"}
		err := tx.Create(board)

		if err != nil {
			return errors.WithStack(err)
		}

		income := &models.Income{Name: "Ejemplo de Entrada", BoardID: board.ID}
		return tx.Create(income)
	})

	return nil
})
