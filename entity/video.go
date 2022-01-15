package entity

type Video struct {
	Title       string `json:"title" binding:"min=3,max=100" validate:"has-space"`
	Description string `json:"description" binding:"max=1000"`
	URL         string `json:"url" binding:"url"`
	Author      Person `json:"author" binding:"required"`
}
