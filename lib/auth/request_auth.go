package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"sort"
)

func MakeSignature(request string, token string) string {
	payload := request + token
	//debug
	//fmt.Println(payload)
	s := []byte(payload)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

	hmac256 := hmac.New(sha256.New, []byte(token))
	hmac256.Write(s)
	sig := base64.StdEncoding.EncodeToString(hmac256.Sum(nil))
	return sig
}
