package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"myapi/model/entity"

)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
    InsertTodo(todo entity.TodoEntity) (id int64, err error)
    UpdateTodo(todo entity.TodoEntity) (err error)
    DeleteTodo(id int)(err error)
}


type todoRepository struct {
}

//コンストラクタ (初期化のための関数) (戻り値に構造体のポインタ)
func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

//取得
func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	todos = []entity.TodoEntity{}
    //Query()メソッドを使ってrowsを受け取る
	rows, err := Db.Query("SELECT * FROM M_User")
	if err != nil {
		log.Print(err)
		return
	}
    //Scan()メソッドで値をコピーするために、Next()メソッドを使ってrowsの結果に対して1行1行繰り返し処理
	for rows.Next() {
		todo := entity.TodoEntity{}
		err = rows.Scan(&todo.Id, &todo.Account, &todo.Name, &todo.Passwd, &todo.Created)
		if err != nil {
			log.Print(err)
			return
		}
        //append関数を使ってスライスにデータを追加
		todos = append(todos, todo)
	}

	return
}


//特定取得
func (tr *todoRepository) GetTodoId(todoId entity.TodoEntity) (todos []entity.TodoEntity, err error){
    todos = []entity.TodoEntity{}
    row, err := Db.Query("SELECT ID = ? FROM M_User",todoId)
    todo := entity.TodoEntity{}
    err = row.Scan(&todoId, &todo.Account, &todo.Name, &todo.Passwd, &todo.Created)
    if err != nil {
        log.Print(err)
        return
    }
    todos = append(todos, todo)
    return
}


//挿入
func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id int64, err error) {
    stmt , err := Db.Prepare("INSERT INTO M_USER(ACCOUNT,NAME,PASSWORD,CREATED) VALUES(?,?,?,now())")
    if err != nil {
       return 0, err
    }
    defer stmt.Close()

    result, err := stmt.Exec(todo.Account, todo.Name, todo.Passwd)
    if err != nil {
       return 0, err
    }
    id, err = result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return
}

//更新
func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error){
   stmt, err := Db.Prepare("UPDATE M_User SET ACCOUNT = ?, NAME = ?, PASSWORD = ?  WHERE ID = ?")
   if err != nil {
      return err
   }
   defer stmt.Close()
   _, err = stmt.Exec(todo.Account, todo.Name, todo.Passwd, todo.Id)
   return
}

//削除
func (tr *todoRepository) DeleteTodo(id int) (err error){
   stmt, err := Db.Prepare("delete from M_User where ID = ?")
   if err != nil {
      return err
   }
   defer stmt.Close()
   _, err = stmt.Exec(id)
   return
}
