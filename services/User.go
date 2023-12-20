package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"ztgo/secure"
	"ztgo/utils"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	Username = "superman"
)

func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
func EnsureSecretFileExist() {
	_, err := os.Stat("./secret")
	if err != nil {
		os.Mkdir("./secret", 0755)
	}
	if !fileIsExist("./secret/secret") {
		fout, err := os.Create("./secret/secret")
		if err != nil {
			fmt.Println("./secret/secret", err)
			return
		}
		defer fout.Close()

		secretStr := generateSecret()
		secretMap := make(map[string]string)
		secretMap["secret"] = secretStr.Secret()
		secretMap["url"] = secretStr.URL()
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(secretMap)
		_, err = fout.WriteString(bf.String())
		if err != nil {
			return
		}
	}
}
func ReadSecretFile() string {
	content, err := os.ReadFile("./secret/secret")
	if err != nil {
		fmt.Println("./secret/secret", err)
		return ""
	}
	secretMap := make(map[string]string)
	err = json.Unmarshal(content, &secretMap)
	if err != nil {
		return ""
	}
	return secretMap["secret"]
}

func generateSecret() *otp.Key {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "superman",
		AccountName: "superman",
	})
	if err != nil {
		panic(err)
	}
	return key
}

func Validate(code string) bool {
	return totp.Validate(code, ReadSecretFile())
}

func ValidateTOTPCode(ctx *gin.Context) {
	postData := utils.GetPostData[string, string](ctx)
	if _, ok := postData["passcode"]; !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	check := Validate(postData["passcode"])
	if !check {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2003,
			"msg":  "校验码错误",
		})
		return
	}

	token, _ := secure.MakeToken(Username)
	// ctx.SetSameSite(http.SameSiteNoneMode)
	// ctx.SetCookie("token", token, 60*60*3, "/", ctx.GetHeader("host"), false, false)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  2000,
		"msg":   "success",
		"token": token,
	})
}
