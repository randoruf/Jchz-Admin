package response

import "jchz-admin/model/system"

type Tags struct {
	ID       int64  `json:"id"`
	TagName  string `json:"tagname" `
	Articles int64  `json:"articles"`
}

type TagPageData struct {
	Total   int64  `json:"total"`
	PageNum int64  `json:"pagenum"`
	Tags    []Tags `json:"tags"`
}

type TagPageResponse struct {
	Data *TagPageData `json:"data"`
	Meta *Meta        `json:"meta"`
}

type AddTagResponse struct {
	TagData *system.Tag `json:"data"`
	Meta    *Meta       `json:"meta"`
}

type UpdateTagResponse struct {
	TagData *system.Tag `json:"data"`
	Meta    *Meta       `json:"meta"`
}
