// Code generated by goctl. DO NOT EDIT.
package types

type AddReq struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type AddResp struct {
	Id int64 `json:"id"`
}

type RemoveReq struct {
	Id int64 `path:"id"`
}

type ModReq struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type FindByNameReq struct {
	Name string `form:"name"`
}

type FindByNameResp struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
}
