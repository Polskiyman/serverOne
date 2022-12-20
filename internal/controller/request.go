package controller

type CreateRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MakeFriendsRequest struct {
	SourceId int `json:"source_Id"`
	TargetId int `json:"target_Id"`
}

type DeleteUserRequest struct {
	UserId int `json:"user_id"`
}

type UpdateAgeRequest struct {
	UserAge int `json:"user_age"`
}
