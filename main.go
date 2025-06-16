package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("ms-event.grid-event.ms-active-highlight", func(e *colly.HTMLElement) {
		var teams []string
		e.ForEach(".participant", func(i int, el *colly.HTMLElement) {
			teams = append(teams, strings.TrimSpace(el.Text))
		})
		fmt.Println(teams)
	})

	c.Visit("https://sports.sportingbet.co.za/en/sports/football-4/betting/world-6/fifa-club-world-cup-0:23")
}
