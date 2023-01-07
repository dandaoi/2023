package main

import (
	"bufio"
	"fmt"
	"os"
)

type food []string
type rated []string

func newFood() food {

	plates := food{}

	mistura := []string{"Carne", "Ovo", "Lasanha", "Escondidinho", "Frango"}
	carbs := []string{"Arroz", "Macarrão", "Batata", "Salada", "Feijoada"}
	acomp := []string{"Feijão", "Molho Branco", "Molho Bolonhesa", "Atum", "Banana"}

	for _, carb := range carbs {
		for _, ac := range acomp {
			for _, mis := range mistura {
				plates = append(plates, carb+" com "+ac+" e "+mis)

			}
		}
	}
	return plates

}

func (f food) linkRate() rated {

	rtd := rated{}

	for _, fud := range f {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("%v Avalie essa comida: ", fud)
		scanner.Scan()
		fmt.Println(scanner.Text())

		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}

		rtd = append(rtd, fud+" tem uma avaliação de: "+scanner.Text())

	}
	fmt.Println(rtd)
	return rtd

}
