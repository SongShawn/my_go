package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"mygo/util"
	"os"
	"regexp"

	"github.com/daviddengcn/go-colortext"
	"github.com/daviddengcn/go-colortext/fmt"
)

func main() {
	//hexEncode()
	//useRegexpPkt()
	//useMap()
	//testInherit()
	//useTypeAssert()
	//useSlice()
	//useMap()
	//useRSACrypt()
	//printColorFul()
	utf8convert()
}

func utf8convert() {
	//a := "仿真"
	//
	//for i, j := range []byte(a) {
	//	fmt.Printf("i=%d c=%x\n", i, j)
	//}
	//
	//for _, j := range []byte(a) {
	//	fmt.Printf("%x", j)
	//}
	//fmt.Println()
	//
	//fmt.Println(utf8.RuneCountInString(a))
	//fmt.Println(len(a))
	//
	//buf := []byte{228, 184, 150} // 世
	//fmt.Println(utf8.FullRune(buf))
	//fmt.Println(utf8.FullRune(buf[:2]))

	ascii2hex := map[uint8]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'a': 0xa, 'b': 0xb, 'c': 0xc, 'd': 0xd, 'e': 0xe, 'f': 0xf,
	}

	a := "e4bbbfe79c9f"
	b, _ := hex.DecodeString(a)
	fmt.Println(string(b))
	fmt.Printf("%x\n", ascii2hex[a[0]]<<4|ascii2hex[a[1]])
}

func printColorFul() {
	util.Error("sdfafsaff")
	util.Info("dafadfafa")
	util.Trace("adfdafadfaasfa")

	fmt.Printf("ppid : %d\n", os.Getppid())

	ctfmt.Printf(ct.Red, true, "%s\n", "sdfadfasf")
}

func getRSASize(key *rsa.PublicKey) int {
	return (key.N.BitLen() + 7) / 8
}

func RsaEncrypt(pubKey *rsa.PublicKey, originData []byte) (cipherData []byte, err error) {
	return nil, nil
}

func useRSACrypt() {
	publicKeyFile := "c:/users/duoyi/rsa_sxn_public_key.pem"
	privateKeyFile := "c:/users/duoyi/rsa_sxn_private_key.pem"

	originData := "宋潇宁"

	pubKeyBuf, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		panic("read public key file failed.")
	}

	pubKeyBlock, _ := pem.Decode(pubKeyBuf)
	pubIntf, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if err != nil {
		panic("x509 parsePKIXPublickKey failed.")
	}

	priKeyBuf, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		panic("read private key file failed.")
	}
	priKeyBlock, _ := pem.Decode(priKeyBuf)
	priKey, err := x509.ParsePKCS1PrivateKey(priKeyBlock.Bytes)
	if err != nil {
		panic("x509 parsePKIXPub")
	}

	pubKey := pubIntf.(*rsa.PublicKey)
	//priKey := priIntf.(*rsa.PrivateKey)

	afterEnc, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(originData))
	base64.StdEncoding.EncodeToString(afterEnc)
	fmt.Printf("msg after encrypted : %s\n", base64.StdEncoding.EncodeToString(afterEnc))

	afterDec, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, afterEnc)
	fmt.Printf("msg after decrypted : %s\n", string(afterDec))

	fmt.Printf("bitlen:%d\n", pubKey.N.BitLen())
}

//func testPkgDependCycle() {
//	xx := apkg.APKG11{}
//	fmt.Println(xx)
//}

func useStringByte() {

}

func useSlice() {
	var a []int

	a = append(a, 10)
	fmt.Println(a)

	xx := &AAA{}
	xx.aaa = append(xx.aaa, 10)
	fmt.Println(xx)
	//
	//b := append(b, 100)
	//fmt.Println(b)
}

func useTypeAssert() {
	a := &AAA{}

	var b interface{}
	b = a

	c, ok := b.(*AAA)
	fmt.Println(c, ok)

	d, ok := b.(*BBB)
	fmt.Println(d, ok)

}

// 继承

type AAA struct {
	aaa []int
}

type AAAer interface {
	func1111()
}

func (x *AAA) func1111() {
	fmt.Println("AAA")
}

type BBB struct {
	AAA
	a int
}

type CCC struct {
	AAA
}

func (x *BBB) func1111() {
	fmt.Println("BBB")
}

func Inherittest(in AAAer) {
	in2 := in.(*AAA)
	in2.func1111()
}

func testInherit() {
	var a AAAer
	a = &BBB{}
	a.func1111()

	var b AAAer
	b = &CCC{}
	b.func1111()

	//Inherittest(a)
}

func useMap() {

	fmt.Println("case 1")
	a := make(map[int]*int)
	var b = 10
	a[10] = &b

	c := a[10]
	if c == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("not nil")
	}

	c = a[20]
	if c == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("not nil")
	}

	fmt.Println(len(a))

	fmt.Println("case 2")
	aa := make(map[string]int)
	aa["a"] = 10
	aa["b"] = 11
	aa["c"] = 12

	for key := range aa {
		fmt.Println(key)
	}
}

func hexEncode() {
	name := []byte("指挥官")
	fmt.Println(hex.EncodeToString(name))
}

func useRegexpPkt() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`)

	content1 := []byte(`<a href="https://sobooks.cc/go.html?url=https://pan.baidu.com/s/1CnEep8caA7MlklKhTeUnCg" target="_blank" rel="nofollow" >百度网盘</a>`)

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	pattern2 := regexp.MustCompile(
		`<a href="https://sobooks.cc/go.html\?url=(?P<key>\w+)".*百度网盘<\/a>`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := []byte("$key=$value\n")

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		fmt.Println(string(content[submatches[0]:submatches[1]]))
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))

	template1 := []byte("$key ")
	result1 := []byte{}

	result1 = pattern2.Expand(result1, template1, content1, []int{0, len(content1)})
	fmt.Println(result1)
}
