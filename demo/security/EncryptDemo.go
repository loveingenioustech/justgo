package main

import (
	"crypto/sha256"
	"io"
	"crypto/sha1"
	"crypto/md5"
	"fmt"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"os"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	demoOneWayHash()

	demoSalt()

	demoBase64()

	demoAes()
}

// 单向哈希算法有一个特征：无法通过哈希后的摘要(digest)恢复原始数据
// 常用的单向哈希算法包括SHA-256, SHA-1, MD5等
func demoOneWayHash(){
	//单向哈希有两个特性：
//	1）同一个密码进行单向哈希，得到的总是唯一确定的摘要。
//	2）计算速度快。随着技术进步，一秒钟能够完成数十亿次单向哈希计算。

	//import "crypto/sha256"
	h1 := sha256.New()
	io.WriteString(h1, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h1.Sum(nil))
	fmt.Println()

	//import "crypto/sha1"
	h2 := sha1.New()
	io.WriteString(h2, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h2.Sum(nil))
	fmt.Println()

	//import "crypto/md5"
	h3 := md5.New()
	io.WriteString(h3, "需要加密的密码")
	fmt.Printf("%x", h3.Sum(nil))
	fmt.Println()
}

func demoSalt(){
	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "需要加密的密码")

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"
	// 在两个salt没有泄露的情况下，黑客如果拿到的是最后这个加密串，就几乎不可能推算出原始的密码是什么了。

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last :=fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(last)
}

func demoBase64(){
	// 如果Web应用足够简单，数据的安全性没有那么严格的要求，那么可以采用一种比较简单的加解密方法是base64
	// encode
	hello := "你好，世界！ hello world"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func demoAes(){
	//需要去加密的字符串
	plaintext := []byte("My name is Robin")
	//如果传入加密串的话，plaint就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	//aes的加密字符串
	key_text := "lalalal12798akljzmknm.ahkjkljl;k" //参数key必须是16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println(len(key_text))

	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}
