package types

type UserReq struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	PassWord string `json:"pass_word" form:"pass_word" binding:"required"`
}

type UserResp struct {
	ID        int64  `json:"id"`
	UserName  string `json:"user_name"`
	CreatedAt int64  `json:"created_at"`
}
