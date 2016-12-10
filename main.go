// EncodeAndDecode project main.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

// 高级加密算法crypto/aes包和crypto/des包
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	// 需要加密的字符串
	plaintext := []byte("My name is Astaxie")
	// 如果传入加密串的话，plaintext就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}
	// aes的加密字符串
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println(len(key_text))

	// 创建加密算法aes，参数key必须是16、24或者32位的[]byte
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes = %s", len(key_text), err.Error())
		os.Exit(1)
	}

	// 加密字符串
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

// base64加密
func base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// base64解密
func base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

// 测试base64加密解密
func testBase64() {
	src := "你好，世界! Hello world!"
	encode := base64Encode(src)
	decode, err := base64Decode(encode)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("原始字符串为:", src)
	fmt.Println("编码后为:", encode)
	fmt.Println("解码后为:", string(decode))
}
