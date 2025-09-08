package portal

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"

	"student-portal-gcuf/internal/ai"
	"student-portal-gcuf/internal/config"
)

type PortalClient struct {
	HTTP   *http.Client
	Config *config.Config
	GenAI  *ai.GenAI
}

func NewPortalClient(cfg *config.Config, genai *ai.GenAI) *PortalClient {
	jar, _ := cookiejar.New(nil)
	return &PortalClient{
		HTTP: &http.Client{Jar: jar,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Config: cfg,
		GenAI:  genai,
	}
}
