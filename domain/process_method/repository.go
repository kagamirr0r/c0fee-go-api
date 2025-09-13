package process_method

type IProcessMethodRepository interface {
	List(processMethods *[]Entity) error
}
