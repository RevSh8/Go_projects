//Package that contains only one func "ToFloat"
package toFloat

import (
	"bufio"
	"strings"
	"os"
	"strconv"
)
//Function that let you to acquire a float64 num from user and
//get it in your project. Also you can get an error.
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