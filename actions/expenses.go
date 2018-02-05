package actions

import (
	"github.com/BryanMoslo/go_finanzly/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// Model: Singular (Expense)
// DB Table: Plural (expenses)
// Resource: Plural (Expenses)
// Path: Plural (/expenses)
// View Template Folder: Plural (/templates/expenses/)

// ExpensesResource is the resource for the Expense model
type ExpensesResource struct {
	buffalo.Resource
}

// GET /expenses
func (v ExpensesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	expenses := &models.Expenses{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Expenses from the DB
	if err := q.All(expenses); err != nil {
		return errors.WithStack(err)
	}

	// Make Expenses available inside the html template
	c.Set("expenses", expenses)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("expenses/index.html"))
}

// Show gets the data for one Expense. This function is mapped to
// the path GET /expenses/{expense_id}
func (v ExpensesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Expense
	expense := &models.Expense{}

	// To find the Expense the parameter expense_id is used.
	if err := tx.Find(expense, c.Param("expense_id")); err != nil {
		return c.Error(404, err)
	}

	// Make expense available inside the html template
	c.Set("expense", expense)

	return c.Render(200, r.HTML("expenses/show.html"))
}

// New renders the form for creating a new Expense.
// This function is mapped to the path GET /expenses/new
func (v ExpensesResource) New(c buffalo.Context) error {
	// Make expense available inside the html template
	c.Set("expense", &models.Expense{})

	return c.Render(200, r.HTML("expenses/new.html"))
}

// Create adds a Expense to the DB. This function is mapped to the
// path POST /expenses
func (v ExpensesResource) Create(c buffalo.Context) error {
	// Allocate an empty Expense
	expense := &models.Expense{}

	// Bind expense to the html form elements
	if err := c.Bind(expense); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(expense)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make expense available inside the html template
		c.Set("expense", expense)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("expenses/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Expense was created successfully")

	// and redirect to the expenses index page
	return c.Redirect(302, "/expenses/%s", expense.ID)
}

// Edit renders a edit form for a Expense. This function is
// mapped to the path GET /expenses/{expense_id}/edit
func (v ExpensesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Expense
	expense := &models.Expense{}

	if err := tx.Find(expense, c.Param("expense_id")); err != nil {
		return c.Error(404, err)
	}

	// Make expense available inside the html template
	c.Set("expense", expense)
	return c.Render(200, r.HTML("expenses/edit.html"))
}

// Update changes a Expense in the DB. This function is mapped to
// the path PUT /expenses/{expense_id}
func (v ExpensesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Expense
	expense := &models.Expense{}

	if err := tx.Find(expense, c.Param("expense_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Expense to the html form elements
	if err := c.Bind(expense); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(expense)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make expense available inside the html template
		c.Set("expense", expense)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("expenses/edit.html"))
	}

	// and redirect to the expenses index page
	return c.Redirect(201, "/expenses/%s", expense.ID)
}

func PaidExpense(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	expense := &models.Expense{}

	err := tx.Find(expense, c.Param("expense_id"))

	if err = c.Bind(expense); err != nil {
		return nil
	}

	_, err = tx.ValidateAndUpdate(expense)
	if err != nil {
		return nil
	}

	tx.Update(expense)
	return nil
}

// Destroy deletes a Expense from the DB. This function is mapped
// to the path DELETE /expenses/{expense_id}
func (v ExpensesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Expense
	expense := models.Expense{}

	// To find the Expense the parameter expense_id is used.
	if err := tx.Find(&expense, c.Param("expense_id")); err != nil {
		return c.Error(404, err)
	}

	boardID := expense.BoardID

	if err := tx.Destroy(&expense); err != nil {
		return errors.WithStack(err)
	}

	// Redirect to the boards list page
	return c.Redirect(302, "/boards/%s", boardID)
}
