package actions

import (
	"github.com/BryanMoslo/go_finanzly/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Board)
// DB Table: Plural (boards)
// Resource: Plural (Boards)
// Path: Plural (/boards)
// View Template Folder: Plural (/templates/boards/)

// BoardsResource is the resource for the Board model
type BoardsResource struct {
	buffalo.Resource
}

// List gets all Boards. This function is mapped to the path
// GET /boards
func (v BoardsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	boards := &models.Boards{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Boards from the DB
	if err := q.All(boards); err != nil {
		return errors.WithStack(err)
	}

	// Make Boards available inside the html template
	c.Set("boards", boards)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("boards/index.html"))
}

// Show gets the data for one Board. This function is mapped to
// the path GET /boards/{board_id}
func (v BoardsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Board
	board := &models.Board{}

	// To find the Board the parameter board_id is used.
	if err := tx.Find(board, c.Param("board_id")); err != nil {
		return c.Error(404, err)
	}

	// Make board available inside the html template
	c.Set("board", board)

	incomes := &models.Incomes{}
	err := tx.BelongsTo(board).All(incomes)

	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("incomes", incomes)

	expenses := &models.Expenses{}
	err = tx.BelongsTo(board).All(expenses)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("expenses", expenses)

	boards := &models.Boards{}
	c.Set("boards", boards)

	return c.Render(200, r.HTML("boards/show.html"))
}

// New renders the form for creating a new Board.
// This function is mapped to the path GET /boards/new
func (v BoardsResource) New(c buffalo.Context) error {
	// Make board available inside the html template
	c.Set("board", &models.Board{})

	return c.Render(200, r.HTML("boards/new.html"))
}

// Create adds a Board to the DB. This function is mapped to the
// path POST /boards
func (v BoardsResource) Create(c buffalo.Context) error {
	// Allocate an empty Board
	board := &models.Board{}

	// Bind board to the html form elements
	if err := c.Bind(board); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(board)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make board available inside the html template
		c.Set("board", board)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("boards/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Board was created successfully")

	// and redirect to the boards index page
	return c.Redirect(302, "/boards/%s", board.ID)
}

// Edit renders a edit form for a Board. This function is
// mapped to the path GET /boards/{board_id}/edit
func (v BoardsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Board
	board := &models.Board{}

	if err := tx.Find(board, c.Param("board_id")); err != nil {
		return c.Error(404, err)
	}

	// Make board available inside the html template
	c.Set("board", board)
	return c.Render(200, r.HTML("boards/edit.html"))
}

// Update changes a Board in the DB. This function is mapped to
// the path PUT /boards/{board_id}
func (v BoardsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Board
	board := &models.Board{}

	if err := tx.Find(board, c.Param("board_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Board to the html form elements
	if err := c.Bind(board); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(board)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make board available inside the html template
		c.Set("board", board)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("boards/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Board was updated successfully")

	// and redirect to the boards index page
	return c.Redirect(302, "/boards/%s", board.ID)
}

// Destroy deletes a Board from the DB. This function is mapped
// to the path DELETE /boards/{board_id}
func (v BoardsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Board
	board := &models.Board{}

	// To find the Board the parameter board_id is used.
	if err := tx.Find(board, c.Param("board_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(board); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Board was destroyed successfully")

	// Redirect to the boards index page
	return c.Redirect(302, "/boards")
}
