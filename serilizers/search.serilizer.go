package serilizers

type SearchReques struct {
	Product			string			`json:"product" form:"product" binding:"required,min=1"`
}