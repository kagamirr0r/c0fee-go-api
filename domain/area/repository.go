package area

type IAreaRepository interface {
	GetById(area *Entity, id uint) error
	List(areas *[]Entity) error
}
