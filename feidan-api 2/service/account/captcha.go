package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CaptchaService struct {
	client *http.Client
}

var (
	ErrGetCaptchaData = errors.New("识别登录验证码失败")
)

var CaptchaServiceApp = &CaptchaService{client: &http.Client{}}

func (c *CaptchaService) GetCaptchaImage(ctx context.Context, url string) (string, string, error) {
	captchaApi := url + "/web/rest/generatecaptcha"
	resp, err := c.client.Get(captchaApi)
	if err != nil {
		return "", "", err
	}

	var captchaRsp struct {
		CaptchImageData string `json:"captchImageData"`
		Cryptograph     string `json:"cryptograph"`
	}

	all, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(all, &captchaRsp); err != nil {
		return "", "", err
	}

	return captchaRsp.CaptchImageData, captchaRsp.Cryptograph, nil
}

func (c *CaptchaService) GetCaptchaData(ctx context.Context, captcha string) (string, error) {
	base64Image := fmt.Sprintf("data:image/png;base64,%s", captcha)

	formData := url.Values{
		"v_pic":  {base64Image},
		"pri_id": {"dn"},
	}

	req, err := http.NewRequest(http.MethodPost, "https://yzmcolor.market.alicloudapi.com/yzmSpeed", strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("APPCODE %s", "f330b3b39f1d414390e3765b29e60b1a"))

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var dataRsp struct {
		Vcode   string `json:"v_code"`
		ErrCode int    `json:"errCode"`
	}
	if err := json.Unmarshal(body, &dataRsp); err != nil {
		return "", err
	}
	if dataRsp.ErrCode != 0 {
		return "", ErrGetCaptchaData
	}

	return dataRsp.Vcode, nil

}
