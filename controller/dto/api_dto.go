package dto

type TodoResponse struct {
   Id       int     `json:"ID"`
   Account  string  `json:"ACCOUNT"`
   Name     string  `json:"NAME"`
   Passwd   string  `json:"PASSWORD"`
   Created   string  `json:"CREATED"`
}

type TodoRequest struct {
   Account  string  `json:"ACCOUNT"`
   Name     string  `json:"NAME"`
   Passwd   string  `json:"PASSWORD"`
}


type TodosResponse struct {
   Todos  []TodoResponse  `json:"Users"`
}

