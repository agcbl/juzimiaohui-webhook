package impl

import (
	"github.com/fatelei/go-feishu/pkg/auth"
	"github.com/fatelei/juzimiaohui-webhook/configs"
	"github.com/fatelei/juzimiaohui-webhook/pkg/controller"
	"github.com/fatelei/juzimiaohui-webhook/pkg/dao/impl"
	"time"
)

type FeishuBotControllerImpl struct {
	auth *auth.Auth
}

var _ controller.FeishuBotController = (*FeishuBotControllerImpl)(nil)
var DefaultFeishuBotController controller.FeishuBotController

func NewFeishuBotController() *FeishuBotControllerImpl {
	auth := auth.NewAuth(configs.DefaultConfig.LarkBot.AppID, configs.DefaultConfig.LarkBot.AppSecret, configs.DefaultConfig.LarkBot.EndPoint)
	return &FeishuBotControllerImpl{auth: auth}
}

func (p *FeishuBotControllerImpl) GetAccessToken() string {
	record := impl.DefaultFeishuBotDAO.GetAccessToken()
	if record == nil {
		accessToken := p.auth.GetAccessToken()
		impl.DefaultFeishuBotDAO.Create(accessToken.Token, accessToken.Expire)
		return accessToken.Token
	} else {
		if time.Now().Unix() - record.Expire <= 10 {
			accessToken := p.auth.GetAccessToken()
			impl.DefaultFeishuBotDAO.Refresh(record.ID, accessToken.Token, accessToken.Expire)
			return accessToken.Token
		}
	}
	return record.Token
}
