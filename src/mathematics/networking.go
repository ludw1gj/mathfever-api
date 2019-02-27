package mathematics

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// BinaryToDecimal outputs the proof and answer of a binary to decimal conversion.
func BinaryToDecimal(binary string) (string, error) {
	const base = 2
	lenBinary := len(binary)

	power := lenBinary - 1
	var proofStepsBuf [4]strings.Builder
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
	for i, buf := range proofStepsBuf {
		proofSteps[i] = strings.TrimSuffix(buf.String(), " + ")
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
	<h4>Question:</h4>
	<p>What is {{.Binary}} in decimal?</p>

	<h4>Answer:</h4>
	<p>The binary number <b>{{.Binary}}</b> in decimal is <b>{{.Answer}}</b>.</p>

	<h4>Here's how to convert it:</h4>
	<p>Using the following continuing equation the accurate answer can be found: 
		(b x 2<sup>n</sup>)...<br>
		Where <strong>b</strong> is the binary number at the n<sup>th</sup> position and 
		<strong>n</strong> is the n<sup>th</sup> position.<br>
		There are {{.LenBinary}} numbers in this binary number, the n<sup>th</sup> position starts from 
		0, and therefore the n<sup>th</sup> position goes up to {{.MaxPosition}} ({{.LenBinary}} - 1). 
		Count down from {{.MaxPosition}} and insert the binary number at that position.</p>
	<p>Using the equation for the binary number <span class="word-break">{{.Binary}}</span>, it would 
		go like this:<br>
		{{range .Proof}}{{.}}<br>{{end}}
		<span class="word-break">{{.Answer}}</span></p>

	<h4>Therefore:</h4>
	<p class="word-break">({{.Binary}})<sub>2</sub> = ({{.Answer}})<sub>10</sub></p>`

	return parseTemplate(tpl, data)
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

	var proof strings.Builder
	var answer strings.Builder
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
	<h4>Question:</h4>
	<p>What is {{.Binary}} in hexadecimal?</p>

	<h4>Answer:</h4>
	<p>The binary number <b>{{.Binary}}</b> in hexadecimal is <b>{{.Answer}}</b>.</p>

	<h4>Here's how to convert it:</h4>
	<p>Break the binary number <span class="word-break">{{.Binary}}</span> into groups of 4 digits as 
		such:<br>
		{{range .GroupedBinary}}{{.}}{{end}}</p>
	<p>Convert those groups (base 2) to a decimal number (base 10), that way it is easy to convert
		them to a Hexadecimal number (base 16):<br>
		Refer to the Conversion Table <a href="/category/networking/conversion-table" 
		rel="noopener noreferrer" target="_blank">here</a>.<br>
		{{.Proof}}</p>

	<h4>Therefore:</h4>
	<p class="word-break">({{.Binary}})<sub>2</sub> = ({{.Answer}})<sub>16</sub></p>`

	return parseTemplate(tpl, data)
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
	var proof strings.Builder
	var answer strings.Builder

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
	<h4>Question:</h4>
	<p>What is {{.Decimal}} in {{if eq .Base 2}}binary{{else}}hexadecimal{{end}}?</p>

	<h4>Answer:</h4>
	<p>The decimal number <b>{{.Decimal}}</b> in {{if eq .Base 2}}binary{{else}}hexadecimal{{end}} is 
		<b>{{.Answer}}</b>.</p>

	<h4>Here's how to convert it:</h4>
	<p>Divide the decimal number by {{.Base}} and record the remainder:<br>
		{{.Proof}}</p>
	<p>Join the remainders from last to first:<br>
		<span class="word-break">{{.Answer}}</span></p>
	{{if .RemaindersHex -}}
	<p>Convert remainders to their hexadecimal equivalents:<br>
		Refer to the Conversion Table <a href="/category/networking/conversion-table" 
		rel="noopener noreferrer" target="_blank">here</a>.<br>
		{{range $i, $r := .Remainders -}}
		({{$r}})<sub>10</sub> = ({{index $.RemaindersHex $i}})<sub>16</sub><br>
		{{- end}}</p>
	{{- end}}

	<h4>Therefore:</h4>
	<p class="word-break">({{.Decimal}})<sub>10</sub> = (<span class="word-break">{{.Answer}}</span>)
		<sub>{{.Base}}</sub></p>`

	return parseTemplate(tpl, data)
}

// HexadecimalToBinary outputs the proof and answer of a hexadecimal to binary conversion.
func HexadecimalToBinary(hexadecimal string) (string, error) {
	var binaries strings.Builder
	var proof strings.Builder

	var buf strings.Builder
	for _, char := range hexadecimal {
		decimalChar, err := strconv.ParseInt(string(char), 16, 0)
		if err != nil {
			// a non hexadecimal character was passed in
			log.Println(err)
			return "", err
		}

		binary := fmt.Sprintf("%b", decimalChar)
		if (len(binary) % 4) == 0 {
			fmt.Fprint(&binaries, binary)
		} else {
			n := 4 - (len(binary) % 4)
			for i := 0; i < n; i++ {
				buf.WriteString("0")
			}
			fmt.Fprint(&binaries, buf.String(), binary)
			buf.Reset()
		}
		fmt.Fprintf(&proof, "(%s)<sub>16</sub> = (%s)<sub>2</sub><br>", string(char), binary)
	}
	// remove last character - binary variable overshoots by one in forloop
	answer := binaries.String()[:binaries.Len()-1]

	data := struct {
		Hexadecimal string
		Proof       string
		Answer      string
	}{
		hexadecimal,
		proof.String(),
		answer,
	}

	const tpl = `
	<h4>Question:</h4>
	<p>What is {{.Hexadecimal}} in binary?</p>

	<h4>Answer:</h4>
	<p>The hexadecimal number <b>{{.Hexadecimal}}</b> in binary is <b>{{.Answer}}</b>.</p>

	<h4>Here's how to convert it:</h4>
	<p>Convert each hexadecimal digit manually to its binary equivalent:<br>
		Refer to the Conversion Table <a href="/category/networking/conversion-table"
		rel="noopener noreferrer" target="_blank">here</a>.<br>
		{{.Proof}}</p>
	<p>Join the digit equivalents together from first to last:<br>
		{{.Answer}}</p>

	<h4>Therefore:</h4>
	<p>({{.Hexadecimal}})<sub>16</sub> = ({{.Answer}})<sub>2</sub></p>`

	return parseTemplate(tpl, data)
}

// HexadecimalToDecimal outputs the proof and answer of a hexadecimal to decimal conversion.
func HexadecimalToDecimal(hexadecimal string) (string, error) {
	hexLength := len(hexadecimal) - 1

	var decimals []int64
	var proof1 strings.Builder
	var proof2Buf [4]strings.Builder
	var answer int64

	for _, char := range hexadecimal {
		decimal, err := strconv.ParseInt(string(char), 16, 0)
		if err != nil {
			log.Println(err)
			return "", err
		}
		decimals = append(decimals, decimal)
		power := int64(math.Pow(16, float64(hexLength)))
		multiplied := decimal * power
		answer += multiplied

		fmt.Fprintf(&proof1, "(%s)<sub>16</sub> = (%d)<sub>10</sub><br>", string(char), decimal)

		fmt.Fprintf(&proof2Buf[0], "(%d x 16<sup>%d</sup>) + ", decimal, hexLength)
		fmt.Fprintf(&proof2Buf[1], "(%d x %d) + ", decimal, power)
		fmt.Fprintf(&proof2Buf[2], "%d + ", decimal*power)
		hexLength--
	}
	fmt.Fprintf(&proof2Buf[3], "%d + ", answer)

	// add "${hexadecimal} = " to front, and remove " + " from the end
	var proof2 [4]string
	for i, line := range proof2Buf {
		proof2[i] = fmt.Sprintf("%s = %s<br>", hexadecimal, strings.TrimSuffix(line.String(), " + "))
	}

	data := struct {
		Hexadecimal string
		Proof1      string
		Proof2      [4]string
		Answer      int64
	}{
		hexadecimal,
		proof1.String(),
		proof2,
		answer,
	}

	const tpl = `
	<h4>Question:</h4>
	<p>What is {{.Hexadecimal}} in decimal?</p>

	<h4>Answer:</h4>
	<p>The hexadecimal number <b>{{.Hexadecimal}}</b> in decimal is <b>{{.Answer}}</b>.</p>

	<h4>Here's how to convert it:</h4>
	<p>Get the decimal equivalent of each hexadecimal digit:<br>
		Refer to the Conversion Table <a href="/category/networking/conversion-table"
		rel="noopener noreferrer" target="_blank">here</a>.<br>
		{{.Proof1}}</p>
	<p>Multiply each hexadecimal digit with 16 power of digit position. Digit position starts from 
		zero, right to left:<br>
		{{range .Proof2}}{{.}}{{end}}</p>

	<h4>Therefore:</h4>
	<p class="word-break">({{.Hexadecimal}})<sub>16</sub> = ({{.Answer}})<sub>10</sub></p>`

	return parseTemplate(tpl, data)
}
