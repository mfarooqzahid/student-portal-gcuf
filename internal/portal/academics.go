package portal

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (c *PortalClient) GetAcademics() {
	url := "https://student.gcuf.edu.pk/ajax/academics-details-model-2020.php"

	req, _ := http.NewRequest("POST", url, nil)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		log.Fatalf("Error sending academics request %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Error reading academics response %v", err)
	}

	doc.Find("a.semester-tab-button").Each(func(i int, s *goquery.Selection) {
		log.Println(i, "SEMESTER RESULT")
		id, _ := s.Attr("data-sem")
		c.AcademicsDetails(id)
	})

}
// func (c *PortalClient) GetAcademics() {
// 	url := c.Config.Portal.BaseURL + c.Config.Portal.Academics

// 	req, _ := http.NewRequest("GET", url, nil)

// 	resp, err := c.HTTP.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error sending academics request %v", err)
// 	}
// 	defer resp.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading academics response %v", err)
// 	}

// 	doc.Find("a.semester-tab-button").Each(func(i int, s *goquery.Selection) {
// 		log.Println(i, "SEMESTER RESULT")
// 		id, _ := s.Attr("data-sem")
// 		c.AcademicsDetails(id)
// 	})

// }

// func (c *PortalClient) AcademicsDetails(id string) {
// 	url := c.Config.Portal.BaseURL + c.Config.Portal.AcademicsDetails

// 	req, _ := http.NewRequest("POST", url, strings.NewReader("SEMESTER_DTL="+id))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	resp, err := c.HTTP.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error sending academics details request %v", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading academics details response %v", err)
// 	}

// 	academicsData := strings.Join(strings.Fields(body.Text()), " ")

// 	fmt.Printf("academicsData: %v\n", academicsData)

// }
