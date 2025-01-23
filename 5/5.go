package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SortCars(cars []car) {
	sort.Slice(cars, func(i, j int) bool {
		if cars[i].StTime != cars[j].StTime {
			return cars[i].StTime < cars[j].StTime // Сортировка по StTime (по возрастанию)
		}
		return cars[i].Id < cars[j].Id // Если StTime равны, сортируем по Id (по возрастанию)
	})
}

type order struct {
	orderId int
	arrTime int
	CarNum  int
	viewed  bool
} // Класс заказ

type car struct {
	StTime  int
	EndTime int
	Cap     int
	Id      int
} // Класс машина
func (c *car) checkTime(t int) bool {
	return c.StTime <= t && t <= c.EndTime
} // проверка на время в границах времени стоянки
func (c *car) checkCap() bool {
	return c.Cap > 0
} // проверка вместимости
func (c *car) decCap() {
	c.Cap--
} // уменьшаем вместимость если выдали заказ

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int
	fmt.Fscanln(in, &num)

	for i := 0; i < num; i++ {
		//получили кол-во заказов
		var int_of_el int
		fmt.Fscanln(in, &int_of_el)

		input_string, _ := in.ReadString('\n')
		input_string = strings.TrimSpace(input_string)
		var mas_el []string = strings.Split(input_string, " ")

		//заполнили мапу с ордерами
		order_map := make([]order, int_of_el)
		for i := 0; i < len(mas_el); i++ {
			elem, _ := strconv.Atoi(mas_el[i])
			order_map[i] = order{i, elem, -1, false}
		}

		var int_of_cars int
		fmt.Fscanln(in, &int_of_cars)

		carMap := make([]car, int_of_cars)
		for i := 0; i < int_of_cars; i++ {
			var elem1, elem2, elem3 int
			fmt.Fscanln(in, &elem1, &elem2, &elem3)
			carMap[i] = car{elem1, elem2, elem3, i + 1}
		} //есть мапа машин
		SortCars(carMap)

		for i := 0; i < int_of_el; i++ {

			minTime := math.MaxInt32
			var ourOrder order
			for k := 0; k < int_of_el; k++ {
				if order_map[k].arrTime <= minTime && order_map[k].CarNum == -1 && order_map[k].viewed == false {
					minTime = order_map[k].arrTime
					ourOrder = order_map[k]
				}
			}
			// нашли минимальное время прибытия в ourOrder найденный заказ
			minTime = math.MaxInt32

			var car1 car = car{-1, -1, -1, -1}

			for k := 0; k < int_of_cars; k++ {
				if carMap[k].checkTime(ourOrder.arrTime) == true && carMap[k].checkCap() == true {
					car1 = carMap[k]
					carMap[k].decCap()
					ourOrder.CarNum = carMap[k].Id
					break
				}
			}

			if car1.Id == -1 {
				for k := 0; k < int_of_el; k++ {
					if order_map[k].orderId == ourOrder.orderId {
						order_map[k].viewed = true
						break
					}
				}
				continue
			} else {

				for k := 0; k < int_of_el; k++ {
					if order_map[k].orderId == ourOrder.orderId {
						order_map[k].CarNum = ourOrder.CarNum
						order_map[k].viewed = true
						break
					}
				}
			}

		}
		for i := 0; i < int_of_el; i++ {
			if i == int_of_el-1 {
				fmt.Print(order_map[i].CarNum, "\n")
			} else {
				fmt.Print(order_map[i].CarNum, " ")
			}
		}

	}
}
