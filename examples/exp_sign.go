package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)


//读取私钥和公钥
func getRsaPriKey(priKey string) *rsa.PrivateKey {
	der, _ := base64.StdEncoding.DecodeString(priKey)

	//block, _ := pem.Decode([]byte(priKey))
	//der := block.Bytes

	if rsaPriKey, err := x509.ParsePKCS8PrivateKey(der); err != nil {
		return nil // 获取失败
	} else {
		return rsaPriKey.(*rsa.PrivateKey) // 读取成功
	}
}

func getRsaPubKey(pubKey string) *rsa.PublicKey {
	der, _ := base64.StdEncoding.DecodeString(pubKey)
	//block, _ := pem.Decode([]byte(pubKey))
	//der := block.Bytes

	if rsaPubKey, err := x509.ParsePKIXPublicKey(der); err != nil {
		return nil // 获取失败
	} else {
		return rsaPubKey.(*rsa.PublicKey) // 读取成功
	}
}


// 签名
func Sign(params map[string]interface{}, priKey *rsa.PrivateKey) string {
	// 对请求参数按照字母顺序进行排序并组合
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var signature string
	var signatureItems []string

	for _, k := range keys {
		signatureItems = append(signatureItems, fmt.Sprintf("%s=%v", k, params[k]))
	}
	signature = strings.Join(signatureItems, "&")

	//var signature string
	//i := 0
	//for _, k := range keys {
	//	v := params[k]
	//	if i != 0 {
	//		signature += "&"
	//	}
	//	signature += k
	//	signature += "="
	//	signature += v
	//	i++
	//}
	// 将排序后的signature进行hash操作
	h := sha256.New()
	h.Write([]byte(signature))
	Sha256Code := h.Sum(nil)
	// 使用rsa算法进行签名
	// 第一个参数是一个随机数参数器，确保每次相同输入产生不同的签名
	// 第二个参数是密钥
	// 第三个参数是我们上面使用的hash函数
	// 第四个参数是被hash函数处理过的原始输入
	signatureAfter, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, Sha256Code)
	if err != nil {
		return ""
	}
	// 返回base64编码的字符串
	return base64.StdEncoding.EncodeToString(signatureAfter)
}

//验签
func Verify(params map[string]interface{}, pubKey *rsa.PublicKey, sign string) (err error) {
	// 和签名步骤相同，对收到的请求参数按照字母顺序进行排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var signature string
	var signatureItems []string

	for _, k := range keys {
		signatureItems = append(signatureItems, fmt.Sprintf("%s=%v", k, params[k]))
	}
	signature = strings.Join(signatureItems, "&")

	//i := 0
	//for _, k := range keys {
	//	v := params[k]
	//	if i != 0 {
	//		signature += "&"
	//	}
	//	signature += k
	//	signature += "="
	//	signature += v
	//	i++
	//}
	// 和签名步骤相同，将排序后的signature进行hash操作
	h := sha256.New()
	h.Write([]byte(signature))
	Sha256Code := h.Sum(nil)
	// 对签名进行base64解码
	decodeSignature, err := base64.StdEncoding.DecodeString(sign)
	// 使用rsa验签函数
	// 第一个参数是公钥
	// 第二个参数是hash函数
	// 第三个参数是被hash函数处理过的原始输入
	// 第四个参数是被处理过的签名
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, Sha256Code, decodeSignature)
	if err != nil {  // 验证失败
		return err
	}
	return nil // 验证成功
}


func test_01() {
	//rsa签名
	//生成公钥私钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	fmt.Println(privateKey)

	// The public key is a part of the *rsa.PrivateKey struct
	publicKey := privateKey.PublicKey
	fmt.Println(publicKey)
}

func test_02()  {
	priKey := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAOjvQISAGctEjUVXLBDdd9No+DquAAnrzENpKWWKD9g5ukSVlH1q7GpF7j7F+YzlKtI2GhKNuxdanfBx+3W4XqnoasELs2IcLeJS8f2VGwJ8tradpEfvotvt2GajosuC4qLQj3hUEDc0VdLQvwm/SzCPbxI/uipDVRy87pbUK7mBAgMBAAECgYBbbe2pFI8LVjENUpYdWJC8DqvAfKPPQRrAKvrwvCxT9vTCDuRmBh4d6K0HPrYhM6KHOG0rcy5IkG0F//dv0Zp13uvv5WqxnALB4layMwuGJI1DoMZ5CMdKiXDNla3ujrwMMZS0rJolDUjbuEYJUBGzEuOx+oXcDrjO1hF2u9oE8QJBAPWqQHLT4nFSh/I4YLzFhaaWOG1IA4p1QW0ZD8uQ5/6UsRdBMHuKeZdyvCM3gemqUnjn2Bxxs4Zs2C0Db9lCxlcCQQDyu+V+saaIOJAG6g17aCU4Vk95i4AS2ohiJGtcHXhjmMrbF9d/NPSXnQ7I5jdSysDS6HeoXC64j0iDp6vNuKfnAkBOEePXWgSdVMgFLut3MiJOFXErMXaGmsrG++7sF0nSurgolrvsls4bVgf5/RrZAhqhd7RTHcMalh6A95Wk6raTAkA3sY8Fk8NPDBNX0soHH60ykPMpQy2+pa+XPR35A7YEvCEHRiAmysK4KswiolfdQjId8kLKOWRq6k1MD+OmjxDPAkEAvKmFG9W8QbanLfN1qRKEUPkGF700XxmKeSghw8/DM8olTZJ0HgNVWC2V15bCHCLGtSVJgDLd2InMKksHFRM3Ug=="
	rsaPriKey := getRsaPriKey(priKey)

	pubKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDo70CEgBnLRI1FVywQ3XfTaPg6rgAJ68xDaSllig/YObpElZR9auxqRe4+xfmM5SrSNhoSjbsXWp3wcft1uF6p6GrBC7NiHC3iUvH9lRsCfLa2naRH76Lb7dhmo6LLguKi0I94VBA3NFXS0L8Jv0swj28SP7oqQ1UcvO6W1Cu5gQIDAQAB"
	rsaPubKey := getRsaPubKey(pubKey)


	paramStr := "nonceStr=a4IWeSdZUq4veICs&reqStatus=50&source=09&timestamp=1664394611249057&unionTok=2VPwbhp/jcLp1HVHr+T8Lq6hNNZTL8ndLoHOqzfzuLWzUwwwWrwot43WWmYr0RfnSDJkBVyMcFSpvpTN880OMPQbQv3GVOFgNJgeCYU8TT5X4jMqXdIWj+ux+eAHmNsy"
	//paramStr := "innerSys=cx0c1kl621asfas25r2134j6sdgdsfdc&nonceStr=a4IWeSdZUq4veICs&secret=MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDo70CEgBnLRI1FVywQ3XfTaPg6rgAJ68xDaSllig/YObpElZR9auxqRe4+xfmM5SrSNhoSjbsXWp3wcft1uF6p6GrBC7NiHC3iUvH9lRsCfLa2naRH76Lb7dhmo6LLguKi0I94VBA3NFXS0L8Jv0swj28SP7oqQ1UcvO6W1Cu5gQIDAQAB&timestamp=1664397998083"

	h := sha256.New()
	h.Write([]byte(paramStr))
	Sha256Code := h.Sum(nil)
	// 使用rsa算法进行签名
	// 第一个参数是一个随机数参数器，确保每次相同输入产生不同的签名
	// 第二个参数是密钥
	// 第三个参数是我们上面使用的hash函数
	// 第四个参数是被hash函数处理过的原始输入
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPriKey, crypto.SHA256, Sha256Code)
	if err != nil {
		return
	}
	// 返回base64编码的字符串
	sign := base64.StdEncoding.EncodeToString(signature)
	fmt.Println(sign)


	// 原始值，不进行转义
	params := map[string]interface{}{
		"unionTok": "2VPwbhp/jcLp1HVHr+T8Lq6hNNZTL8ndLoHOqzfzuLWzUwwwWrwot43WWmYr0RfnSDJkBVyMcFSpvpTN880OMPQbQv3GVOFgNJgeCYU8TT5X4jMqXdIWj+ux+eAHmNsy",
		"source": "09",
		"reqStatus": "50",
		"timestamp": 1664394611249057,
		"nonceStr": "a4IWeSdZUq4veICs",
	}

	fmt.Println(Sign(params, rsaPriKey))

	err = Verify(params, rsaPubKey, sign)
	if err != nil {
		fmt.Println("验签失败")
		return
	}
	fmt.Println("验签成功")
}

func test_03()  {
	priKey := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMKzDKJQ8o1KVVI/7QKtLPb6Ct2GxKvtan5Cgn0pytdEa0btwv6ehfJ2y1SqCv4CsorGj25SxtaSxv4sISbhJKU0QUJ5iOqHwdaj7W9yF0Dsu9ayZ3HIESCWr5BtipCx7EQK1mn3QdLhQA9Di3znuTWQDSKwQGaR5Yfj9tCUm435AgMBAAECgYAPBUE0RIpx3Ao5DjV6v2F4rFGqRXtmyCovE3J9tXuwn9am4jDo89HhosiM/5FZh1u+2RK3sCsm4qv3b/Aez7/D3G+yyuvzIoJz4lGBFCXjfNOxM6tx1yWLEryu81jtUK+d08Ds7jm25f1nckdIo7IQPrV2tQ0TJq4PxjkvaRicAQJBAPOzuuzSw4Wf4zxGWqCEatA6KPF9kLfw91I8BCyQZ4W15EFNzMK7YSl9nEiMknQjVcl9jJayz1CVZL6KersFFnUCQQDMhkcDDb6Ue40WXf94G1x0cmLZYx6AW/+IG/f/xKivm7ccOs6qd5o9Nm1TXSXYTzGPLugzSR/59eLt4YS4xtD1AkEA2Sqje2DB7XSyiUbkgcAecxow3CoZ6WBo9LYYviGcrtGM6wjKXAmgptmRiDEEXUwTqW8g9jW3esb2NLmwbrI09QJAJoXEKlcUZnM/B898purAZ06EP93jiNnLL5/U/l/URf9WZgLNml4ctxrhTKpmzm9cLlPveHl1hGdD7vpjPUp1YQJADy+FQOeFyzhuvkECdhiVFNsWMmlGV7YG+5/CvBCGasXw+CWahYAxAcTasFex7TWBx1ZRpeOQmq4YfcGlL2n23Q=="
	rsaPriKey := getRsaPriKey(priKey)

	//pubKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDo70CEgBnLRI1FVywQ3XfTaPg6rgAJ68xDaSllig/YObpElZR9auxqRe4+xfmM5SrSNhoSjbsXWp3wcft1uF6p6GrBC7NiHC3iUvH9lRsCfLa2naRH76Lb7dhmo6LLguKi0I94VBA3NFXS0L8Jv0swj28SP7oqQ1UcvO6W1Cu5gQIDAQAB"
	//rsaPubKey := getRsaPubKey(pubKey)


	//paramStr := "nonceStr=a4IWeSdZUq4veICs&reqStatus=50&source=09&timestamp=1664394611249057&unionTok=2VPwbhp/jcLp1HVHr+T8Lq6hNNZTL8ndLoHOqzfzuLWzUwwwWrwot43WWmYr0RfnSDJkBVyMcFSpvpTN880OMPQbQv3GVOFgNJgeCYU8TT5X4jMqXdIWj+ux+eAHmNsy"
	paramStr := "innerSys=175&nonceStr=a4IWeSdZUq4veICs&secret=iSPCrZxmuzgUIzpS&timestamp=1664419345677"

	h := sha256.New()
	h.Write([]byte(paramStr))
	Sha256Code := h.Sum(nil)
	// 使用rsa算法进行签名
	// 第一个参数是一个随机数参数器，确保每次相同输入产生不同的签名
	// 第二个参数是密钥
	// 第三个参数是我们上面使用的hash函数
	// 第四个参数是被hash函数处理过的原始输入
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPriKey, crypto.SHA256, Sha256Code)
	if err != nil {
		return
	}
	// 返回base64编码的字符串
	sign := base64.StdEncoding.EncodeToString(signature)
	fmt.Println(sign)
}

func main()  {
	//test_01()
	test_02()
	//test_03()
}
