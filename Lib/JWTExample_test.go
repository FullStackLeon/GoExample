package Lib

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/yaml.v2"
)

// 定义JWT签名的密钥
var config Config

// User 定义用户信息结构体
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Config struct {
	JWT struct {
		SecretKey string `yaml:"secretKey"`
	} `yaml:"jwt"`
}

func parseConfig() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Read config file err: ", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Unmarshal config file err: ", err)
	}
}

// GenerateToken 生成JWT token
func GenerateToken(user *User) (string, error) {
	// 创建一个新的JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为1天
	})

	// 使用签名密钥签名token
	tokenString, err := token.SignedString([]byte(config.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken 验证JWT token
func VerifyToken(tokenString string) (*jwt.Token, error) {
	// 解析JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回签名密钥
		return []byte(config.JWT.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TestJWT(t *testing.T) {
	// 创建一个新的用户信息
	user := &User{
		Username: "username",
		Email:    "username@gmail.com",
	}

	parseConfig()

	// 生成JWT token
	tokenString, err := GenerateToken(user)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("JWT token:", tokenString)

	// 验证JWT token
	fmt.Println("JWT token verify and parse")
	token, err := VerifyToken(tokenString)
	if err != nil {
		fmt.Println("Verify JWT token err", err)
		return
	}

	// 获取token中的claims
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("Username:", claims["username"])
	fmt.Println("Email:", claims["email"])
	fmt.Println("Expiration:", time.Unix(int64(claims["exp"].(float64)), 0))
}
