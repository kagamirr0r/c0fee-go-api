package roast_level

type IRoastLevelRepository interface {
	GetAll(roastLevels *[]Entity) error
	GetById(roastLevel *Entity, id uint) error
}
