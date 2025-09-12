package farm

type IFarmRepository interface {
	GetById(farm *Entity, id uint) error
	List(farms *[]Entity) error
}
