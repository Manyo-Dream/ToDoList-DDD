package types

type List[T any] struct {
	Count int64 `json:"count"`
	Items []T   `json:"items"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type TaskResp struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

type TaskCreateReq struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

type TaskListReq struct {
	Pagination
}

type TaskDetailReq struct {
	ID uint `json:"id" form:"id"`
}

type TaskUpdateReq struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

type TaskSearchReq struct {
	Info string `json:"info" form:"info"`
	Pagination
}

type TaskDeteleReq struct {
	ID uint `json:"id" form:"id"`
}
