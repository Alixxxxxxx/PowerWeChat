package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	*response.ResponseMiniProgram

	List     []*object.HashMap `json:"list"`
}
