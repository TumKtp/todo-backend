package todo

type TodoService interface {
	ListTodos(sort, title, description string) ([]*Todo, error)
	CreateNewTodo(todo *TodoRequest) (*Todo, error)
	UpdateTodo(id string, todo *TodoRequest) (*Todo, error)
}

type TodoRepository interface {
	GetTodos(sort, title, description string) ([]*Todo, error)
	SaveTodo(todo *TodoRequest) (*Todo, error)
	UpdateTodo(id string, todo *TodoRequest) (*Todo, error)
}
