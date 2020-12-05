package main

import "fmt"
import "crypto/hmac"
import "crypto/sha256"

func main() {

    passwd  := "abc123";
    use     := "huang"
    hash := hmac.New(sha256.New,[]byte(passwd)) // 创建对应的sha256哈希加密算法
    hash.Write([]byte(use+passwd)) // 写入加密数据
    fmt.Printf("%x\n",hash.Sum(nil)) // c10a04b78bcbcc1c4cba37f6afe0fa60cbf08f6e0a1d93b09387f7069be1aeff
    hashString := hash.Sum(nil)

    passwd_2 := "abc";
    if CheckMAC([]byte(use+passwd),hashString,[]byte(passwd)) {
       fmt.Println("CheckMAC Pass!");
    } else {
       fmt.Println("CheckMAC Fail!");
    }

    if CheckMAC([]byte(use+passwd),hashString,[]byte(passwd_2)) {
       fmt.Println("CheckMAC2 Pass!");
    } else {
       fmt.Println("CheckMAC2 Fail!");
    }

}

// CheckMAC 报告 messageMAC 是否是消息的有效 HMAC 标记。
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
