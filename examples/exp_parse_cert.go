package main

import (
	//  "crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"io/ioutil"
)

type CertInfo struct {
	SerialNumber       string    `binding:"required" form:"serial_number" json:"serial_number"`             // 证书序列号
	Version            int      `binding:"required" form:"version" json:"version"`                         // 证书版本（0:V1,1:V2,2:V3）
	IssuerName         string    `binding:"required" form:"issuer_name" json:"issuer_name"`                 // 颁发机构
	SignatureAlgorithm string    `binding:"required" form:"signature_algorithm" json:"signature_algorithm"` // 签名算法
	NotBefore          time.Time `binding:"required" form:"not_before" json:"not_before"`                   // 有效期开始时间
	NotAfter           time.Time `binding:"required" form:"not_after" json:"not_after"`                     // 有效期结束时间
}

func parseCert(crt, privateKey string) *tls.Certificate {
	var cert tls.Certificate
	//加载PEM格式证书到字节数组
	certPEMBlock, err := ioutil.ReadFile(crt)
	if err != nil {
		return nil
	}
	//获取下一个pem格式证书数据 -----BEGIN CERTIFICATE-----   -----END CERTIFICATE-----
	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		return nil
	}
	//附加数字证书到返回
	cert.Certificate = append(cert.Certificate, certDERBlock.Bytes)
	//继续解析Certifacate Chan,这里要明白证书链的概念
	certDERBlockChain, _ := pem.Decode(restPEMBlock)
	if certDERBlockChain != nil {
		//追加证书链证书到返回
		cert.Certificate = append(cert.Certificate, certDERBlockChain.Bytes)
		fmt.Println("存在证书链")
	}

	//读取RSA私钥进文件到字节数组
	keyPEMBlock, err := ioutil.ReadFile(privateKey)
	if err != nil {
		return nil
	}

	//解码pem格式的私钥------BEGIN RSA PRIVATE KEY-----   -----END RSA PRIVATE KEY-----
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		return nil
	}
	//打印出私钥类型
	fmt.Println(keyDERBlock.Type)
	fmt.Println(keyDERBlock.Headers)
	var key interface{}
	var errParsePK error
	if keyDERBlock.Type == "RSA PRIVATE KEY" {
		//RSA PKCS1
		key, errParsePK = x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	} else if keyDERBlock.Type == "PRIVATE KEY" {
		//pkcs8格式的私钥解析
		key, errParsePK = x509.ParsePKCS8PrivateKey(keyDERBlock.Bytes)
	}

	if errParsePK != nil {
		return nil
	} else {
		cert.PrivateKey = key
	}
	//第一个叶子证书就是我们https中使用的证书
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		fmt.Println("x509证书解析失败")
		return nil
	} else {
		switch x509Cert.PublicKeyAlgorithm {
		case x509.RSA:
			{
				fmt.Println("Plublic Key Algorithm:RSA")
			}
		case x509.DSA:
			{
				fmt.Println("Plublic Key Algorithm:DSA")
			}
		case x509.ECDSA:
			{
				fmt.Println("Plublic Key Algorithm:ECDSA")
			}
		case x509.UnknownPublicKeyAlgorithm:
			{
				fmt.Println("Plublic Key Algorithm:Unknow")
			}
		}
	}
	return &cert
}

func parsePemFile(path string) (x509Cert *x509.Certificate, err error) {
	certPEMBlock, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	//获取证书信息 -----BEGIN CERTIFICATE-----   -----END CERTIFICATE-----
	//这里返回的第二个值是证书中剩余的 block, 一般是rsa私钥 也就是 -----BEGIN RSA PRIVATE KEY 部分
	//一般证书的有效期，组织信息等都在第一个部分里
	fmt.Println(certPEMBlock)
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		err = errors.New("cert is empty")
		return
	}
	fmt.Println(certDERBlock)
	x509Cert, err = x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return
	}
	//log.Printf("certFile=%s, validation time %s ~ %s", path,
	//	x509Cert.NotBefore.Format("2006-01-02 15:04:05"), x509Cert.NotAfter.Format("2006-01-02 15:04:05"))
	fmt.Println(x509Cert.PublicKey)
	certInfo := &CertInfo{
		SerialNumber: x509Cert.SerialNumber.String(),
		Version: x509Cert.Version,
		IssuerName: x509Cert.Issuer.CommonName,
		SignatureAlgorithm: x509Cert.SignatureAlgorithm.String(),
		NotBefore: x509Cert.NotBefore,
		NotAfter: x509Cert.NotAfter,
	}
	certInfoJson, err := json.Marshal(certInfo) //转换成JSON返回的是byte[]
	if err != nil {
		return
	}
	var certInfoMap map[string]interface{}
	err = json.Unmarshal(certInfoJson, &certInfoMap)
	if err != nil {
		return
	}
	fmt.Println(string(certInfoJson))
	fmt.Println(certInfoMap)
	fmt.Println(certInfo)
	return
}

func main() {
	//fmt.Println("---------pkcs8 private key ---------------")
	//parseCert("./server.crt", "pkcs8_server.key")
	//fmt.Println("---------pkcs1  private key ---------------")
	//parseCert("./server.crt", "server.key")
	parsePemFile("examples/test_cert.der")
}
