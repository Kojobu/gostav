package cmd

import (
	"fmt"
	"os/exec"
	"time"
	"net/http"
	"io"
	"encoding/json"
	"strconv"



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

func ozon_scrap() string {

	type Ozon struct {
		Table struct {
			Header []struct {
				Label   string `json:"label"`
				Abbr    string `json:"abbr"`
				Options struct {
					MinWidth            string `json:"minWidth"`
					Align               string `json:"align"`
					Sorting             bool   `json:"sorting"`
					Highlighted         bool   `json:"highlighted"`
					Template            string `json:"template"`
					TemplateHighlighted string `json:"templateHighlighted"`
				} `json:"options,omitempty"`
				Tooltip  string `json:"tooltip,omitempty"`
				Options0 struct {
					Sorting  bool   `json:"sorting"`
					DataType string `json:"dataType"`
				} `json:"options,omitempty"`
				SubHeaders []struct {
					Label   string `json:"label"`
					Abbr    string `json:"abbr"`
					Options struct {
						Align    string `json:"align"`
						Sorting  bool   `json:"sorting"`
						DataType string `json:"dataType"`
						Limits   struct {
							Warning float64 `json:"warning"`
							Error   float64 `json:"error"`
						} `json:"limits"`
						Template string `json:"template"`
					} `json:"options"`
				} `json:"subHeaders,omitempty"`
				Options1 struct {
					Align   string `json:"align"`
					Sorting bool   `json:"sorting"`
				} `json:"options,omitempty"`
				Options2 struct {
					Align   string `json:"align"`
					Sorting bool   `json:"sorting"`
				} `json:"options,omitempty"`
			} `json:"header"`
			Data []struct {
				Data       []any `json:"data"`
				Parameters struct {
					ID         string `json:"id"`
					O3Measured bool   `json:"O3-measured"`
				} `json:"parameters"`
				Highlighted bool `json:"highlighted,omitempty"`
			} `json:"data"`
		} `json:"table"`
	}

	resp, err := http.Get("https://lupo-cloud.de/air-app/table?component=5&id=DEBW009")
	if err != nil {
		// handle err
	}
	var ozon Ozon
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal(body, &ozon)
	ozon_hd := ozon.Table.Data[9].Data
	returnstring := ""
	returnstring += "Ozonvales for " + ozon_hd[0].(string) + " in [µg/m3]\n"
	returnstring += "1h mean: " + strconv.FormatFloat(ozon_hd[1].(float64), 'f', -1, 64)+"\n"
	returnstring += "1h max: " +  ozon_hd[2].(string) +"\n"
	returnstring += "1h max (yesterday): " +  ozon_hd[3].(string)
 	return returnstring

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
