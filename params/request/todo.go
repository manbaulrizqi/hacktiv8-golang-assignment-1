package request

type CreateTodo struct {
	Title       string `example:"Title TODO"`
	Description string `example:"Desc TODO"`
}

type UpdateQueryParam struct {
	Id *int `query:"id"        example:"1"  default:"1"`
}
