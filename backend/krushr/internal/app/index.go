
func CreateApp(config *config.Config) (*App, error) {

}

func Initialize() *gin.Engine {
	r := gin.Default()
	handlers.InitHandlers(r)

	return r
}

func Run() {
	r := initRouter()
	r.Run(":8080")
}
