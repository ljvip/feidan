package bet

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"feidan-api/global"
	accountCommomModel "feidan-api/model/account"
	betReq "feidan-api/model/bet/request"
	betRsp "feidan-api/model/bet/response"
	"feidan-api/model/common"
	"feidan-api/service/account"
	"feidan-api/utils"
	"fmt"
	"github.com/spf13/cast"
	"golang.org/x/sync/syncmap"
	"gorm.io/gorm"
	"io"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"
)

type BetService struct{}

var BetServiceApp = &BetService{}

var inUseUser = syncmap.Map{}

type (
	SendBetData struct {
		Lottery    string     `json:"lottery"`
		DrawNumber string     `json:"drawNumber"`
		Bets       []*BetData `json:"bets"`
	}
	BetData struct {
		Amount   string `json:"amount"`
		Odds     any    `json:"odds"`
		Game     string `json:"game"`
		Contents string `json:"contents"`
		Multiple int32  `json:"multiple,omitempty"`
		State    int32  `json:"State,omitempty"`
		Title    string `json:"title,omitempty"`
	}
)

func (s *BetService) OutFly(ctx context.Context, req *betReq.OutFlyReq) (*betRsp.BetRspData, error) {
	brd := &betRsp.BetRspData{
		State:   0,
		Success: make([]*betRsp.BetResult, 0),
		Failure: make([]*betRsp.BetResult, 0),
	}
	if len(req.Data) == 0 || req.Token != "f3785tg3b48f237fg8243yt5" {
		return brd, nil
	}

	var adminIds []int32
	ces := strings.Split(req.Ce, ",")
	err := global.GVA_DB.Model(&accountCommomModel.AdminUser{}).
		Where("username in ?", ces).
		Pluck("id", &adminIds).Error
	if err != nil {
		return nil, err
	}
	var users []*accountCommomModel.Platform
	q := global.GVA_DB.Where("admin_user_id in (?)", adminIds)
	for _, ce := range ces {
		q = q.Or("FIND_IN_SET(?, ce) > 0", ce)
	}
	err = q.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var sum float64
	for _, d := range req.Data {
		for _, l := range d.List {
			sum += cast.ToFloat64(l.Amount)
		}
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Polling < users[j].Polling
	})

	var hasUser bool
	for _, user := range users {
		_, ok := inUseUser.Load(user.ID)
		if ok {
			continue
		}
		inUseUser.Store(user.ID, struct{}{})

		info, err := account.AccountServiceApp.FetchUserInfo(ctx, user)
		if err != nil || len(info.Accounts) == 0 {
			global.GVA_LOG.Error(fmt.Sprintf("账号:%s,刷新用户余额失败", user.Username))
			inUseUser.Delete(user.ID)
			continue
		}
		balance := info.Accounts[0].Balance

		tSum := sum
		redouble := cast.ToFloat64(user.Redouble) / 100
		if redouble != 1 {
			// 重新计算sum
			tSum = 0
			for _, d := range req.Data {
				for _, l := range d.List {
					tSum += math.Round(cast.ToFloat64(l.Amount) * redouble)
				}
			}
		}

		if balance < tSum {
			inUseUser.Delete(user.ID)
			continue
		}
		hasUser = true

		bReq := &betReq.BetReq{
			Token:   user.Token,
			Rows:    req.Rows,
			Ce:      req.Ce,
			Data:    req.Data,
			URL:     user.Url,
			Version: "v1",
		}
		for _, d := range bReq.Data {
			for _, l := range d.List {
				amount := math.Round(cast.ToFloat64(l.Amount) * redouble)
				l.Amount = cast.ToString(amount)
			}
		}
		send, _ := json.Marshal(&bReq)

		log := &accountCommomModel.PlatformLog{
			GVA_MODEL: global.GVA_MODEL{
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
			AdminUserId: user.AdminUserId,
			Url:         user.Url,
			Username:    user.Username,
			Send:        send,
			Input:       nil,
		}
		err = global.GVA_DB.Create(log).Error
		if err != nil {
			inUseUser.Delete(user.ID)
			continue
		}

		err = global.GVA_DB.Model(&accountCommomModel.Platform{}).Where("id = ?", user.ID).Update("polling", gorm.Expr("polling + ?", 1)).Error
		if err != nil {
			inUseUser.Delete(user.ID)
			continue
		}

		br, err := s.Bet(ctx, bReq)
		input, _ := json.Marshal(br)
		_ = global.GVA_DB.Model(&accountCommomModel.PlatformLog{}).Where("id = ?", log.ID).Update("input", string(input)).Error
		if err != nil {
			global.GVA_LOG.Error(fmt.Sprintf("账号:%s,飞单失败", user.Username))
			inUseUser.Delete(user.ID)
			continue
		}

		inUseUser.Delete(user.ID)

		return br.Data, nil
	}

	message := "飞单失败"
	if !hasUser {
		message = "无可用账户"
	}
	for _, d := range req.Data {
		for _, l := range d.List {
			l.Error = message
		}
		brd.Failure = append(brd.Failure, &betRsp.BetResult{
			Code:  d.Code,
			Issue: d.Issue,
			List:  d.List,
		})
	}
	return brd, nil

}

func (s *BetService) Bet(ctx context.Context, req *betReq.BetReq) (*betRsp.BetRsp, error) {

	rsp := &betRsp.BetRsp{
		Data: &betRsp.BetRspData{
			State:   0,
			Success: make([]*betRsp.BetResult, 0),
			Failure: make([]*betRsp.BetResult, 0),
		},
	}

	for _, d := range req.Data {
		gameName := utils.GetGameName(d.Code)
		odds, err := s.GetOdds(ctx, gameName, req.URL, req.Token)
		if err != nil {
			ret := &betRsp.BetResult{
				Code:  d.Code,
				Issue: d.Issue,
			}
			for _, l := range d.List {
				l.Error = err.Error()
				ret.List = append(ret.List, l)
			}
			rsp.Data.Failure = append(rsp.Data.Failure, ret)

			continue
		}

		sbd := &SendBetData{
			Lottery:    gameName,
			DrawNumber: d.Issue,
		}
		for _, l := range d.List {
			pumData := utils.GetPumInfo(cast.ToString(l.Pum))
			if pumData == nil {
				l.Error = "错误下注标识码"
				continue
			}

			pumInt := cast.ToInt(l.Pum)
			bd := &BetData{
				Amount:   cast.ToString(l.Amount),
				Odds:     odds[fmt.Sprintf("%s_%s", pumData.Game, pumData.Contents)],
				Game:     pumData.Game,
				Contents: pumData.Contents,
				Title:    pumData.Title,
				State:    pumData.State,
				Multiple: pumData.Multiple,
			}
			if (pumInt >= 105000 && pumInt <= 105900) || (pumInt >= 10600 && pumInt <= 106200) {
				bd.Odds = fmt.Sprintf("%s_%d", pumData.Game, pumData.State)
				bd.Contents = l.Log
			}
			sbd.Bets = append(sbd.Bets, bd)
		}
		err = s.Send(ctx, sbd, req.URL, req.Token)
		if err != nil {
			ret := &betRsp.BetResult{
				Code:  d.Code,
				Issue: d.Issue,
			}
			for _, l := range d.List {
				l.Error = err.Error()
				ret.List = append(ret.List, l)
			}
			rsp.Data.Failure = append(rsp.Data.Failure, ret)

			continue
		}

		failBr := &betRsp.BetResult{
			Code:  d.Code,
			Issue: d.Issue,
			List:  make([]*common.BetDataList, 0),
		}
		successBr := &betRsp.BetResult{
			Code:  d.Code,
			Issue: d.Issue,
			List:  make([]*common.BetDataList, 0),
		}
		for _, l := range d.List {
			if l.Error != "" {
				failBr.List = append(failBr.List, l)
			} else {
				successBr.List = append(successBr.List, l)
			}
		}
		if len(failBr.List) > 0 {
			rsp.Data.Failure = append(rsp.Data.Failure, failBr)
		}
		if len(successBr.List) > 0 {
			rsp.Data.Success = append(rsp.Data.Success, successBr)
		}
	}

	if len(rsp.Data.Failure) == 0 {
		rsp.Data.State = 1
	}

	info, err := account.AccountServiceApp.FetchUserInfo(ctx, &accountCommomModel.Platform{
		Url:   req.URL,
		Token: req.Token,
	})
	if err == nil {
		if len(info.Accounts) > 0 {
			acc := info.Accounts[0]
			rsp.Info = &betRsp.BetInfo{
				Balance:  acc.Balance,
				Betting:  acc.Betting,
				MaxLimit: acc.MaxLimit,
				Result:   acc.Result,
				Type:     int32(acc.Type),
			}
		}

	}
	return rsp, nil
}

func (s *BetService) GetOdds(ctx context.Context, name string, pUrl string, token string) (map[string]any, error) {
	infoUrl := fmt.Sprintf(pUrl+"/web/rest/member/odds?lottery=%s", name)

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
		StatusCode int            `json:"statusCode"`
		Status     string         `json:"status"`
		Result     map[string]any `json:"result"`
		Message    string         `json:"message"`
	}

	if err := json.Unmarshal(all, &infoRsp); err != nil {
		return nil, err
	}
	if infoRsp.StatusCode != 0 || infoRsp.Status != "success" {
		msg := "获取赔率失败"
		if infoRsp.Message != "" {
			msg = infoRsp.Message
		}
		return nil, errors.New(msg)
	}

	return infoRsp.Result, nil
}

func (s *BetService) Send(ctx context.Context, data *SendBetData, pUrl string, token string) error {
	infoUrl := pUrl + "/web/rest/member/placebet"

	marshal, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, infoUrl, bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}
	req.Header.Set("Token", token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var infoRsp struct {
		StatusCode int            `json:"statusCode"`
		Status     string         `json:"status"`
		Result     map[string]any `json:"result"`
		Message    string         `json:"message"`
	}

	if err := json.Unmarshal(all, &infoRsp); err != nil {
		return err
	}
	if infoRsp.StatusCode != 0 || infoRsp.Status != "success" {
		msg := "下注失败"
		if infoRsp.Message != "" {
			msg = infoRsp.Message
		}
		return errors.New(msg)
	}

	return nil
}
