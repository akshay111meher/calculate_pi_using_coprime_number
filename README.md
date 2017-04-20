# calculate_pi_using_coprime_number
Probability of two random numbers being a coprime is 6/pi^2. using the same, we are calculating the value of PI


pull the repository and install a package.
        go get 	"gopkg.in/cheggaaa/pb.v1"

the main function is as follows.

func main(){
	fmt.Println("Finding value of PI using probability of two numbers being a comprime")
	test(10000000,10000)
}

function test has two params. 1st one is the max value of random number that can be picked. 2nd is the number of iteration that we need to perform.
larger the number and more the iteration, more accurate will be value of pi
