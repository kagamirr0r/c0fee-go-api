package bean_rating

type IBeanRatingRepository interface {
	Create(beanRating *Entity) error
	GetByBeanID(beanRatings *[]Entity, beanID uint) error
	UpdateByID(beanRating *Entity) error
}
