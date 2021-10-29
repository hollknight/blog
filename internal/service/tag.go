package service

// CountTagRequest /api/v1/tags [GET] Get
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// TagListRequest /api/v1/tags [GET] List
type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// CreateTagRequest /api/v1/tags [POST] Create
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// UpdateTagRequest /api/v1/tags/{id} [PUT] Update
type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

// DeleteTagRequest /api/v1/tags/{id} [DELETE] Delete
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
