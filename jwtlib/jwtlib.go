package jwtlib

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/iskraman/golang-modules/jsonlib"
	"github.com/iskraman/golang-modules/syslog"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Payload interface{} `json:"payload,omitempty"`
	jwt.StandardClaims
}

func New(data interface{}, expireSec int64) (string, error) {
	expirationTime := time.Now().Add(time.Second * time.Duration(expireSec))
	//jsonMsg, _ := jsonlib.Encoding(data)
	claims := &Claims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		syslog.ERRLN(err)
	}

	return tokenString, err
}

func GetData(tokenString string) ([]byte, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if !token.Valid || err != nil {
		syslog.ERRLN(err)
	}

	payload, err := jsonlib.Encoding(claims.Payload)
	return payload, err
}

func RefreshExpireTime(tokenString string, expireSec int64) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if !token.Valid || err != nil {
		syslog.WARLN(err)
	}

	expirationTime := time.Now().Add(time.Second * time.Duration(expireSec))
	claims = &Claims{
		Payload: claims.Payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		syslog.ERRLN(err)
	}

	return tokenString, err
}

func GetExpireTime(tokenString string) (float64, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if !token.Valid || err != nil {
		syslog.WARLN(err)
	}

	expireT := time.Unix(claims.ExpiresAt, 0).Sub(time.Now()).Seconds()
	return expireT, err
}

/*
type UserInfo struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func main() {
	user := UserInfo{Username: "iskraman", Age: 12}
	tokenString, _ := New(user, 10)
	syslog.STDLN("[JWT]", tokenString)

	userdata, _ := GetData(tokenString)
	syslog.STD("%+v", string(userdata))

	userInfo := &UserInfo{}
	jsonlib.Decoding(userdata, userInfo)
	syslog.STD("%+v", userInfo)

	for {
		expireTime, _ := GetExpireTime(tokenString)
		syslog.STD("%.2f", expireTime)
		time.Sleep(time.Second * 1)

		if expireTime < 1 {
			tokenString, _ = RefreshExpireTime(tokenString, 30)
		}
	}
}
*/

/*
func Auth(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}

func Confirm(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		syslog.ERRLN(err)
	}

	tokenStr := cookie.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if !token.Valid || err != nil {
		syslog.ERRLN(err)
	}

	syslog.STD("Hello, %s", claims.Username)
}
*/
