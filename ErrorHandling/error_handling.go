package main

import (
	//"errors"
	"fmt"
)

type customError struct{
	code int
	message string
}
func (c customError) Error() string{
	return fmt.Sprintf("Error %d: %s", c.code, c.message)
}

func errorHandling(param int) (int,error)  {
	if param == 60{
		//return -1, errors.New("i dont like 60")
		return -1, customError{code: 1, message: "i dont like 60"}
	} else{
		return param + 10, nil
	}
}
func main()  {
	var number int
	_, err1 := fmt.Scan(&number)
	if err1 != nil { // means if there is an error ...
		fmt.Println("we got an error")
		fmt.Println(err1.Error())
	} else{
		fmt.Println(number)
	}
	ret, err2 := errorHandling(80)
	if err2 != nil{ // ==nil means ;[if nothing went wrong continue (nil means nothing/no value)
		fmt.Println("Error encountered")
		fmt.Println(err2)
	}else{
		fmt.Println(ret)
	}

	result, err := divide(10, 3)
	if err != nil{
		fmt.Println("error", err)
		return
	}
	fmt.Println("result", result)
}

func divide(a, b float64) (float64, error)  {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by 0 %f", b)//use errors.new when u want to return a new error but errorf wen u want to return a formatted string
	}
	return a / b, nil
}
