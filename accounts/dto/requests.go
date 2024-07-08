package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type PatchAccountRequest struct {
	Name string `json:"name"`
}

type ChangeAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountRequest_name struct {
	Name     string `json:"name"`
	New_name string `json:"new_name"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
