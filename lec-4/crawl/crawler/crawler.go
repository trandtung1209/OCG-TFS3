package crawler

import (
	"fmt"

	"github.com/trandtung1209/crawl/database"

	"github.com/gocolly/colly/v2"
	"github.com/trandtung1209/crawl/entity"
)

func Crawl()  {
	db:= database.ConnectDB()

	c:=colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting ...", r.URL)
	})

	c.OnHTML("tr", func(e *colly.HTMLElement) {
		film:= entity.Film{
			Title: e.ChildText(".titleColumn > a"),
			Year:e.ChildText(".titleColumn .secondaryInfo"),
			Rating:e.ChildText(".ratingColumn > strong"),
		}
		db.Create(&film)
	})

	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
}