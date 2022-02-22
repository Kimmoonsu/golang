package main

import (
	"fmt"
	"net/url"
	"server/autocoin/model"
	"server/autocoin/model/trade"
	"server/autocoin/model/trade/order"
)

// PurchaseOrder 매수하기
//
// [BODY PARAMS]
//
// market : REQUIRED. Market ID
//
// side : REQUIRED. 주문 종류
//
// volume : REQUIRED. 주문 수량. 지정가, 시장가 매도 시 필수
//
// price : REQUIRED. 유닛당 주문 가격. 지정가, 시장가 매수 시 필수
// ex) KRW-BTC 마켓에서 1BTC당 1,000 KRW로 거래할 경우, 값은 1000 이 된다.
// ex) KRW-BTC 마켓에서 1BTC당 매도 1호가가 500 KRW 인 경우, 시장가 매수 시 값을 1000으로 세팅하면 2BTC가 매수된다.
// (수수료가 존재하거나 매도 1호가의 수량에 따라 상이할 수 있음)
//
// orderType : REQUIRED. 주문 타입
//
// identifier : 조회용 사용자 지정 값
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) PurchaseOrder(market, volume, price, orderType, identifier string) (*order.Order, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	if len(price) == 0 {
		return nil, nil, fmt.Errorf("price length is 0")
	}

	switch orderType {
	case trade.ORDER_TYPE_LIMIT:
	case trade.ORDER_TYPE_PRICE:
	default:
		return nil, nil, fmt.Errorf("invalid orderType")
	}

	api, e := GetApiInfo(FuncPurchaseOrder)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market":     []string{market},
		"side":       []string{trade.ORDER_SIDE_BID},
		"volume":     []string{volume},
		"price":      []string{price},
		"ord_type":   []string{orderType},
		"identifier": []string{identifier},
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, values, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	order, e := order.OrderFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return order, model.RemainingFromHeader(resp.Header), nil
}

// SellOrder 매도하기
//
// [BODY PARAMS]
//
// market : REQUIRED. Market ID
//
// side : REQUIRED. 주문 종류
//
// volume : REQUIRED. 주문 수량. 지정가, 시장가 매도 시 필수
//
// price : REQUIRED. 유닛당 주문 가격. 지정가, 시장가 매수 시 필수
// ex) KRW-BTC 마켓에서 1BTC당 1,000 KRW로 거래할 경우, 값은 1000 이 된다.
// ex) KRW-BTC 마켓에서 1BTC당 매도 1호가가 500 KRW 인 경우, 시장가 매수 시 값을 1000으로 세팅하면 2BTC가 매수된다.
// (수수료가 존재하거나 매도 1호가의 수량에 따라 상이할 수 있음)
//
// orderType : REQUIRED. 주문 타입
//
// identifier : 조회용 사용자 지정 값
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) SellOrder(market, volume, price, orderType, identifier string) (*order.Order, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	if len(volume) == 0 {
		return nil, nil, fmt.Errorf("volume length is 0")
	}

	switch orderType {
	case trade.ORDER_TYPE_LIMIT:
	case trade.ORDER_TYPE_MARKET:
	default:
		return nil, nil, fmt.Errorf("invalid orderType")
	}

	api, e := GetApiInfo(FuncSellOrder)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market":     []string{market},
		"side":       []string{trade.ORDER_SIDE_ASK},
		"volume":     []string{volume},
		"price":      []string{price},
		"ord_type":   []string{orderType},
		"identifier": []string{identifier},
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, values, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	order, e := order.OrderFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return order, model.RemainingFromHeader(resp.Header), nil
}
