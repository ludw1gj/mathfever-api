package mathematics

import (
	"math"
)

// TsaPythagoreanTheorem outputs the proof and answer of finding the hypotenuse
// using the pythagorean theorem.
func TsaPythagoreanTheorem(a float64, b float64) (string, error) {
	/*
		Use the Pythagorean Theorem to find side lengths (a or b), or the
		hypotenuse (c) of a right-angle triangle.

		Pythagorean Theorem: right-angle triangle with side lengths (a, b) and hypotenuse (c).
		a^2 + b^2 = c^2

		Example:
		When a = 25 and b = 17, find c.
		c^2 	= 25^2 + 17^2
			= 625 + 289
			= 914
		c	= √914
			≈ 30.2
	*/
	aSqr := math.Pow(a, 2)
	bSqr := math.Pow(b, 2)
	ab := aSqr + bSqr
	answer := math.Sqrt(ab)

	data := struct {
		A      float64
		B      float64
		ASqr   float64
		BSqr   float64
		AB     float64
		Answer float64
	}{
		round(a, .5, 2),
		round(b, .5, 2),
		round(aSqr, .5, 2),
		round(bSqr, .5, 2),
		round(ab, .5, 2),
		round(answer, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>Where a = {{.A}} and b = {{.B}}, find c (hypotenuse).</p>

	<h6>Answer:</h6>

	<p>c is equal to {{.Answer}}.</p>

	<h6>Here's how to calculate it:</h6>

	<p>The Pythagorean theorem, also known as Pythagoras's theorem, can be used when two sides of a right angled
			triangle are known, and then we can find the third by using the equation:<br>
			a<sup>2</sup> + b<sup>2</sup> = c<sup>2</sup></p>

	<p>Where a = {{.A}} and b = {{.B}}, find c (hypotenuse):<br>
			{{.A}}<sup>2</sup> + {{.B}}<sup>2</sup> = c<sup>2</sup></p>

	<p>Square a and b:<br>
			{{.ASqr}} + {{.BSqr}} = c<sup>2</sup></p>

	<p>Add a and b together:<br>
			{{.AB}} = c<sup>2</sup></p>

	<p>Square root both sides:<br>
			&radic;{{.AB}} = &radic;c<sup>2</sup><br>
			c = {{.Answer}}</p>

	<h6>Therefore:</h6>

	<p>c is equal to {{.Answer}}.</p>`

	return parseTemplate(tpl, data)
}

// TsaCone outputs the proof and answer of finding the total surface area of a cone.
func TsaCone(radius float64, slantHeight float64) (string, error) {
	/*
		Find the Total Surface Area of a cone, with the equation:
		Where S is Slant Height, r is radius
		TSA     = tsa of base (circle) + tsa of curved surface
			= πr^2 + πrS
			= πr(r + S)
	*/
	baseArea := math.Pi * math.Pow(radius, 2)
	curvedSurfaceArea := math.Pi * radius * slantHeight
	tsa := baseArea + curvedSurfaceArea

	data := struct {
		Radius            float64
		SlantHeight       float64
		BaseArea          float64
		CurvedSurfaceArea float64
		TSA               float64
	}{
		round(radius, .5, 2),
		round(slantHeight, .5, 2),
		round(baseArea, .5, 2),
		round(curvedSurfaceArea, .5, 2),
		round(tsa, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>Where r (radius) is {{.Radius}}, and S (slant height) is {{.SlantHeight}}, find the total surface area of a cone.</p>
	
	<h6>Answer:</h6>
	
	<p>The total surface area of a cone with the radius of {{.Radius}} and a slant height of {{.SlantHeight}}
			is {{.TSA}}.</p>
	
	
	<h6>Here's how to calculate it:</h6>
	
	<p>Find the Total Surface Area of a cone, with the equation:<br>
			TSA = tsa of the base (circle) + tsa of the curved surface<br>
			TSA = π r<sup>2</sup> + π rS<br>
	
	<p>r is the radius and S is the Slant Height.</p>
	
	<p>Where r = {{.Radius}} and S = {{.SlantHeight}}:<br>
			Find the tsa of the base:<br>
			π r<sup>2</sup><br>
			π &times {{.Radius}}<sup>2</sup><br>
			{{.BaseArea}}</p>
	
	<p>Find the tsa of the curved surface:<br>
			π rS<br>
			π &times {{.Radius}} &times {{.SlantHeight}}<br>
			{{.CurvedSurfaceArea}}</p>
	
	<p>Add the two surface areas together:<br>
			{{.BaseArea}} + {{.CurvedSurfaceArea}}<br>
			{{.TSA}}</p>
	
	<h6>Therefore:</h6>
	
	<p>The total surface area of a cone with the radius of {{.Radius}} and a slant height of {{.SlantHeight}}
			equals {{.TSA}}.</p>`

	return parseTemplate(tpl, data)
}

// TsaCube outputs the proof and answer of finding the total surface area of a cube.
func TsaCube(length float64) (string, error) {
	/*
		Find the Total Surface Area of a cube, with the equation:
		TSA = 6L^2
	*/
	lengthSqr := math.Pow(length, 2)
	tsa := 6 * lengthSqr

	data := struct {
		Length    float64
		LengthSqr float64
		TSA       float64
	}{
		round(length, .5, 2),
		round(lengthSqr, .5, 2),
		round(tsa, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>Where the length is {{.Length}}, find the total surface area of a cube.</p>
	
	<h6>Answer:</h6>
	
	<p>The total surface area of a cube with the length of {{.Length}} is {{.TSA}}.</p>
	
	<h6>Here's how to calculate it:</h6>
	
	<p>Find the Total Surface Area of a cube, with the equation:<br>
			TSA = 6L<sup>2</sup></p>
	
	<p>L is the length of each side of the cube. The surface tsa of a single side is a square, thus L &times L, which
			is the same as L<sup>2</sup>. Since all of the sides are the same, and there are 6 sides, you times L<sup>2</sup> by
			6.</p>
	
	<p>Where L = {{.Length}}:<br>
			TSA = 6 &times {{.Length}}<sup>2</sup><br>
			TSA = 6 &times {{.LengthSqr}}<br>
			TSA = {{.TSA}}</p>
	
	<h6>Therefore:</h6>
	
	<p>The total surface area of a cube with the length of {{.Length}} equals {{.TSA}}.</p>`

	return parseTemplate(tpl, data)
}

// TsaCylinder outputs the proof and answer of finding the total surface area of a cylinder.
func TsaCylinder(radius float64, height float64) (string, error) {
	/*
		Find the Total Surface Area of a cylinder, with the equation:
		TSA = tsa of 2 circles + curved surface
		    = 2πr^2 + 2πrh
		    = 2πr(r + h)
	*/
	radiusSqr := math.Pow(radius, 2)
	circlesArea := 2 * math.Pi * radiusSqr
	rh := radius * height
	curvedSurfaceArea := 2 * math.Pi * rh
	tsa := circlesArea + curvedSurfaceArea

	data := struct {
		Radius            float64
		Height            float64
		RadiusSqr         float64
		CirclesArea       float64
		RH                float64
		CurvedSurfaceArea float64
		TSA               float64
	}{
		round(radius, .5, 2),
		round(height, .5, 2),
		round(radiusSqr, .5, 2),
		round(circlesArea, .5, 2),
		round(rh, .5, 2),
		round(curvedSurfaceArea, .5, 2),
		round(tsa, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>Where the radius is {{.Radius}} and the height is {{.Height}}, find the total surface area of a cylinder.</p>
	
	<h6>Answer:</h6>
	
	<p>The total surface area of a cylinder with the radius of {{.Radius}}, and a height of {{.Height}} is {{.TSA}}.</p>
	
	<h6>Here's how to calculate it:</h6>
	
	<p>Find the Total Surface Area of a cylinder, with the equation:<br>
			TSA = tsa of 2 circles + tsa of curved surface<br>
			TSA = 2 π r<sup>2</sup> + 2 π rh</p>
	
	<p>r is the radius and h is the height. A cylinder consists of two circles and a curved surface. Find the tsa of one
			of the circles and multiply that by 2, find the the tsa of the curved surface, and then add those two together.</p>
	
	<p>Where r = {{.Radius}}, and h = {{.Height}}:<br>
			Find the tsa of the two circles:<br>
			2 π r<sup>2</sup><br>
			2 &times π &times {{.Radius}}<sup>2</sup><br>
			2 &times π &times {{.RadiusSqr}}<br>
			{{.CirclesArea}}</p>
	
	<p>Find the tsa of curved surface:<br>
			2 π rh<br>
			2 &times π &times {{.Radius}} &times {{.Height}}<br>
			2 &times π &times {{.RH}}<br>
			{{.CurvedSurfaceArea}}</p>
	
	<p>Add the two surface areas together:<br>
			{{.CirclesArea}} + {{.CurvedSurfaceArea}}<br>
			{{.TSA}}</p>
	
	<h6>Therefore:</h6>
	
	<p>The total surface area of a cylinder with the radius of {{.Radius}}, and a height of {{.Height}} equals {{.TSA}}.</p>`

	return parseTemplate(tpl, data)
}

// TsaRectangularPrism outputs the proof and answer of finding the total
// surface area of a rectangular prism.
func TsaRectangularPrism(height float64, length float64, width float64) (string, error) {
	/*
		Find the Total Surface Area of a rectangular prism, with the equation:
		TSA = 2(wh + lw + lh)
	*/
	add := (width * height) + (length * width) + (length * height)
	tsa := 2 * add

	data := struct {
		Height float64
		Length float64
		Width  float64
		Add    float64
		TSA    float64
	}{
		round(height, .5, 2),
		round(length, .5, 2),
		round(width, .5, 2),
		round(add, .5, 2),
		round(tsa, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>Where the height is {{.Height}}, the length is {{.Length}}, and the width is {{.Width}}, find the total surface area
			of a rectangular prism.</p>
	
	<h6>Answer:</h6>
	
	<p>The total surface area of a rectangular prism with the height of {{.Height}}, a length of {{.Length}}, and a
			width of {{.Width}} is {{.TSA}}.</p>
	
	<h6>Here's how to calculate it:</h6>
	
	<p>Find the Total Surface Area of a rectangular prism, with the equation:<br>
			TSA = 2(WH + LW + LH)</p>
	
	<p>A rectangular prism has a surface tsa of six rectangles. Two sides of each: width &times height, length &times
			width, and length &times height. That would be 2WH + 2LW + 2LH, which would be the same as 2(WH + LW + LH).</p>
	
	<p>Where H = 2, L = 4, W = 3:<br>
			TSA = 2({{.Width}} &times {{.Height}} + {{.Length}} &times {{.Width}} + {{.Length}} &times {{.Height}})<br>
			TSA = 2({{.Add}})<br>
			TSA = {{.TSA}}</p>
	
	<h6>Therefore:</h6>
	
	<p>The total surface area of a rectangular prism with the height of {{.Height}}, a length of {{.Length}}, and a
			width of {{.Width}} is equal to {{.TSA}}.</p>`

	return parseTemplate(tpl, data)
}

// TsaSphere outputs the proof and answer of finding the total surface area of a sphere.
func TsaSphere(radius float64) (string, error) {
	/*
		Find the Total Surface Area of a sphere, with the equation:
		TSA = 4πr^2
	*/
	tsa := 4 * math.Pi * math.Pow(radius, 2)

	data := struct {
		Radius float64
		TSA    float64
	}{
		round(radius, .5, 2),
		round(tsa, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>Where the r (radius) is {{.Radius}}, find the total surface area of a sphere.</p>

	<h6>Answer:</h6>

	<p>The total surface area of a sphere with the radius of {{.Radius}} equals {{.TSA}}.</p>

	<h6>Here's how to calculate it:</h6>

	<p>Find the Total Surface Area of a sphere, with the equation:<br>
			TSA = 4 &#928; r<sup>2</sup></p>

	<p>r is the radius of the sphere.</p>

	<p>Where r = {{.Radius}}:<br>
			TSA = 4 &times &#928; &times {{.Radius}}<sup>2</sup><br>
			TSA = {{.TSA}}</p>

	<h6>Therefore:</h6>

	<p>The total surface area of a sphere with the radius of {{.Radius}} equals {{.TSA}}.</p>`

	return parseTemplate(tpl, data)
}

// TsaSquareBaseTriangle outputs the proof and answer of finding the total
// surface area of a square base triangle.
func TsaSquareBaseTriangle(baseLength float64, height float64) (string, error) {
	/*
		Find the Total Surface Area of a square based triangle, with the equation:
		TSA = tsa of square + tsa of 4 triangles
		    = b^2 + 4 * (1/2bh)
		    = b^2 + 2bh
	*/
	squareArea := math.Pow(baseLength, 2)
	fourTrianglesArea := 2 * (baseLength * height)
	tsa := squareArea + fourTrianglesArea

	data := struct {
		BaseLength        float64
		Height            float64
		SquareArea        float64
		FourTrianglesArea float64
		TSA               float64
	}{
		round(baseLength, .5, 2),
		round(height, .5, 2),
		round(squareArea, .5, 2),
		round(fourTrianglesArea, .5, 2),
		round(tsa, .5, 2),
	}

	const tpl = `
	<h6>Question:</h6>

	<p>Where the base length is {{.BaseLength}}, and the height is {{.Height}}, find the total surface area of a square
			based
			triangle.</p>

	<h6>Answer:</h6>

	<p>The total surface area of a square based triangle with the base length of {{.BaseLength}}, and a height is
			{{.Height}} is {{.TSA}}.</p>

	<h6>Here's how to calculate it:</h6>

	<p>Find the Total Surface Area of a square based triangle, with the equation:<br>
			TSA = tsa of the square + tsa of 4 triangles<br>
			TSA = b<sup>2</sup> + 4 &times (1/2bh)<br>
			TSA = b<sup>2</sup> + 2bh</p>

	<p>b is the base length, and h is the height.</p>

	<p>Where b = {{.BaseLength}}, and h = {{.Height}}:<br>
			Area of the square:<br>
			b<sup>2</sup><br>
			{{.BaseLength}}<sup>2</sup><br>
			{{.SquareArea}}</p>

	<p>Area of 4 triangles:<br>
			2bh<br>
			2 &times {{.BaseLength}} &times {{.Height}}<br>
			{{.FourTrianglesArea}}</p>

	<p>Add the two surface areas together:<br>
			{{.SquareArea}} + {{.FourTrianglesArea}}<br>
			{{.TSA}}</p>

	<h6>Therefore:</h6>

	<p>The total surface area of a square based triangle with the base length of {{.BaseLength}}, and a height is
			{{.Height}} equals {{.TSA}}.</p>`

	return parseTemplate(tpl, data)
}
