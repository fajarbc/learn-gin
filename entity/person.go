package entity

type Person struct {
	Name  string `json:"name" binding:"required"`
	Age   uint   `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" validate:"email"`
}
