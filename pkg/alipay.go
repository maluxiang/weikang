package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

type NewAlipay struct {
	Subject     string
	OutTradeNo  string
	TotalAmount float64
}

func (a NewAlipay) Alipay() string {
	var appId = "9021000143642533"
	var privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCPvf8oB3gKgaNPRdSGR+VQ7XGzxuoBTM5+6DJbp7PLCQsmulJIz5UBkV4QWXLhIwE1OnOInyJXZoaPSiDDHMtUgGiVqs3KPKRxvu7JoxjFpIVCajFUf31XjqhfDRagI5brALHNzpQu6YqTZBsHORWp/0vDDeDFW6F3iAYmlL7Tmi5SuPYsI5nmIqhl8he0A2TYMsaLMUAm8Bl0eDPet3VGS1O11CVlAGvZBeT2eLRGkZIDOgHkUPTtxog0//7in2IvxfPP9hRzhKSsErrnww81AVf0YSZh7+nusnAO8WSBeeHSOboKem5ikXop27ZkEVbWvxZqOdpXF+ll56R9wrPDAgMBAAECggEAStyVCu/44N/rgdnIebbwA1nmoc6zCMJLuWSM9Zc92Dx5hk37lNgnro1bBADSB9cHRgWPLlRSulo3MmFU9skETJoj3e77BpzWuDbMfk9nE2g/zvsnfe54BSXTIs3aPsTKUNInyBwSqOwQs/qq6PQeCE9h0dBkg+TZDzvUmDtXtzQJZyS6bjXtDv/APylgi+q5x0XDXYjZzSmLIfa8usSa7omcU39ma7CrtT+AROKa3HFIyxUaSaKf/B8l6/0tyt4fHCEM0qkP9M3pNEwmSb0bZkN3yQHF+ufCvkUBcTt/mE8Qle0yFc7ks5xUJ6ePAwaIAceVazn7U7GMSvY8HRl3YQKBgQDJanT8Gg5Wg7uQvBwWuNxmNcAorl0a7jxbfH+O7DE8bpQUAYEbkf0O2RyrFoQRVb7vg23hjhJa81oN4wWP9D8UQ6Sl0xk96XrJiFErSARvFrI8pXPgx8pWtgInwq84yHGE9XH8kmQf9DfingzIMHO/W1E77rFbJ/m20vO3cKLfNQKBgQC2slZiwK/su+uNot0OaXEswumU3hZNjtckKATlBMFIFvRHncuooJaGBuCTCRDPOsZLhedEYLVGImRNWt2HYbs5HVxeq6pTkfHeuSbbCGSsZMKgNb6B9Zt3opg5QCGF977DXnZBwHFPpG8xz/FxG3uq66qkTzIqDuPgGNRAbfTOFwKBgDzHIwXv1uPEXJUQLmms3tT440NjWjUGLrBsoRE/tJvHwmdHDO6E3xfb7Aq0gKW6eiNMRZwKgv9u733BQ6xsx5wVzVk8miFUkvi9acDlunDLKH2kb5MktqnzwjK9TKKV2auFZSASDKSXzUVU7AZ8mHDl3V14aYxYQ8InZeO02XEBAoGANWR/LJMllukAmT8cnYahRbSc8R3KpNX2+CEd9RRjrD7RG1D5YQm1k+vUnAQNPpLtusqiYPBmad3JNHY2wccFIVb8VMqUl6HSbwXrXh3g1iUIYCv0xiRSUC9bj3e3lGBoBs7HfsXQF7d1q7ga1rRyeuwuzaA7h+EcJbT76ux3m/cCgYEAwYlAvyWXPS3bs8u3YcZBIs0JnMuixR4E37HcGhh1GpthWWWEjK6MEGYjZCYmJCDdhpauRsHeOO7bF9WB0VTGp4kQSWuI6MGVJu4lTJyVKJoYq7g+5ywlHFScRiGwlQUuNGTDPjHG2b1zjggKBDYbt4YhbGV8/3p7GxQsayEQcks=" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	client, err := alipay.New(appId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "https://www.baidu.com"
	p.Subject = a.Subject
	p.OutTradeNo = a.OutTradeNo
	p.TotalAmount = a.TotalAmount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
