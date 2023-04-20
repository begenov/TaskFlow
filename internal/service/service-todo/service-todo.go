package servicetodo

type todoProvider interface{}

type ServiceTodo struct {
	todo todoProvider
}

func NewServiceTodo(todo todoProvider) *ServiceTodo {
	return &ServiceTodo{
		todo: todo,
	}
}
