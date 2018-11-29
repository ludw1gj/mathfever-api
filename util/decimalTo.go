package util

import (
	"bytes"
	"fmt"
	"strconv"
)

func DecimalToBinaryHexadecimal(decimal int, base int) (string, error) {
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

	return ParseTemplate(tpl, data)
}
