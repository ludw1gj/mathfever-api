package math

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/ludw1gj/mathfever/util"
)

// BinaryToDecimal outputs the proof and answer of a binary to decimal conversion.
func BinaryToDecimal(binary string) (string, error) {
	const base = 2
	lenBinary := len(binary)

	power := lenBinary - 1
	var proofStepsBuf [4]bytes.Buffer
	var answer int
	for _, digit := range binary {
		digit, _ := strconv.ParseInt(string(digit), base, 0)
		power := int(math.Pow(base, float64(power)))
		stepValue := int(digit) * power

		fmt.Fprintf(&proofStepsBuf[0], "(b=%d x %d<sup>n=%d</sup>) + ", digit, base, power)
		fmt.Fprintf(&proofStepsBuf[1], "(%d x %d<sup>%d</sup>) + ", digit, base, power)
		fmt.Fprintf(&proofStepsBuf[2], "(%d x %d) + ", digit, power)
		fmt.Fprintf(&proofStepsBuf[3], "(%d) + ", stepValue)

		answer += stepValue
		power--
	}

	// remove " + " at end of strings and write string to proofSteps
	var proofSteps [4]string
	for i, ps := range proofStepsBuf {
		proofStepsBuf[i].Truncate(len(ps.String()) - 3)
		proofSteps[i] = proofStepsBuf[i].String()
	}

	data := struct {
		Binary      string
		LenBinary   int
		MaxPosition int
		Proof       [4]string
		Answer      int
	}{
		binary,
		lenBinary,
		lenBinary - 1,
		proofSteps,
		answer,
	}

	const tpl = `
	<h6>Question:</h6>

	<p>What is {{.Binary}} in decimal?</p>

	<h6>Answer:</h6>

	<p>The binary number <b>{{.Binary}}</b> in decimal is <b>{{.Answer}}</b>.</p>

	<h6>Here's how to convert it:</h6>

	<p>Using the following continuing equation the accurate answer can be found: (b x 2<sup>n</sup>)Ã¢â‚¬Â¦<br>
			Where <strong>b</strong> is the binary number at the n<sup>th</sup> position and <strong>n</strong> is the
			n<sup>th</sup> position.<br>
			There are {{.LenBinary}} numbers in this binary number, the n<sup>th</sup> position starts from 0, and therefore
			the n<sup>th</sup> position goes up to {{.MaxPosition}} ({{.LenBinary}} - 1). Count down from {{.MaxPosition}}
			and insert the binary number at that position.</p>

	<p>Using the equation for the binary number <span class="word-break">{{.Binary}}</span>, it would go like this:<br>
			{{range .Proof}}{{.}}<br>{{end}}
			<span class="word-break">{{.Answer}}</span></p>

	<h6>Therefore:</h6>

	<p class="word-break">({{.Binary}})<sub>2</sub> = ({{.Answer}})<sub>10</sub></p>`

	return util.ParseTemplate(tpl, data)
}

// BinaryToHexadecimal outputs the proof and answer of a binary to hexadecimal conversion.
func BinaryToHexadecimal(binary string) (string, error) {
	zeroedBinary := binary
	nLength := len(zeroedBinary)

	// if input doesn't divide into groups of length of 4, add 0"s to the beginning of the string
	if nLength%4 != 0 {
		n := 4 - (nLength % 4)
		for i := 0; i < n; i++ {
			zeroedBinary = fmt.Sprintf("0%s", zeroedBinary)
		}
	}

	// split input into groups of 4 digits
	var groupedBinary []string
	for i := 0; i < len(zeroedBinary)/4; i++ {
		groupedBinary = append(groupedBinary, zeroedBinary[4*i:4*(i+1)])
	}

	var proof bytes.Buffer
	var answer bytes.Buffer
	for _, i := range groupedBinary {
		decimalAnswer, err := strconv.ParseInt(i, 2, 0)
		if err != nil {
			return "", fmt.Errorf("incorrect input: is not a binary number: %s", binary)
		}
		fmt.Fprintf(&proof, "(%s)<sub>2</sub> = (%d)<sub>10</sub> = (%X)<sub>16</sub><br>",
			i, decimalAnswer, decimalAnswer)
		fmt.Fprintf(&answer, "%X", decimalAnswer)
	}

	data := struct {
		Binary        string
		GroupedBinary []string
		Proof         string
		Answer        string
	}{
		binary,
		groupedBinary,
		proof.String(),
		answer.String(),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>What is {{.Binary}} in hexadecimal?</p>

	<h6>Answer:</h6>

	<p>The binary number <b>{{.Binary}}</b> in hexadecimal is <b>{{.Answer}}</b>.</p>

	<h6>Here's how to convert it:</h6>

	<p>Break the binary number <span class="word-break">{{.Binary}}</span> into groups of 4 digits as such:<br>
	{{range .GroupedBinary}}{{.}}{{end}}</p>

	<p>Convert those groups (base 2) to a decimal number (base 10), that way it is easy to convert them to a Hexadecimal
	number (base 16):<br>
	Refer to the Conversion Table <a href="/category/networking/conversion-table" rel="noopener noreferrer"
															target="_blank">here</a>.<br>
	{{.Proof}}</p>

	<h6>Therefore:</h6>

	<p class="word-break">({{.Binary}})<sub>2</sub> = ({{.Answer}})<sub>16</sub></p>`

	return util.ParseTemplate(tpl, data)
}

// DecimalToBinary outputs the proof and answer of a decimal to binary conversion.
func DecimalToBinary(decimal int) (string, error) {
	return decimalToBinaryHexadecimal(decimal, 2)
}

// DecimalToHexadecimal outputs the proof and answer of a decimal to hexadecimal conversion.
func DecimalToHexadecimal(decimal int) (string, error) {
	return decimalToBinaryHexadecimal(decimal, 16)
}

func decimalToBinaryHexadecimal(decimal int, base int) (string, error) {
	decimalInt := decimal

	var remainders []int
	var remaindersHex []string
	var proof bytes.Buffer
	var answer bytes.Buffer

	for decimalInt != 0 {
		currentValue := decimalInt
		decimalInt /= base
		remainder := currentValue % base

		fmt.Fprintf(&proof, "%d &divide; %d = %d, Remainder: %d<br>",
			currentValue, base, decimalInt, remainder)
		remainders = append(remainders, remainder)
	}

	if base == 16 {
		for _, r := range remainders {
			remaindersHex = append(remaindersHex, fmt.Sprintf("%X", r))
		}

		for i := len(remaindersHex) - 1; i >= 0; i-- {
			fmt.Fprint(&answer, remaindersHex[i])
		}
	} else {
		for i := len(remainders) - 1; i >= 0; i-- {
			fmt.Fprint(&answer, strconv.Itoa(remainders[i]))
		}
	}
	data := struct {
		Decimal       int
		Base          int
		Proof         string
		Remainders    []int
		RemaindersHex []string
		Answer        string
	}{
		decimal,
		base,
		proof.String(),
		remainders,
		remaindersHex,
		answer.String(),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>What is {{.Decimal}} in {{if eq .Base 2}}binary{{else}}hexadecimal{{end}}?</p>

	<h6>Answer:</h6>

	<p>The decimal number <b>{{.Decimal}}</b> in {{if eq .Base 2}}binary{{else}}hexadecimal{{end}} is <b>{{.Answer}}</b>.
	</p>

	<h6>Here's how to convert it:</h6>

	<p>Divide the decimal number by {{.Base}} and record the remainder:<br>
			{{.Proof}}</p>

	<p>Join the remainders from last to first:<br>
			<span class="word-break">{{.Answer}}</span></p>
	{{if .RemaindersHex -}}
	<p>Convert remainders to their hexadecimal equivalents:<br>
			Refer to the Conversion Table <a href="/category/networking/conversion-table" rel="noopener noreferrer"
																			target="_blank">here</a>.<br>
			{{range $i, $r := .Remainders -}}
			({{$r}})<sub>10</sub> = ({{index $.RemaindersHex $i}})<sub>16</sub><br>
			{{- end}}</p>
	{{- end}}

	<h6>Therefore:</h6>

	<p class="word-break">({{.Decimal}})<sub>10</sub> = (<span class="word-break">{{.Answer}}</span>)<sub>{{.Base}}</sub>
	</p>`

	return util.ParseTemplate(tpl, data)
}
