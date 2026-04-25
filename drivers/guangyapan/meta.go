package guangyapan

import (
	"github.com/OpenListTeam/OpenList/v4/internal/driver"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
)

type Addition struct {
	driver.RootID // 核心底层组件，保留不动
	PhoneNumber    string `json:"手机号" type:"text" help:"短信登录手机号，例如 +86 13800000000"`
	CaptchaToken   string `json:"人机验证Token" type:"text" help:"填写抓包获取的 Captcha Token (如果不需要可留空)"`
	SendCode       bool   `json:"发送验证码" type:"bool" help:"填好手机号后，开启此开关并【保存】即可发送短信验证码 (发送后会自动重置)"`
	VerifyCode     string `json:"短信验证码" type:"text" help:"收到短信后填入此处，再次【保存】即可完成最终登录"`
	VerificationID string `json:"验证ID (自动)" type:"text" help:"发送短信后系统自动生成，请勿手动修改"`
	AccessToken    string `json:"访问令牌 (自动)" type:"text" help:"登录成功后的授权凭证 (有 Refresh Token 时可留空)"`
	RefreshToken   string `json:"刷新令牌 (自动)" type:"text" help:"用于自动续期的 Refresh Token"`
	ClientID       string `json:"客户端ID" default:"aMe-8VSlkrbQXpUR" help:"默认客户端ID，一般保持默认即可"`
	DeviceID       string `json:"设备ID" help:"可选，自定义设备ID(32位字符)，留空则自动生成"`
	PageSize       int    `json:"分页大小" type:"number" default:"100"`
	OrderBy        int    `json:"排序字段" type:"number" default:"3" help:"0:文件名, 1:大小, 2:创建时间, 3:更新时间"`
	SortType       int    `json:"排序规则" type:"number" default:"1" help:"0:升序(从小到大), 1:降序(从大到小)"`
}

var config = driver.Config{
	Name:              "光鸭云盘",
	LocalSort:         false,
	OnlyProxy:         false,
	NoCache:           false,
	NoUpload:          false,
	NeedMs:            false,
	DefaultRoot:       "",
	CheckStatus:       true,
	// Alert 提示语也全部汉化，并分段让用户看得更清楚
	Alert:             "info|【光鸭云盘短信登录步骤】<br>1. 填写【手机号】(若需要则填人机验证Token)，开启【发送验证码】开关并点击底部保存；<br>2. 收到短信后，填写【短信验证码】，再次点击保存即可完成登录！系统会自动获取并保存令牌。",
	NoOverwriteUpload: true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &GuangYaPan{}
	})
}
