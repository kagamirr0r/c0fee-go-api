package variety

type IVarietyRepository interface {
	List(varieties *[]Entity) error
}
