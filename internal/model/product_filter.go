package model

type ProductFilter struct {
	ID         int
	Search     string
	CategoryID int
	Page       int
	Size       int

	SortBy    string
	SortOrder string
}
