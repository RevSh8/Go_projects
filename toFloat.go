package toFloat

import (
	"bufio"
	"strings"
	"os"
	"strconv"
)

func ToFloat()(float64, error){
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err!=nil{
		return 0, err
	}
	input = strings.TrimSpace(input)
	res, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}