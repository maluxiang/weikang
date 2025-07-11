package logic

import "weikang/pkg"

type Payment interface {
	Pay(subject, outTradeNo string, amount float64) (string, error)
}

type AlipayPayment struct{}

func (a AlipayPayment) Pay(subject, outTradeNo string, amount float64) (string, error) {
	pay := pkg.NewAlipay{
		Subject:     subject,
		OutTradeNo:  outTradeNo,
		TotalAmount: amount,
	}
	url := pay.Alipay()
	return url, nil
}

type WechatPayment struct{}

func (w WechatPayment) Pay(subject, outTradeNo string, amount float64) (string, error) {
	// 这里调用微信支付SDK，生成支付链接
	url := "https://pay.wechat.com/xxx" // 真实项目需对接微信支付API
	return url, nil
}

type BankcardPayment struct{}

func (b BankcardPayment) Pay(subject, outTradeNo string, amount float64) (string, error) {
	// 这里对接银行支付网关，生成支付链接或二维码
	url := "https://bankpay.com/xxx" // 真实项目需对接银行支付API
	return url, nil
}

func GetPayment(method string) Payment {
	switch method {
	case "alipay":
		return AlipayPayment{}
	case "wechat":
		return WechatPayment{}
	case "bankcard":
		return BankcardPayment{}
	default:
		return nil
	}
}
