package services

import (
	"math"
	"math/rand"

	"github.com/BigbossXD/calculate-xyz/models/responses"
)

type CalculateService struct {
}

func NewCalculateService() *CalculateService {
	return &CalculateService{}
}

var X int
var Y int
var Z int
var dataSet = []int{1, X, 8, 17, Y, Z, 78, 113}

func mappingdata(x int, y int, z int) []int {
	dataSet[1] = X
	dataSet[4] = Y
	dataSet[5] = Z
	if X < dataSet[0] {
		X = dataSet[0]
	}
	if Y < dataSet[3] {
		Y = dataSet[3]
	}
	if Z < dataSet[4] {
		Z = dataSet[4]
	}
	return dataSet
}

func mappingNegativeData(x int, y int, z int) []int {
	dataSet[1] = X
	dataSet[4] = Y
	dataSet[5] = Z

	return dataSet
}

func randomOddNumber(min int, max int) int {
	r := rand.Intn(max-min) + min
	if math.Mod(float64(r), 2) == 0 {
		return r - 1
	}
	return r
}

func randomEvenNumber(min int, max int) int {
	r := rand.Intn(max-min) + min
	if math.Mod(float64(r), 2) == 0 {
		return r
	}
	return r - 1
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func (s *CalculateService) Fixed() (*responses.BasicResultResponse, error) {
	X = 4
	Y = 36
	Z = 54
	result := &responses.BasicResultResponse{X: 4, Y: 36, Z: 54, DataSet: mappingdata(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) BasicRandomInt() (*responses.BasicResultResponse, error) {
	X = rand.Intn(200)
	Y = rand.Intn(200)
	Z = rand.Intn(200)
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingdata(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) BasicRandomIntSort() (*responses.BasicResultResponse, error) {
	X = rand.Intn(dataSet[2]-dataSet[0]) + dataSet[0]
	Y = rand.Intn((dataSet[6]-1)-dataSet[3]) + dataSet[3]
	Z = rand.Intn((dataSet[6])-Y) + Y
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingdata(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) RandomOdd() (*responses.BasicResultResponse, error) {
	X = randomOddNumber(0, 200)
	Y = randomOddNumber(0, 200)
	Z = randomOddNumber(0, 200)
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingdata(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) RandomOddSort() (*responses.BasicResultResponse, error) {
	X = randomOddNumber(dataSet[0], dataSet[2])
	Y = randomOddNumber(dataSet[3], dataSet[6]-1)
	Z = randomOddNumber(Y, dataSet[6])
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingdata(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) RandomEven() (*responses.BasicResultResponse, error) {
	X = randomEvenNumber(0, 200)
	Y = randomEvenNumber(0, 200)
	Z = randomEvenNumber(0, 200)
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingdata(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) RandomEvenSort() (*responses.BasicResultResponse, error) {
	X = randomEvenNumber(dataSet[0], dataSet[2])
	Y = randomEvenNumber(dataSet[3], dataSet[6]-1)
	Z = randomEvenNumber(Y, dataSet[6])
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingdata(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) RandomNegativeInt() (*responses.BasicResultResponse, error) {
	X = rand.Intn(200) * -1
	Y = rand.Intn(200) * -1
	Z = rand.Intn(200) * -1
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingNegativeData(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) RandomNegativeOddInt() (*responses.BasicResultResponse, error) {
	X = randomOddNumber(0, 200) * -1
	Y = randomOddNumber(0, 200) * -1
	Z = randomOddNumber(0, 200) * -1
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingNegativeData(X, Y, Z)}
	return result, nil
}

func (s *CalculateService) RandomNegativeEvenInt() (*responses.BasicResultResponse, error) {
	X = randomEvenNumber(0, 200) * -1
	Y = randomEvenNumber(0, 200) * -1
	Z = randomEvenNumber(0, 200) * -1
	result := &responses.BasicResultResponse{X: X, Y: Y, Z: Z, DataSet: mappingNegativeData(X, Y, Z)}
	return result, nil
}
