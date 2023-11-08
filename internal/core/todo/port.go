package todo

type TodoService interface {
	ListTodos() ([]*Todo, error)
	CreateNewTodo(todo *TodoRequest) (*Todo, error)
	UpdateTodo(id string, todo *TodoRequest) (*Todo, error)
}

type TodoRepository interface {
	GetTodos() ([]*Todo, error)
	SaveTodo(todo *TodoRequest) (*Todo, error)
	UpdateTodo(id string, todo *TodoRequest) (*Todo, error)
}
