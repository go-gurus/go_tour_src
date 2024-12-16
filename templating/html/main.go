package main

import (
	"html/template"
	"os"
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
<!DOCTYPE html>
    <title>Travel Information</title>
</head>
<body>
    <h1>Travel Information</h1>
    <p>Distance: {{.Distance}} {{.DistanceUnit}}</p>
    <p>Speed: {{.Speed}} {{.SpeedUnit}}</p>
    <p>You are travelling at a speed of {{.Speed}} {{.SpeedUnit}}.</p>
    <p>You will cover a distance of {{.Distance}} {{.DistanceUnit}} in {{.Time}} {{.TimeUnit}}.</p>
</body>
</html>
`)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("output.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	metricData := Data{1000, "km", 100, "<script>alert('XSS-Angriff!')</script>", 10, "h"}
	err = tmpl.Execute(file, metricData)
	if err != nil {
		panic(err)
	}

	println("HTML file 'output.html' has been created.")
}