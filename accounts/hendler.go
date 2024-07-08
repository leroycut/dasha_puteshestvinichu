package accounts

import (
	"net/http"
	"sync"

	"example.com/m/accounts/dto"
	"example.com/m/accounts/models"

	"github.com/labstack/echo/v4"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	if len(name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account doesn't exists")
	}

	delete(h.accounts, name)

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) PatchAccount_name(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account doesn't exists")
	}

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusConflict, "this name exists")
	}

	h.accounts[request.Name].Name = request.Name

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) PatchAccount_amount(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account doesn't exists")
	}

	h.accounts[request.Name].Amount = request.Amount

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}
