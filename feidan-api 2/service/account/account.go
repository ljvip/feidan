package account

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"feidan-api/global"
	"feidan-api/model/account"
	"fmt"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"sync/atomic"
)

type AccountService struct{}

var AccountServiceApp = &AccountService{}

type (
	CeLottery struct {
		AccountType    int      `json:"accountType"`
		Balls          []string `json:"balls"`
		CanBack        bool     `json:"canBack"`
		DrawCount      int      `json:"drawCount"`
		ID             string   `json:"id"`
		LotterySubType int      `json:"lotterySubType"`
		MaxBall        int      `json:"maxBall"`
		MinBall        int      `json:"minBall"`
		Name           string   `json:"name"`
		Repeatable     bool     `json:"repeatable"`
		SortResult     bool     `json:"sortResult"`
		Tb             int      `json:"tb"`
		Template       string   `json:"template"`
		Type           int      `json:"type"`
	}
	CeUser struct {
		GameEnable    bool   `json:"gameEnable"`
		Username      string `json:"username"`
		WechatEnabled int    `json:"wechatEnabled"`
	}
	CeAccount struct {
		Balance  float64 `json:"balance"`
		Betting  float64 `json:"betting,omitempty"`
		MaxLimit float64 `json:"maxLimit"`
		Type     int     `json:"type"`
		Result   float64 `json:"result"`
	}
	CeUserInfo struct {
		Accounts []*CeAccount `json:"accounts"`
		Atypes   []int        `json:"atypes"`
		//Lotterys any          `json:"lotterys"`
		User *CeUser `json:"user"`
	}
)

func (s *AccountService) GetToken(ctx context.Context, url, username, password string) (string, error) {
	captchImageData, cryptograph, err := CaptchaServiceApp.GetCaptchaImage(ctx, url)
	if err != nil {
		return "", err
	}

	vcode, err := CaptchaServiceApp.GetCaptchaData(ctx, captchImageData)
	if err != nil {
		return "", err
	}

	loginUrl := url + "/web/rest/login"
	params := map[string]any{
		"code":        vcode,
		"cryptograph": cryptograph,
		"username":    username,
		"password":    password,
	}
	paramsBytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(loginUrl, "application/json", bytes.NewBuffer(paramsBytes))
	if err != nil {
		return "", err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var loginRsp struct {
		StatusCode int    `json:"statusCode"`
		Status     string `json:"status"`
		Token      string `json:"token"`
		Message    string `json:"message"`
	}
	if err := json.Unmarshal(all, &loginRsp); err != nil {
		return "", err
	}
	if loginRsp.StatusCode != 0 || loginRsp.Status != "success" {
		msg := "登陆失败"
		if loginRsp.Message != "" {
			msg = loginRsp.Message
		}
		return "", errors.New(msg)
	}

	return loginRsp.Token, nil
}

func (s *AccountService) GetUserInfo(ctx context.Context, pUrl, token string) (*CeUserInfo, error) {
	infoUrl := pUrl + "/web/rest/member/info"

	req, err := http.NewRequest(http.MethodGet, infoUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Token", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var infoRsp struct {
		StatusCode int         `json:"statusCode"`
		Status     string      `json:"status"`
		Result     *CeUserInfo `json:"result"`
		Message    string      `json:"message"`
	}
	if err := json.Unmarshal(all, &infoRsp); err != nil {
		return nil, err
	}
	if infoRsp.StatusCode != 0 || infoRsp.Status != "success" {
		msg := "登陆失败"
		if infoRsp.Message != "" {
			msg = infoRsp.Message
		}
		return nil, errors.New(msg)
	}

	return infoRsp.Result, nil
}

func (s *AccountService) GetList(ctx context.Context, pUrl, token string) (any, error) {
	infoUrl := pUrl + "/web/rest/member/history"

	req, err := http.NewRequest(http.MethodGet, infoUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Token", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var infoRsp struct {
		StatusCode int    `json:"statusCode"`
		Status     string `json:"status"`
		Result     any    `json:"result"`
		Message    string `json:"message"`
	}
	if err := json.Unmarshal(all, &infoRsp); err != nil {
		return "", err
	}
	if infoRsp.StatusCode != 0 || infoRsp.Status != "success" {
		msg := "获取报表失败"
		if infoRsp.Message != "" {
			msg = infoRsp.Message
		}
		return nil, errors.New(msg)
	}

	return infoRsp, nil
}

func (s *AccountService) AutoLoginConcurrency(ctx context.Context) (string, error) {
	var ps []*account.Platform
	err := global.GVA_DB.Find(&ps).Order("polling Asc").Error
	if err != nil {
		return "", err
	}

	var (
		accChannel = make(chan *account.Platform, 1024)
	)
	var countLogin atomic.Int32

	var eg errgroup.Group
	for _, p := range ps {
		eg.Go(func() error {
			_, err := s.FetchUserInfo(ctx, p)
			if err != nil {
				token, err := s.GetToken(ctx, p.Url, p.Username, p.Password)
				if err != nil {
					global.GVA_LOG.Error(fmt.Sprintf("平台:%s 账号:%s 密码:%s 自动登陆失败", p.Url, p.Username, p.Password), zap.Error(err))
					return nil
				}
				p.Token = token
				_, err = s.FetchUserInfo(ctx, p)
				countLogin.Add(1)
			}
			if err == nil {
				accChannel <- p
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return "", err
	}

	var data string
	onlineCount := len(accChannel)
	for i := 0; i < onlineCount; i++ {
		v := <-accChannel
		data += "平台名称: <a href=\"" + v.Url + "\">" + v.PlatformName + "</a><br>用户名: " + v.Username + "<hr>"
	}

	return fmt.Sprintf("全部账户: %d<br>本次登录: %d<br>在线账户: %d<br><br>%s", len(ps), countLogin.Load(), onlineCount, data), nil
}

func (s *AccountService) AutoLogin(ctx context.Context) (string, error) {
	var ps []*account.Platform
	err := global.GVA_DB.Find(&ps).Order("polling Asc").Error
	if err != nil {
		return "", err
	}

	var (
		acc []*account.Platform
	)
	var countLogin int32

	for _, p := range ps {
		_, err := s.FetchUserInfo(ctx, p)
		if err != nil {
			var token string
			token, err = s.Login(ctx, p)
			if err != nil {
				global.GVA_LOG.Error(fmt.Sprintf("平台:%s 账号:%s 密码:%s 自动登陆失败", p.Url, p.Username, p.Password), zap.Error(err))
				continue
			}
			p.Token = token
			_, err = s.FetchUserInfo(ctx, p)
			if err != nil {
				countLogin++
			}
		}
		if err == nil {
			acc = append(acc, p)
		}
	}

	var data string
	onlineCount := len(acc)
	for i := 0; i < onlineCount; i++ {
		v := acc[i]
		data += "平台名称: <a href=\"" + v.Url + "\">" + v.PlatformName + "</a><br>用户名: " + v.Username + "<hr>"
	}

	return fmt.Sprintf("全部账户: %d<br>本次登录: %d<br>在线账户: %d<br><br>%s", len(ps), countLogin, onlineCount, data), nil
}

func (s *AccountService) Login(ctx context.Context, p *account.Platform) (string, error) {
	token, err := s.GetToken(ctx, p.Url, p.Username, p.Password)
	if err != nil {

		errr := global.GVA_DB.Model(&account.Platform{}).Where("id = ?", p.ID).Updates(map[string]any{
			"online":         0,
			"offline_reason": err.Error(),
		}).Error
		if errr != nil {
			return "", errr
		}

		return "", err
	}

	err = global.GVA_DB.Model(&account.Platform{}).Where("id = ?", p.ID).Updates(map[string]any{
		"online":         1,
		"offline_reason": "",
		"token":          token,
	}).Error
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AccountService) FetchUserInfo(ctx context.Context, p *account.Platform) (*CeUserInfo, error) {
	info, err := s.GetUserInfo(ctx, p.Url, p.Token)
	if err != nil {
		return nil, err
	}

	if len(info.Accounts) > 0 {
		accInfo := info.Accounts[0]

		balance := cast.ToString(accInfo.Balance)
		if balance == "0" {
			balance = "0.00"
		}
		betting := cast.ToString(accInfo.Betting)
		if betting == "0" {
			betting = "0.00"
		}
		result := cast.ToString(accInfo.Result)
		if result == "0" {
			result = "0.00"
		}

		err := global.GVA_DB.Model(&account.Platform{}).Where("id = ? or token = ?", p.ID, p.Token).Updates(map[string]any{
			"balance": balance,
			"betting": betting,
			"result":  result,
		}).Error
		if err != nil {
			return nil, err
		}
	}

	return info, nil
}
