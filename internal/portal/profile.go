package portal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"student-portal-gcuf/internal/model"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/genai"
)

func (c *PortalClient) GetProfile() (*model.Profile, error) {

	var profile model.Profile

	profileUrl := c.Config.Portal.BaseURL + c.Config.Portal.Profile

	req, _ := http.NewRequest("GET", profileUrl, nil)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile: %v", err)
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	profiledata := doc.Find("#content > div:nth-child(4) > div > div > div").Text()
	profiledata = strings.Join(strings.Fields(profiledata), " ")

	instruction := fmt.Sprintf("Return profile data into given json structure for the image url follow (upload/{cnic}.png): %s", string(profiledata))
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"name":       {Type: genai.TypeString},
				"fathername": {Type: genai.TypeString},
				"rollno":     {Type: genai.TypeString},
				"email":      {Type: genai.TypeString},
				"gender":     {Type: genai.TypeString},
				"dob":        {Type: genai.TypeString},
				"phone":      {Type: genai.TypeString},
				"city":       {Type: genai.TypeString},
				"address":    {Type: genai.TypeString},
				"imageurl":   {Type: genai.TypeString},
				"program":    {Type: genai.TypeString},
				"session":    {Type: genai.TypeString},
				"regno":      {Type: genai.TypeString},
			},
		},
	}
	res, err := c.GenAI.Generate(context.Background(), config, instruction)
	if err != nil {
		return nil, fmt.Errorf("failed to generate profile: %v", err)
	}

	json.Unmarshal([]byte(res), &profile)

	return &profile, nil
}
