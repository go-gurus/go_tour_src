package main

import (
	"os"
	"text/template"
)

type Data struct {
	Distance     float64
	DistanceUnit string
	Speed        float64
	SpeedUnit    string
	Time         float64
	TimeUnit     string
}

func main() {
	tmpl, err := template.New("example").Parse(`
Distance: {{.Distance}} {{.DistanceUnit}}
Speed:    {{.Speed}} {{.SpeedUnit}}
You are travelling at a speed of {{.Speed}} {{.SpeedUnit}}.
You will cover a distance of {{.Distance}} {{.DistanceUnit}} in {{.Time}} {{.TimeUnit}}.`)
	if err != nil {
		panic(err)
	}

	// metric system
	metric_data := Data{1000, "km", 100, "km/h", 10, "h"}
	err = tmpl.Execute(os.Stdout, metric_data)
	if err != nil {
		panic(err)
	}

	// imperial system
	imperial_data := Data{621.371, "mi", 62.1371, "mph", 10, "h"}
	err = tmpl.Execute(os.Stdout, imperial_data)
	if err != nil {
		panic(err)
	}

	//maritime/aviation system
	maritime_data := Data{539.957, "nm", 53.9957, "kn", 10, "h"}
	err = tmpl.Execute(os.Stdout, maritime_data)
	if err != nil {
		panic(err)
	}
}
