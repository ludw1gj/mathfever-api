package mathematics

// ChangeByPercentage outputs the proof and answer of calculating the change
// of a number by percentage.
func ChangeByPercentage(n float64, p float64) (string, error) {
	/*
		ChangeByPercentage:
		Increase/Decrease a number by a percentage.
		(n / 100) * p + n

		First work out 1% of 250, 250 ÷ 100 = 2.5 then multiply the answer by 23,
		because there was a 23% increase. 2.5 × 23 = 57.5.
	*/
	onePercent := n / 100
	addValue := onePercent * p
	answer := addValue + n

	var noun string
	if p < 0 {
		noun = "decrease"
	} else {
		noun = "increase"
	}

	data := struct {
		Number     float64
		Percentage float64
		OnePercent float64
		AddValue   float64
		Noun       string
		Answer     float64
	}{
		round(n, .5, 2),
		round(p, .5, 2),
		round(onePercent, .5, 2),
		round(addValue, .5, 2),
		noun,
		round(answer, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>
	<p>Where the {{.Noun}} in percentage is {{.Percentage}}% and the number is {{.Number}}, find the 
		resulting number.</p>

	<h6>Answer:</h6>
	<p>The number {{.Number}} {{.Noun}}d by {{.Percentage}}% is {{.Answer}}.</p>

	<h6>Here's how to calculate it:</h6>
	<p>Calculate using the equation shown below:<br>
		(Number &#xf7; 100) &times Percentage + Number<br>
		({{.Number}} &#xf7; 100) &times {{.Percentage}} + {{.Number}}</p>
	<p>Find 1% of {{.Number}}:<br>
		Number &#xf7; 100<br>
		{{.Number}} &#xf7; 100<br>
		{{.OnePercent}}</p>
	<p>Multiply the one percent value by {{.Percentage}} to find {{.Percentage}}% of {{.Number}}:<br>
		{{.OnePercent}} &times Percentage<br>
		{{.OnePercent}} &times {{.Percentage}}<br>
		{{.AddValue}}</p>
	<p>The value {{.AddValue}} is {{.Percentage}}% of {{.Number}}, add it to the Number to get the 
		final answer:<br>
		{{.AddValue}} + Number<br>
		{{.AddValue}} + {{.Number}}<br>
		{{.Answer}}</p>

	<h6>Therefore:</h6>
	<p>{{.Number}} {{.Noun}}d by {{.Percentage}}% is {{.Answer}}.</p>`

	return parseTemplate(tpl, data)
}

// NumberFromPercentage outputs the proof and answer of calculating finding the
// result of a percentage corresponding to a number.
func NumberFromPercentage(p float64, n float64) (string, error) {
	/*
		Find the value of the percentage of a number.
		(p / 100) * n
	*/
	divided := p / 100
	answer := divided * n

	data := struct {
		Percentage float64
		Number     float64
		Divided    float64
		Answer     float64
	}{
		round(p, .5, 2),
		round(n, .5, 2),
		round(divided, .5, 2),
		round(answer, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>
	<p>Where the percentage is {{.Percentage}}% and the number is {{.Number}}, find the resulting 
		number of the percentage.</p>

	<h6>Answer:</h6>
	<p>{{.Percentage}}% of {{.Number}} is {{.Answer}}.</p>

	<h6>Here's how to calculate it:</h6>
	<p>Use the equation below:<br>
		(Percentage &#xf7; 100) &times Number<br>
		({{.Percentage}} &#xf7; 100) &times {{.Number}}</p>
	<p>Find the decimalised version of the Percentage:<br>
		Percentage &#xf7; 100<br>
		{{.Percentage}} &#xf7; 100<br>
		{{.Divided}}</p>
	<p>Now multiply it with the Number to get the final answer:<br>
		({{.Divided}}) &times Number<br>
		({{.Divided}}) &times {{.Number}}<br>
		{{.Answer}}</p>

	<h6>Therefore:</h6>
	<p>{{.Percentage}}% of {{.Number}} is {{.Answer}}.</p>`

	return parseTemplate(tpl, data)
}

// PercentageChange outputs the proof and answer of calculating the percentage
// change from one number to another.
func PercentageChange(n float64, newN float64) (string, error) {
	/*
		Find the percentage change from one number to another.
		(new_n - original_n) / original_n * 100
	*/
	change := newN - n
	decimalisedPercentage := change / n
	answer := decimalisedPercentage * 100

	data := struct {
		Number                float64
		NewNumber             float64
		Change                float64
		DecimalisedPercentage float64
		Answer                float64
	}{
		round(n, .5, 2),
		round(newN, .5, 2),
		round(change, .5, 2),
		round(decimalisedPercentage, .5, 2),
		round(answer, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>
	<p>Where original number is {{.Number}} and the new number is {{.NewNumber}}, find the difference 
		between the numbers in percentage.</p>

	<h6>Answer:</h6>
	<p>The difference in percentage between {{.Number}} and {{.NewNumber}} is {{.Answer}}%.</p>

	<h6>Here's how to calculate it:</h6>
	<p>Calculate using the equation shown below:<br>
		((New Number - Original Number) &#xf7; Original Number) &times 100<br>
	<p>Find the change in value between the two numbers:<br>
		New Number - Original Number<br>
		{{.NewNumber}} - {{.Number}}<br>
		{{.Change}}</p>
	<p>Find the decimalised version of the percentage change using the change in value:<br>
		{{.Change}} &#xf7; Original Number<br>
		{{.Change}} &#xf7; {{.Number}}<br>
		{{.DecimalisedPercentage}}</p>
	<p>Convert decimalised percentage to percentage to get the final answer:<br>
		{{.DecimalisedPercentage}} &times 100<br>
		{{.Answer}}</p>

	<h6>Therefore:</h6>
	<p>The percentage of change from {{.Number}} to {{.NewNumber}} is {{.Answer}}%.</p>`

	return parseTemplate(tpl, data)
}

// PercentageFromNumber outputs the proof and answer of calculating the
// percentage of a number of a total number.
func PercentageFromNumber(n float64, totalN float64) (string, error) {
	/*
		Find the percentage of a number of a total number.
		(n / total_n) * 100
	*/
	decimalisedPercentage := n / totalN
	answer := decimalisedPercentage * 100

	data := struct {
		Number                float64
		TotalNumber           float64
		DecimalisedPercentage float64
		Answer                float64
	}{
		round(n, .5, 2),
		round(totalN, .5, 2),
		round(decimalisedPercentage, .5, 2),
		round(answer, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>
	<p>Where the number is {{.Number}} and the total number is {{.TotalNumber}}, find how much 
		percentage {{.Number}} is of {{.TotalNumber}}.</p>

	<h6>Answer:</h6>
	<p>{{.Number}} is {{.Answer}}% of {{.TotalNumber}}.</p>

	<h6>Here's how to calculate it:</h6>
	<p>Calculate using the equation shown below:<br>
		(Number &#xf7; Total Number) &times 100</p>
	<p>Find the decimalised version of the percentage:<br>
		Number &#xf7; Total Number<br>
		{{.Number}} &#xf7; {{.TotalNumber}}<br>
		{{.DecimalisedPercentage}}</p>
	<p>Convert decimalised percentage to percentage to get the answer:<br>
		{{.DecimalisedPercentage}} &times 100<br>
		{{.Answer}}</p>

	<h6>Therefore:</h6>
	<p>{{.Number}} is {{.Answer}}% of {{.TotalNumber}}.</p>`

	return parseTemplate(tpl, data)
}
