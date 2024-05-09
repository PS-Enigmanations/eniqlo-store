package request

type PaginationParams struct {
	Limit  string `form:"limit" default:"5"`
	Offset string `form:"offset" default:"0"`
}
