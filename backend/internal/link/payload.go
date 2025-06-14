package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}
type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash,omitempty"`
}

type LinkGetAllParamsRequest struct {
	Limit  int `schema:"limit,default:10"`
	Offset int `schema:"offset,default:0"`
}
type LinkGetAllResponse struct {
	Count int64  `json:"count"`
	Links []Link `json:"links"`
}
