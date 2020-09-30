package App
import(
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestDevuleveNumeroEnString(test *testing.T) {
	result:=Calculate(2)
	assert.Equal(test,result,"2")
}  
func TestFizz(test *testing.T) {
	result:=Calculate(3)
	assert.Equal(test,result,"Fizz")
}  
func TestBuzz(test *testing.T) {
	result:=Calculate(5)
	assert.Equal(test,result,"Buzz")
}  
func TestFizzBuzz(test *testing.T) {
	result:=Calculate(15)
	assert.Equal(test,result,"FizzBuzz")
} 