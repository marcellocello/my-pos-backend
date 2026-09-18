package model

type Unit struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type CreateUnitRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required,min=2"`
	Description *string `json:"description,omitempty" binding:"min=5,max=50"`
}

type UpdateUnitRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required,min=2"`
	Description *string `json:"description,omitempty" binding:"min=5,max=50"`
}
