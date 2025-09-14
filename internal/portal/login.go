package portal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *PortalClient) Login(username string, password string) (bool, error) {

	loginUrl := c.Config.Portal.BaseURL + c.Config.Portal.Login

	params := url.Values{}
	params.Set("userName", username)
	params.Set("password", password)

	req, _ := http.NewRequest("POST", loginUrl, strings.NewReader(params.Encode()))
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to login: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("body: %v\n", string(body))

	if strings.Contains(string(body), "login-otp.php") {
		var otp string
		// VERIFY OTP
		fmt.Println("enter otp")
		fmt.Scanln(&otp)
		ok, err := c.VerifyOtp(otp)
		if err != nil || !ok {
			log.Fatalf("otp failed:%v", err)
			return false, err
		}
		log.Println("otp verified")
		return true, nil
	}
	if strings.Contains(string(body), "index.php") {
		return true, nil
	}

	if strings.Contains(string(body), "Username and Password wrong") {
		return false, fmt.Errorf("invalid credentials")
	}

	return false, fmt.Errorf("unexpected login response")
}
