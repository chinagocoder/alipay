// @Title  alipay.go
// @Description  支付宝初始化
// @Author  xushuai  2022/11/23 10:27 AM

package alipay

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/smartwalle/crypto4go"
)

var (
	ErrSignNotFound         = errors.New("alipay: sign content not found")
	ErrAliPublicKeyNotFound = errors.New("alipay: alipay public key not found")
)

const (
	kAliPayPublicKeySN = "alipay-public-key"
	kAppAuthToken      = "app_auth_token"
)

type Client struct {
	mu                 sync.Mutex
	isProduction       bool
	appId              string
	apiDomain          string
	notifyVerifyDomain string
	Client             *http.Client
	location           *time.Location

	appPrivateKey    *rsa.PrivateKey // 应用私钥
	appCertSN        string
	rootCertSN       string
	aliPublicCertSN  string
	aliPublicKeyList map[string]*rsa.PublicKey
}

type OptionFunc func(c *Client)

func WithTimeLocation(location *time.Location) OptionFunc {
	return func(c *Client) {
		c.location = location
	}
}

func WithHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		c.Client = client
	}
}

// New
// @Description  初始化支付宝客户端
// @param  appId 支付宝应用 id
// @param  privateKey 应用私钥，开发者自己生成
// @param  isProduction 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
// @param  opts 可选方法
// @return  client 支付宝客户端实例
// @return  err 错误信息
func New(appId, privateKey string, isProduction bool, opts ...OptionFunc) (client *Client, err error) {
	priKey, err := crypto4go.ParsePKCS1PrivateKey(crypto4go.FormatPKCS1PrivateKey(privateKey))
	if err != nil {
		priKey, err = crypto4go.ParsePKCS8PrivateKey(crypto4go.FormatPKCS8PrivateKey(privateKey))
		if err != nil {
			return nil, err
		}
	}
	client = &Client{}
	client.isProduction = isProduction
	client.appId = appId

	if client.isProduction {
		client.apiDomain = kProductionURL
		client.notifyVerifyDomain = kProductionMAPIURL
	} else {
		client.apiDomain = kSandboxURL
		client.notifyVerifyDomain = kSandboxURL
	}
	client.Client = http.DefaultClient
	client.location = time.Local

	client.appPrivateKey = priKey
	client.aliPublicKeyList = make(map[string]*rsa.PublicKey)

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

// IsProduction
// @Description 是否是生产环境
// @receiver  c 客户端实例
// @return  bool true/false
func (c *Client) IsProduction() bool {
	return c.isProduction
}

// LoadAliPayPublicKey
// @Description  加载支付宝公钥
// @receiver  c 客户端实例
// @param  aliPublicKey 支付宝公钥串
// @return  error
func (c *Client) LoadAliPayPublicKey(aliPublicKey string) error {
	var pub *rsa.PublicKey
	var err error
	if len(aliPublicKey) < 0 {
		return ErrAliPublicKeyNotFound
	}
	pub, err = crypto4go.ParsePublicKey(crypto4go.FormatPublicKey(aliPublicKey))
	if err != nil {
		return err
	}
	c.mu.Lock()
	c.aliPublicCertSN = kAliPayPublicKeySN
	c.aliPublicKeyList[c.aliPublicCertSN] = pub
	c.mu.Unlock()
	return nil
}

// LoadAppPublicCert
// @Description  加载应用公钥证书
// @receiver  c 客户端实例
// @param  s 应用公钥串
// @return  error 错误
func (c *Client) LoadAppPublicCert(s string) error {
	cert, err := crypto4go.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}
	c.appCertSN = getCertSN(cert)
	return nil
}

// LoadAppPublicCertFromFile
// @Description  加载应用公钥证书,通过文件方式
// @receiver  c 客户端实例
// @param  filename 应用公钥文件名
// @return  error 错误
func (c *Client) LoadAppPublicCertFromFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return c.LoadAppPublicCert(string(b))
}

// LoadAliPayPublicCert
// @Description  加载支付宝公钥证书，通过文件方式
// @receiver  c 客户端实例
// @param  filename 应用公钥文件名
// @return  error 错误
func (c *Client) LoadAliPayPublicCert(s string) error {
	cert, err := crypto4go.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return nil
	}

	c.mu.Lock()
	c.aliPublicCertSN = getCertSN(cert)
	c.aliPublicKeyList[c.aliPublicCertSN] = key
	c.mu.Unlock()

	return nil
}

// LoadAliPayPublicCertFromFile
// @Description  加载支付宝公钥证书，通过文件
// @receiver  c 客户端实例
// @param  filename 证书文件名
// @return  error 错误信息
func (c *Client) LoadAliPayPublicCertFromFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return c.LoadAliPayPublicCert(string(b))
}

// LoadAliPayRootCert
// @Description  加载支付宝根证书
// @receiver  c 客户端实例
// @param  s 证书字符串
// @return  error 错误信息
func (c *Client) LoadAliPayRootCert(s string) error {
	var certStrList = strings.Split(s, kCertificateEnd)

	var certSNList = make([]string, 0, len(certStrList))

	for _, certStr := range certStrList {
		certStr = certStr + kCertificateEnd

		var cert, _ = crypto4go.ParseCertificate([]byte(certStr))
		if cert != nil && (cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA1WithRSA) {
			certSNList = append(certSNList, getCertSN(cert))
		}
	}

	c.rootCertSN = strings.Join(certSNList, "_")
	return nil
}

// LoadAliPayRootCertFromFile
// @Description  加载支付宝根证书，通过文件试
// @receiver  c 客户端实例
// @param  filename 证书文件名
// @return  error 错误信息
func (c *Client) LoadAliPayRootCertFromFile(filename string) error {
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	return c.LoadAliPayRootCert(string(b))
}

// URLValues
// @Description  生成请求参数
// @receiver  c 客户端实例
// @param  param 参数
// @return  value 组装后的参数
// @return  err 错误信息
func (c *Client) URLValues(param Param) (value url.Values, err error) {
	var p = url.Values{}
	p.Add("app_id", c.appId)
	p.Add("method", param.APIName())
	p.Add("format", kFormat)
	p.Add("charset", kCharset)
	p.Add("sign_type", kSignTypeRSA2)
	p.Add("timestamp", time.Now().In(c.location).Format(kTimeFormat))
	p.Add("version", kVersion)
	if c.appCertSN != "" {
		p.Add("app_cert_sn", c.appCertSN)
	}
	if c.rootCertSN != "" {
		p.Add("alipay_root_cert_sn", c.rootCertSN)
	}

	bytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	p.Add("biz_content", string(bytes))

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			if key == kAppAuthToken {
				continue
			}
			p.Set(key, value)
		}
	}

	sign, err := signWithPKCS1v15(p, c.appPrivateKey, crypto.SHA256)
	if err != nil {
		return nil, err
	}
	p.Add("sign", sign)
	return p, nil
}

// doRequest
// @Description  发起支付宝请求
// @receiver  c 实例
// @param  method 请求方式 GET/POST/PUT/DELETE
// @param  param 请求参数
// @param  result 响应结果
// @return  err 错误信息
func (c *Client) doRequest(method string, param Param, result interface{}) (err error) {
	var buf io.Reader
	if param != nil {
		p, err := c.URLValues(param)
		if err != nil {
			return err
		}
		buf = strings.NewReader(p.Encode())
	}

	req, err := http.NewRequest(method, c.apiDomain, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", kContentType)

	resp, err := c.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dataStr = string(data)

	var rootNodeName = strings.Replace(param.APIName(), ".", "_", -1) + kResponseSuffix

	var rootIndex = strings.LastIndex(dataStr, rootNodeName)
	var errorIndex = strings.LastIndex(dataStr, kErrorResponse)

	var content string
	var certSN string
	var sign string

	if rootIndex > 0 {
		content, certSN, sign = parseJSONSource(dataStr, rootNodeName, rootIndex)
		if sign == "" {
			var errRsp *ErrorRsp
			if err = json.Unmarshal([]byte(content), &errRsp); err != nil {
				return err
			}

			// alipay.open.app.alipaycert.download(应用支付宝公钥证书下载) 没有返回 sign 字段，所以再判断一次 code
			if errRsp.Code != CodeSuccess {
				if errRsp != nil {
					return errRsp
				}
				return ErrSignNotFound
			}
		}
	} else if errorIndex > 0 {
		content, certSN, sign = parseJSONSource(dataStr, kErrorResponse, errorIndex)
		if sign == "" {
			var errRsp *ErrorRsp
			if err = json.Unmarshal([]byte(content), &errRsp); err != nil {
				return err
			}
			return errRsp
		}
	} else {
		return ErrSignNotFound
	}

	if sign != "" {
		publicKey, err := c.getAliPayPublicKey(certSN)
		if err != nil {
			return err
		}

		if ok, err := verifyData([]byte(content), sign, publicKey); ok == false {
			return err
		}
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) DoRequest(method string, param Param, result interface{}) (err error) {
	return c.doRequest(method, param, result)
}

func (c *Client) VerifySign(data url.Values) (ok bool, err error) {
	var certSN = data.Get(kCertSNNodeName)
	publicKey, err := c.getAliPayPublicKey(certSN)
	if err != nil {
		return false, err
	}

	return verifySign(data, publicKey)
}

func (c *Client) getAliPayPublicKey(certSN string) (key *rsa.PublicKey, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if certSN == "" {
		certSN = c.aliPublicCertSN
	}

	key = c.aliPublicKeyList[certSN]

	if key == nil {
		if c.isProduction {
			cert, err := c.downloadAliPayCert(certSN)
			if err != nil {
				return nil, err
			}

			var ok bool
			key, ok = cert.PublicKey.(*rsa.PublicKey)
			if ok == false {
				return nil, ErrAliPublicKeyNotFound
			}
		} else {
			return nil, ErrAliPublicKeyNotFound
		}
	}
	return key, nil
}

func (c *Client) CertDownload(param CertDownload) (result *CertDownloadRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

func (c *Client) downloadAliPayCert(certSN string) (cert *x509.Certificate, err error) {
	var cp = CertDownload{}
	cp.AliPayCertSN = certSN
	rsp, err := c.CertDownload(cp)
	if err != nil {
		return nil, err
	}
	certBytes, err := base64.StdEncoding.DecodeString(rsp.Content.AliPayCertContent)
	if err != nil {
		return nil, err
	}

	cert, err = crypto4go.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return nil, nil
	}

	c.aliPublicCertSN = getCertSN(cert)
	c.aliPublicKeyList[c.aliPublicCertSN] = key

	return cert, nil
}

func parseJSONSource(rawData string, nodeName string, nodeIndex int) (content, certSN, sign string) {
	var dataStartIndex = nodeIndex + len(nodeName) + 2
	var signIndex = strings.LastIndex(rawData, "\""+kSignNodeName+"\"")
	var certIndex = strings.LastIndex(rawData, "\""+kCertSNNodeName+"\"")
	var dataEndIndex int

	if signIndex > 0 && certIndex > 0 {
		dataEndIndex = int(math.Min(float64(signIndex), float64(certIndex))) - 1
	} else if certIndex > 0 {
		dataEndIndex = certIndex - 1
	} else if signIndex > 0 {
		dataEndIndex = signIndex - 1
	} else {
		dataEndIndex = len(rawData) - 1
	}

	var indexLen = dataEndIndex - dataStartIndex
	if indexLen < 0 {
		return "", "", ""
	}
	content = rawData[dataStartIndex:dataEndIndex]

	if certIndex > 0 {
		var certStartIndex = certIndex + len(kCertSNNodeName) + 4
		certSN = rawData[certStartIndex:]
		var certEndIndex = strings.Index(certSN, "\"")
		certSN = certSN[:certEndIndex]
	}

	if signIndex > 0 {
		var signStartIndex = signIndex + len(kSignNodeName) + 4
		sign = rawData[signStartIndex:]
		var signEndIndex = strings.LastIndex(sign, "\"")
		sign = sign[:signEndIndex]
	}

	return content, certSN, sign
}

func signWithPKCS1v15(param url.Values, privateKey *rsa.PrivateKey, hash crypto.Hash) (s string, err error) {
	if param == nil {
		param = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var src = strings.Join(pList, "&")
	sig, err := crypto4go.RSASignWithKey([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func verifySign(data url.Values, key *rsa.PublicKey) (ok bool, err error) {
	sign := data.Get(kSignNodeName)

	var keys = make([]string, 0, 0)
	for k := range data {
		if k == kSignNodeName || k == kSignTypeNodeName || k == kCertSNNodeName {
			continue
		}
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var buf strings.Builder

	for _, k := range keys {
		vs := data[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	s := buf.String()
	return verifyData([]byte(s), sign, key)
}

func verifyData(data []byte, sign string, key *rsa.PublicKey) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	if err = crypto4go.RSAVerifyWithKey(data, signBytes, key, crypto.SHA256); err != nil {
		return false, err
	}
	return true, nil
}

func getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}
