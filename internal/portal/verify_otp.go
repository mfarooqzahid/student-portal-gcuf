package portal

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *PortalClient) VerifyOtp(otp string) (bool, error) {

	url := fmt.Sprintf("%s%s?TOKEN=%s",
		c.Config.Portal.BaseURL,
		c.Config.Portal.VerifyOtp,
		otp,
	)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to verify otp: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("body: %v\n", string(body))

	if strings.Contains(string(body), "success index") {
		return true, nil
	}

	if strings.Contains(string(body), "success otp") {
		return false, fmt.Errorf("otp is invalid or expired")
	}

	return false, fmt.Errorf("an error occurred, please try again")
}