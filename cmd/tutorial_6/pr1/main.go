package main

import "fmt"

var emo = struct {
	lying    string // "\U0001F925"
	vomiting string // "\U0001F922"
	hot      string // "\U0001F975"
	drooling string // "\U0001F924"
	poo      string // "\U0001F4A9"
	neutral  string // "\U0001F610"
}{lying: "\U0001F925", vomiting: "\U0001F922", hot: "\U0001F975", drooling: "\U0001F924", poo: "\U0001F4A9", neutral: "\U0001F610"}

type engine struct {
	name     string
	cylinder uint8
	capacity float32
	hp       uint16
	torque   uint16
}
type car struct {
	name      string
	engine    engine
	top_speed uint16
}

func rating_car(car car) {
	fmt.Print("you engine is : ")
	switch car.engine.cylinder {
	case 2:
		fmt.Println(emo.vomiting)
	case 3:
		fmt.Println(emo.poo)
	case 4:
		fmt.Println(emo.neutral)
	case 6, 5:
		fmt.Println(emo.drooling)
	case 8:
		fmt.Println(emo.hot)
	default:
		fmt.Println(emo.lying)
	}
}

func main() {
	var eng engine = engine{name: "v8", cylinder: 8, capacity: 4.5, hp: 555, torque: 800}
	fmt.Println(eng)
	var car1 car = car{name: "g63", engine: eng, top_speed: 350}
	fmt.Println(car1)
	rating_car(car1)
	var eng2 engine = engine{name: "v8", cylinder: 16, capacity: 8.0, hp: 1600, torque: 800}
	var car2 car = car{name: "bugatti", engine: eng2, top_speed: 450}
	rating_car(car2)

}
