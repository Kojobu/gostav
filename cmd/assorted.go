package cmd

import (
	"fmt"
	"os/exec"
	"time"



	"github.com/gocolly/colly"
)

func mensa_scrap(days int) string {

	c := colly.NewCollector(
	//colly.AllowedDomains("studentenwerk.uni-heidelberg.de"),
	)

	type Food struct {
		date []string
		meal []string
	}
	food := Food{}
	c.OnHTML("table", func(e *colly.HTMLElement){
		found_food := false
		e.ForEach("tr", func(i int, el *colly.HTMLElement) {
			//fmt.Println(i, el.Text)
			if i == 7 {
				food.meal = append(food.meal, el.ChildText("td:nth-child(1)"))
				found_food = !found_food
			}
		})
		if !found_food {
			food.meal = append(food.meal, "Die Mensa ist an diesem Tag geschlossen.")
		}
	})

	returnstring := ""
	for i := 0; i<days; i++ {
		food.date = append(food.date, time.Now().AddDate(0,0,i).Format("02.01.2006"))
		c.Visit("https://www.stw.uni-heidelberg.de/external-tools/speiseplan/speiseplan.php?lang=de&mode=Mensa+Im+Neuenheimer+Feld+304&date=" + food.date[i])
		returnstring += food.date[i] + "\n"
		returnstring += food.meal[i] + "\n\n"
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
