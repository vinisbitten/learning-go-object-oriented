package Vehicle

type Car struct {
	Name string
	Brand string
	Year  int
	topSpeed int
}

func ImportedCar() Car{
	myCar := Car{
		Name: "Gol",
		Brand: "Volkswagen",
		Year: 2010,
		topSpeed: 200,
	}

	return myCar
}