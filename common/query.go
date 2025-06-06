// common/query.go
package common

type QueryParams struct {
	Limit    int    `query:"limit"`
	NameLike string `query:"name_like"`
}
