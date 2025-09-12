package roaster

import "c0fee-api/common"

type IRoasterRepository interface {
	List(roasters *[]Entity) error
	Search(roasters *[]Entity, params common.QueryParams) error
}
