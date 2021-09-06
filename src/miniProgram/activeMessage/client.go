package activeMessage

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	response2 "github.com/ArtisanCloud/power-wechat/src/kernel/response"
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/activeMessage/response"
)

type Client struct {
	*kernel.BaseClient
}

// 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func (comp *Client) GetDailyRetain(unionID string, openID string) (*response.ResponseActiveMessageCreateActiveID, error) {

	result := &response.ResponseActiveMessageCreateActiveID{}

	params := &object.StringMap{
		"unionid": unionID,
		"openid":  openID,
	}

	_, err := comp.HttpGet("cgi-bin/message/wxopen/activityid/create", params, nil, result)

	return result, err
}

// 修改被分享的动态消息。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html
func (comp *Client) BusinessLicense(activityID string, targetState int8, templateInfo *power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"activity_id":   activityID,
		"target_state":  targetState,
		"template_info": templateInfo,
	}

	_, err := comp.HttpPostJson("cgi-bin/message/wxopen/updatablemsg/send", data, nil, nil, result)

	return result, err
}
