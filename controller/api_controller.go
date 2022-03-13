package controller

import (
	"encoding/json"
	"net/http"
    "strconv"
    "path"

	"myapi/controller/dto"
	"myapi/model/repository"
    "myapi/model/entity"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
    PostTodo(w http.ResponseWriter, r *http.Request)
    PutTodo(w http.ResponseWriter, r *http.Request)
    DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) *todoController {
	return &todoController{tr}
}

//取得
func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
    //repositoryパッケージで作ったGetTodosメソッドでデータ取得
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}
    // 取得したTODOのentityをDTOに詰め替え
	var todoResponses []dto.TodoResponse //スライス定義
	for _, v := range todos {
        todoResponses = append(todoResponses, dto.TodoResponse{Id: v.Id, Account: v.Account, Name: v.Name, Passwd: v.Passwd, Created: v.Created})
	}

	var todosResponse dto.TodosResponse
	todosResponse.Todos = todoResponses
    //構造体をjsonに変換
	output, _ := json.MarshalIndent(todosResponse.Todos, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

//特定取得
func (tc *todoController) GetTodoName(w http.ResponseWriter, r *http.Request) {
    Id, err := strconv.Atoi(path.Base(r.URL.Path))
    if err != nil {
        w.WriteHeader(500)
        return
    }
    todoId := entity.TodoEntity.Id
    todoId = Id
    todo, err := tc.tr.GetTodoId(todoId)
    if err != nil {
        w.WriteHeader(500)
        return
    }
    todoResponses := dto.TodoResponse{Id: todoId, Account: todo.Account, Name: todo.Name, Passwd: todo.Passwd, Created: todo.Created}
    output, _ := json.MarshalIndent(todoResponses, "", "\t")

    w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

//挿入
func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
   body := make([]byte, r.ContentLength)
   r.Body.Read(body)
   var todoRequest dto.TodoRequest
   json.Unmarshal(body, &todoRequest)

   todo := entity.TodoEntity{Account: todoRequest.Account, Name: todoRequest.Name, Passwd: todoRequest.Passwd}
   id, err := tc.tr.InsertTodo(todo)
   if err != nil {
       w.WriteHeader(500)
	   return
   }
    w.Header().Set("Location", r.Host+r.URL.Path+strconv.FormatInt(id, 10))
    //r.URL.Path: /api/users/
    w.WriteHeader(201)
}

//更新
func (tc *todoController) PutTodo(w http.ResponseWriter, r *http.Request){
     todoId, err := strconv.Atoi(path.Base(r.URL.Path))
     if err != nil{
        w.WriteHeader(400)
        return
     }
     body := make([]byte, r.ContentLength)
     r.Body.Read(body)
     var todoRequest dto.TodoRequest
     json.Unmarshal(body, &todoRequest)

     todo := entity.TodoEntity{Id: todoId, Account: todoRequest.Account, Name: todoRequest.Name, Passwd: todoRequest.Passwd}
     err = tc.tr.UpdateTodo(todo)
     if err != nil {
         w.WriteHeader(500)
         return
     }
     w.WriteHeader(200)
}

//削除
func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request){
    todoId, err := strconv.Atoi(path.Base(r.URL.Path))
    if err != nil {
        w.WriteHeader(400)
        return
    }
    err = tc.tr.DeleteTodo(todoId)
    if err != nil {
        w.WriteHeader(500)
        return
    }
    w.WriteHeader(204)
}
