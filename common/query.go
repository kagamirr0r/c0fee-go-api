// common/query.go
package common

type QueryParams struct {
	Limit    int    `query:"limit"`
	Cursor   uint   `query:"cursor"`
	NameLike string `query:"name_like"`
}
