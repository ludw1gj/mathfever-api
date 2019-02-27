package mathematics

import (
	"fmt"
	"math"
	"strings"
)

// IsPrime checks if n is a prime number and returns if is/not a prime and why.
func IsPrime(n int) (string, error) {
	var answer string
	if n <= 0 {
		answer = fmt.Sprintf("%d is not a prime number, because prime numbers are defined for "+
			"integers greater than 1.", n)
	} else if n == 1 {
		answer = "1 is not considered to be a prime number."
	} else if n == 2 {
		answer = "2 is a prime number, because its only whole-number factors are 1 and itself."
	} else if n%2 == 0 {
		answer = fmt.Sprintf("%d is not a prime number, because it is divisible by 2.", n)
	} else {
		sqr := int(math.Sqrt(float64(n))) + 1
		for i := 3; i < sqr; {
			if n%i == 0 {
				answer = fmt.Sprintf("%d is not a prime number, because it is divisible by %d.", n, i)
				break
			}
			i += 2
		}
	}
	if len(answer) == 0 {
		answer = fmt.Sprintf("%d is a prime number, because its only whole-number factors are 1 and "+
			"itself.", n)
	}

	data := struct {
		Number int
		Answer string
	}{
		n,
		answer,
	}

	const tpl = `
	<h4>Question:</h4>
	<p>Is {{.Number}} a prime number?</p>

	<h4>Answer:</h4>
	<p>{{.Answer}}</p>

	<h4>Helpful Tips:</h4>
	<p>Prime numbers are numbers whose only whole-number factors are 1 and itself.</p>
	<p>To determine if a number (n) is a prime number, here are some rules:</p>
	<ul>
		<li>Prime numbers are defined for whole-numbers greater than 1.</li>
		<li>1 is not considered a prime number.</li>
		<li>2 is considered a prime number because its only whole-number factors are 1 and itself.</li>
		<li>Any number that is divisible by 2 is not a prime number.</li>
		<li>After the steps above, start by dividing the number by 3, then 5, then 7, and so on, 
			checking if the number can be divided by those divisors. As we have determined it cannot be 
			divided by 2, there is no need to divide by 4, 6, 8, and so on.
		</li>
		<li>The maximum divisor to look for is the square root of your number, because: n=a&timesb and 
			if both values were greater than the square root of n, a&timesb would be larger than n. 
			Therefore at least one of those factors must be less than or equal to the square root of n.
		</li>
	</ul>`

	return parseTemplate(tpl, data)
}

func findPrimeFactors(n int) (primeFactors []int, table string, proof string,
	factorFrequency map[int]int) {
	generatePrimeFactorsAndTable := func(num int) ([]int, string) {
		var primes []int
		var tableBuf strings.Builder

		fmt.Fprint(&tableBuf, `
		<table class="mdl-data-table mdl-js-data-table">
		<tbody>`)
		for i := 2; i*i <= num; {
			if num%i == 0 {
				primes = append(primes, i)
				fmt.Fprintf(&tableBuf, `
			<tr>
				<td>%d | %d</td>
				<td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
				<td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td>
			</tr>`, num, i, i, num, num, i, num/i)

				num /= i
			} else {
				i++
			}
		}

		if num > 1 {
			primes = append(primes, num)
			fmt.Fprintf(&tableBuf, `
			<tr>
			<td>%d | %d</td>
			<td class="column-wrap mdl-data-table__cell--non-numeric">%d is a factor of %d</td>
			<td class="column-wrap mdl-data-table__cell--non-numeric">%d divided by %d is %d</td>
			</tr>`, num, num, num, num, num, num, 1)
		}

		fmt.Fprint(&tableBuf, `
		<tr>
		<td>1&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</td>
		<td></td>
		<td></td>
		</tr>
		</tbody></table>`)
		return primes, tableBuf.String()
	}

	generateProof := func(num int, primes []int, frequency map[int]int) string {
		var proofBuf strings.Builder
		for _, factor := range primes {
			if len(primes) > 0 {
				fmt.Fprintf(&proofBuf, "%d", factor)
			} else {
				fmt.Fprintf(&proofBuf, "%d &times ", factor)
			}
		}
		fmt.Fprintf(&proofBuf, "= %d<br>", num)

		for factor := range frequency {
			fmt.Fprintf(&proofBuf, "%d<sup>%d</sup> &times ", factor, frequency[factor])
		}

		proof := fmt.Sprintf("%s = %d", strings.TrimSuffix(proofBuf.String(), " &times "), num)
		return proof
	}

	primeFactors, table = generatePrimeFactorsAndTable(n)
	factorFrequency = findElementFrequency(primeFactors)
	proof = generateProof(n, primeFactors, factorFrequency)
	return primeFactors, table, proof, factorFrequency
}

// HighestCommonFactor outputs the proof and answer of calculating the highest
// common factor of two numbers.
func HighestCommonFactor(n1 int, n2 int) (string, error) {
	primeFactors1, table1, proof1, _ := findPrimeFactors(n1)
	primeFactors2, table2, proof2, _ := findPrimeFactors(n2)
	commonNumbers := findCommonIntegers(primeFactors1, primeFactors2)

	answer := 1
	var sharedPrimesBuf strings.Builder
	var sharedPrimesProofBuf strings.Builder
	var sharedPrimesProof string
	if len(commonNumbers) != 0 {
		// shared primes
		for _, n := range commonNumbers {
			fmt.Fprintf(&sharedPrimesBuf, "%d, ", n)
		}

		// shared primes proof
		for _, n := range commonNumbers {
			fmt.Fprintf(&sharedPrimesProofBuf, "%d &times ", n)
			answer *= n
		}
		sharedPrimesProof = fmt.Sprintf("%s = %d", strings.TrimSuffix(sharedPrimesProofBuf.String(),
			" &times "), answer)
	} else {
		message := "There are no shared factors."
		fmt.Fprint(&sharedPrimesBuf, message)
		fmt.Fprint(&sharedPrimesProofBuf, message)
		sharedPrimesProof = sharedPrimesProofBuf.String()
	}
	// sharedPrimes might have ", " at the end
	sharedPrimes := strings.TrimSuffix(sharedPrimesBuf.String(), ", ")

	data := struct {
		FirstNumber       int
		SecondNumber      int
		Table1            string
		Table2            string
		Proof1            string
		Proof2            string
		SharedPrimes      string
		SharedPrimesProof string
		Answer            int
	}{
		n1,
		n2,
		table1,
		table2,
		proof1,
		proof2,
		sharedPrimes,
		sharedPrimesProof,
		answer,
	}

	const tpl = `
	<h4>Question:</h4>
	<p>Find the highest common factor of {{.FirstNumber}} and {{.SecondNumber}}.</p>

	<h4>Answer:</h4>
	<p>The highest common factor of {{.FirstNumber}} and {{.SecondNumber}} is {{.Answer}}.</p>

	<h4>Here's the working out:</h4>
	<p>Finding all prime factors of {{.FirstNumber}}:</p>
	{{.Table1}}
	<p>Finding all prime factors of {{.SecondNumber}}:</p>
	{{.Table2}}
	<p>Prime factors for the first number are:<br>
		{{.Proof1}}</p>
	<p>Prime factors for the second number are:<br>
		{{.Proof2}}</p>
	<p>Find the primes that are shared between the two numbers:<br>
		{{.SharedPrimes}}</p>
	{{if .SharedPrimesProof}}
	<p>Take the shared primes and multiply them together:<br>
		{{.SharedPrimesProof}}</p>
	{{end}}

	<h4>Therefore:</h4>
	<p>The highest common factor of {{.FirstNumber}} and {{.SecondNumber}} is {{.Answer}}.</p>`

	return parseTemplate(tpl, data)
}

// LowestCommonMultiple outputs the proof and answer of calculating the lowest
// common multiple of two numbers.
func LowestCommonMultiple(n1 int, n2 int) (string, error) {
	_, table1, proof1, factorFrequency1 := findPrimeFactors(n1)
	_, table2, proof2, factorFrequency2 := findPrimeFactors(n2)

	// find sets with the highest exponent
	m := make(map[int]int)
	for i := range factorFrequency1 {
		if factorFrequency1[i] > factorFrequency2[i] {
			m[i] = factorFrequency1[i]
		} else {
			m[i] = factorFrequency2[i]
		}
	}
	for i := range factorFrequency2 {
		if factorFrequency1[i] > factorFrequency2[i] {
			m[i] = factorFrequency1[i]
		} else {
			m[i] = factorFrequency2[i]
		}
	}

	var compareProofBuf strings.Builder
	answer := 1
	for factor := range m {
		fmt.Fprintf(&compareProofBuf, "%d<sup>%d</sup> &times ", factor, m[factor])
		answer *= int(math.Pow(float64(factor), float64(m[factor])))
	}
	compareProof := fmt.Sprintf("%s = %d", strings.TrimSuffix(compareProofBuf.String(), " &times "),
		answer)

	data := struct {
		FirstNumber  int
		SecondNumber int
		Table1       string
		Table2       string
		Proof1       string
		Proof2       string
		CompareProof string
		Answer       int
	}{
		n1,
		n2,
		table1,
		table2,
		proof1,
		proof2,
		compareProof,
		answer,
	}

	const tpl = `
	<h4>Question:</h4>
	<p>Find the lowest common multiple of {{.FirstNumber}} and {{.SecondNumber}}.</p>

	<h4>Answer:</h4>
	<p>The lowest common multiple of {{.FirstNumber}} and {{.SecondNumber}} is {{.Answer}}.</p>

	<h4>Here's the working out:</h4>
	<p>Finding all prime factors of {{.FirstNumber}}:</p>
	{{.Table1}}
	<p>Finding all prime factors of {{.SecondNumber}}:</p>
	{{.Table2}}
	<p>Prime factors for the first number are:<br>
		{{.Proof1}}</p>
	<p>Prime factors for the second number are:<br>
		{{.Proof2}}</p>
	<p>Compare the primes of the first and second numbers and use the sets with the highest exponent:
		<br>
		{{.CompareProof}}</p>

	<h4>Therefore:</h4>
	<p>The lowest common multiple of {{.FirstNumber}} and {{.SecondNumber}} is {{.Answer}}.</p>`

	return parseTemplate(tpl, data)
}
