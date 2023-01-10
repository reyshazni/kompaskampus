package query_param

type UniversityQuery struct {
	Name  *string `query:"search"`
	Page  int     `query:"page"`
	Limit int     `query:"limit"`
}
