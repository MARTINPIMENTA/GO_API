package dependencies

type Container struct {
	database   string
	restClient string
}

func StartDependencies() *Container {
	return &Container{
		database:   "",
		restClient: "",
	}
}

func (container *Container) Database() string {
	return container.database
}

func (container *Container) RestyClient() string {
	return container.restClient
}
