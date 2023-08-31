package cmd

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/gocolly/colly"
)

func mensa_scrap(debug bool) string {

	type Food struct {
		date string
		meal string
	}

	c := colly.NewCollector(
	//colly.AllowedDomains("studentenwerk.uni-heidelberg.de"),
	)

	page := 0
	line_date := 0
	line_meal := 5

	date := []string{}
	meal := []string{}

	c.OnHTML(".mensa-carousel-wrapper-2", func(e *colly.HTMLElement) {
		if page == 2 {
			food := Food{}
			text := e.Text
			text = strings.TrimSpace(text)
			text = strings.Replace(text, "\t", "", -1)
			//text = strings.Replace(text, "", "", -1)
			scanner := bufio.NewScanner(strings.NewReader(text))
			line := 0
			for scanner.Scan() {
				if debug {
					fmt.Println(line, ":\t", scanner.Text())
					line++
				}
				if line_date%50 == 0 {
					if strings.Replace(scanner.Text(), "\n", "", -1) == "" {
						line_date--
					} else {
						food.date = scanner.Text()
						date = append(date, food.date)
					}
				}
				if line_meal%50 == 0 {
					if strings.Replace(scanner.Text(), "\n", "", -1) == "" {
						line_meal--
					} else {
						food.meal = scanner.Text()
						meal = append(meal, food.meal)
					}

				}

				line_date++
				line_meal++
			}
		}
		page++
	})

	c.Visit("https://www.studentenwerk.uni-heidelberg.de/de/speiseplan_neu")

	returnstring := ""

	for i := 0; i < 3; i++ {
		returnstring += date[i]
		returnstring += "\n"
		returnstring += meal[i]
		returnstring += "\n\n"

	}

	return returnstring

}

func ozon_scrap(debug bool) string {

	c := colly.NewCollector(
		colly.AllowedDomains(""),
	)

	if debug {

		c.OnHTML("title", func(e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})

		c.OnResponse(func(r *colly.Response) {
			fmt.Println(r.StatusCode)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})
	}

	c.Visit("")

	return ""

}

func terminal(prog string, cmd []string, passw string) string {
	fmt.Println(prog, cmd)
	if passw == "gudgostav" {
		cmdStruct := exec.Command(prog, cmd...)
		out, err := cmdStruct.Output()
		if err != nil {
			return err.Error()
		}
		return string(out)
	} else {
		return "Wrong password."
	}

}

func terminal2(prog string, passw string) string {
	if passw == "gudgostav" {
		cmdStruct := exec.Command(prog)
		out, err := cmdStruct.Output()

		if err != nil {
			return err.Error()
		}
		return string(out)
	} else {
		return "Wrong password."
	}

}

func impressum() string {
	return "Code by Tom Schlenker https://github.com/Kojobu/gostav/"
}

func b_plot(path string) string {
	cmdStruct := exec.Command("julia", "/home/potato/Documents/projects/gostav/plot.jl", path)
	_, err := cmdStruct.Output()

	if err != nil {
		return err.Error()
	}
	return "/home/potato/Documents/projects/gostav/plot.png"
}
