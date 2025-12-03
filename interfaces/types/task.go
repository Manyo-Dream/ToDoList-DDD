package types

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

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type TaskList struct {
}

type TaskDetail struct {
}

type TaskUpdate struct {
}

type TaskSearch struct {
}

type TaskDetele struct {
}
