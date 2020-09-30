package App
import "strconv"

func Calculate(x int) (result string) {
	if (x%3==0 && x%5!=0){
		result := "Fizz"
		return result
	}
	if (x%5==0 && x%3!=0){
		result := "Buzz"
		return result
	}
	if (x%5==0 && x%3==0){
		result := "FizzBuzz"
		return result
	}
	result=strconv.Itoa(x)
	return result
}