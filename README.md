# goapi
`Just made a simple router, controller and model based on golang`

## Router and Controller
`goapi/src/routers/api.go `
```golang
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/books", controllers.ListBook)
	}

	return r
}
```

# Models
`goapi/src/models/book.go`
```golang
type Book struct {
	Name 	 string `json:"name"`
	Author 	 string `json:"author"`
	Category string `json:"category"`
}

func GetAllBook(b *[]Book) (err error) {
	if err = config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}
```
