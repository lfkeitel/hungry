package main

import (
	"bufio"
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	bindaddr string
	bindport string
	places   string
)

func init() {
	flag.StringVar(&bindaddr, "addr", "0.0.0.0", "Bind address")
	flag.StringVar(&bindport, "port", "8081", "Bind port")
	flag.StringVar(&places, "f", "places.txt", "File list of restaurants")

	rand.Seed(time.Now().Unix())
}

func main() {
	flag.Parse()

	restaurants := loadRestaurants(places)
	choiceSlice := makeChoiceSlice(restaurants)

	http.HandleFunc("/", indexHandler(choiceSlice))
	http.HandleFunc("/settings", currentSettingsHandler(restaurants))

	log.Print("Now listening")
	http.ListenAndServe(fmt.Sprintf("%s:%s", bindaddr, bindport), nil)
}

func indexHandler(choices []*restaurant) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := choices[rand.Intn(len(choices))]
		indexTmpl.Execute(w, res)
	}
}

func currentSettingsHandler(choices []*restaurant) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, r := range choices {
			w.Write([]byte(fmt.Sprintf("Name: %s\nSchedule: %s\nWeight: %d\n\n", r.Name, r.schedule.String(), r.score)))
		}
	}
}

type openDays struct {
	monday, tuesday, wednesday, thursday, friday, saturday, sunday bool
}

func openDaysFromString(days string) *openDays {
	if days == "" {
		return &openDays{true, true, true, true, true, true, true}
	}

	openDay := &openDays{}
	ranges := strings.Split(days, ",")
	for _, r := range ranges {
		switch strings.ToLower(r) {
		case "m":
			openDay.monday = true
		case "t":
			openDay.tuesday = true
		case "w":
			openDay.wednesday = true
		case "r", "th":
			openDay.thursday = true
		case "f":
			openDay.friday = true
		case "a", "sa":
			openDay.saturday = true
		case "u", "su":
			openDay.sunday = true
		default:
			return nil
		}
	}

	return openDay
}

func (o *openDays) String() string {
	return fmt.Sprintf(`
  Monday: %t
  Tuesday: %t
  Wednesday: %t
  Thursday: %t
  Friday: %t
  Saturday: %t
  Sunday: %t`, o.monday, o.tuesday, o.wednesday, o.thursday, o.friday, o.saturday, o.sunday)
}

type restaurant struct {
	Name     string
	schedule *openDays
	score    int
}

func loadRestaurants(filename string) []*restaurant {
	choices := make([]*restaurant, 0, 20)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linenum := 0

	for scanner.Scan() {
		linenum++

		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		lineItems := strings.SplitN(scanner.Text(), ":", 3)
		if len(lineItems) != 3 {
			log.Fatalf("error on line %d, missing section", linenum)
		}

		name := lineItems[0]
		scheduleStr := lineItems[1]
		score := lineItems[2]

		schedule := openDaysFromString(scheduleStr)
		if schedule == nil {
			log.Fatalf("error on line %d, invalid schedule", linenum)
		}

		if score == "" {
			score = "0"
		}

		scoreCost, err := strconv.Atoi(score)
		if err != nil || scoreCost < -4 || scoreCost > 5 {
			log.Fatalf("error on line %d, invalid score, must be between -4 to 5", linenum)
		}

		r := &restaurant{
			Name:     name,
			schedule: schedule,
			score:    scoreCost,
		}

		choices = append(choices, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading restaurant file")
	}

	return choices
}

func makeChoiceSlice(restaurants []*restaurant) []*restaurant {
	choices := make([]*restaurant, 0, 20)

	for _, r := range restaurants {
		multiplier := 5 + r.score
		for i := 0; i < multiplier; i++ {
			choices = append(choices, r)
		}
	}

	return choices
}

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">

	<title>Where to eat?</title>

	<style>
		html, body {
			height: 100%;
		}

		body {
			background-color: lawngreen;
			font-family: sans-serif;
			display: flex;
			align-items: center;
			justify-content: center;
			align-content: center;
			text-align: center;
		}

		h1 {
			font-size: 3rem;
		}
	</style>
</head>

<body>
	<h1>{{.Name}}</h1>

	<script type="text/javascript">
		window.addEventListener("click", () => location.reload());
		window.addEventListener("mousedown", () => document.body.style.backgroundColor = "rgb(162, 255, 73)");
	</script>
</body>
</html>
`))
