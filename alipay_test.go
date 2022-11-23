package alipay_test

import (
	"fmt"
	"os"

	alipay "github.com/chinagocoder/alipay"
)

var (
	appID = ""

	privateKey = ""

	aliPublicKey = ""
)

var client *alipay.Client

func init() {
	var err error
	client, err = alipay.New(appID, privateKey, true)

	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		os.Exit(-1)
	}

	// 下面两种方式只能二选一
	var cert = false
	if cert {
		// 使用支付宝证书
		fmt.Println("加载证书", client.LoadAppPublicCertFromFile("appPublicCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayRootCertFromFile("alipayRootCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayPublicCertFromFile("alipayPublicCert.crt"))
	} else {
		// 使用支付宝公钥
		fmt.Println("加载公钥", client.LoadAliPayPublicKey(aliPublicKey))
	}
}
