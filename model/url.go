package model

type Urls struct {
	List []string `json:"urls" binding:"required"`
}
