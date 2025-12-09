package types

type List[T any] struct {
	Count int64 `json:"count"`
	Items []T   `json:"items"`
}

type Pagination struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
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

type TaskDeleteReq struct {
	ID uint `json:"id" form:"id"`
}
