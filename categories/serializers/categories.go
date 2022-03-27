package categoriesSerializers

type CategoriesResponse struct {
	Category string `form:"category"`
}

type CategoriesResponseArray struct {
	Categories []CategoriesResponse `form:"categories"`
}
