package logic

import (
	"errors"
)

const categoriesJson = `[
  {
	"name": "Total Surface Area",
	"slug": "total-surface-area",
	"description": "Calculations that you might use for Total Surface Area: Pythagorean Theorem (also known as Pythagoras's Theorem), Total Surface Area of Cone, Total Surface Area of Cube, Total Surface Area of Cylinder, Total Surface Area of Rectangular Prism, Total Surface Area of Sphere, and Total Surface Area of Square Based Triangle.",
	"calculations": [
	  {
		"name": "Pythagorean Theorem",
		"slug": "tsa-pythagorean-theorem",
		"description": "Solve: Pythagorean Theorem, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Side A", "tag": "side_a" },
		  { "name": "Side B", "tag": "side_b" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where a = 25 and b = 17, find c (hypotenuse).</p>\n\n<h4>Answer:</h4>\n\n<p>c is equal to 30.23.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>The Pythagorean theorem, also known as Pythagoras's theorem, can be used when two sides of a right angled\n    triangle are known, and then we can find the third by using the equation:<br>\n    a<sup>2</sup> + b<sup>2</sup> = c<sup>2</sup></p>\n\n<p>Where a = 25 and b = 17, find c (hypotenuse):<br>\n    25<sup>2</sup> + 17<sup>2</sup> = c<sup>2</sup></p>\n\n<p>Square a and b:<br>\n    625 + 289 = c<sup>2</sup></p>\n\n<p>Add a and b together:<br>\n    914 = c<sup>2</sup></p>\n\n<p>Square root both sides:<br>\n    √914 = √c<sup>2</sup><br>\n    c = 30.23</p>\n\n<h4>Therefore:</h4>\n\n<p>c is equal to 30.23.</p>\n",
		"category": "Total Surface Area",
		"categorySlug": "total-surface-area"
	  },
	  {
		"name": "Cone",
		"slug": "tsa-cone",
		"description": "Solve: Total Surface Area of Cone, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Radius", "tag": "radius" },
		  { "name": "Slant Height", "tag": "slant_height" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where r (radius) is 3, and S (slant height) is 5, find the total surface area of a cone.</p>\n\n<h4>Answer:</h4>\n\n<p>The total surface area of a cone with the radius of 3 and a slant height of 5\n    is 75.4.</p>\n\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Find the Total Surface Area of a cone, with the equation:<br>\n    TSA = tsa of the base (circle) + tsa of the curved surface<br>\n    TSA = π r<sup>2</sup> + π rS<br>\n\n<p>r is the radius and S is the Slant Height.</p>\n\n<p>Where r = 3 and S = 5:<br>\n    Find the tsa of the base:<br>\n    π r<sup>2</sup><br>\n    π &times 3<sup>2</sup><br>\n    28.27</p>\n\n<p>Find the tsa of the curved surface:<br>\n    π rS<br>\n    π &times 3 &times 5<br>\n    47.12</p>\n\n<p>Add the two surface areas together:<br>\n    28.27 + 47.12<br>\n    75.4</p>\n\n<h4>Therefore:</h4>\n\n<p>The total surface area of a cone with the radius of 3 and a slant height of 5\n    equals 75.4.</p>",
		"category": "Total Surface Area",
		"categorySlug": "total-surface-area"
	  },
	  {
		"name": "Cube",
		"slug": "tsa-cube",
		"description": "Solve: Total Surface Area of Cube, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Length", "tag": "length" }],
		"example": "<h4>Question:</h4>\n\n<p>Where the length is 3, find the total surface area of a cube.</p>\n\n<h4>Answer:</h4>\n\n<p>The total surface area of a cube with the length of 3 is 54.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Find the Total Surface Area of a cube, with the equation:<br>\n    TSA = 6L<sup>2</sup></p>\n\n<p>L is the length of each side of the cube. The surface tsa of a single side is a square, thus L &times L, which\n    is the same as L<sup>2</sup>. Since all of the sides are the same, and there are 6 sides, you times L<sup>2</sup> by\n    6.</p>\n\n<p>Where L = 3:<br>\n    TSA = 6 &times 3<sup>2</sup><br>\n    TSA = 6 &times 9<br>\n    TSA = 54</p>\n\n<h4>Therefore:</h4>\n\n<p>The total surface area of a cube with the length of 3 equals 54.</p>",
		"category": "Total Surface Area",
		"categorySlug": "total-surface-area"
	  },
	  {
		"name": "Cylinder",
		"slug": "tsa-cylinder",
		"description": "Solve: Total Surface Area of Cylinder, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Radius", "tag": "radius" },
		  { "name": "Height", "tag": "height" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where the radius is 2 and the height is 5, find the total surface area of a cylinder.</p>\n\n<h4>Answer:</h4>\n\n<p>The total surface area of a cylinder with the radius of 2, and a height of 5 is 87.96.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Find the Total Surface Area of a cylinder, with the equation:<br>\n    TSA = tsa of 2 circles + tsa of curved surface<br>\n    TSA = 2 π r<sup>2</sup> + 2 π rh</p>\n\n<p>r is the radius and h is the height. A cylinder consists of two circles and a curved surface. Find the tsa of one\n    of the circles and multiply that by 2, find the the tsa of the curved surface, and then add those two together.</p>\n\n<p>Where r = 2, and h = 5:<br>\n    Find the tsa of the two circles:<br>\n    2 π r<sup>2</sup><br>\n    2 &times π &times 2<sup>2</sup><br>\n    2 &times π &times 4<br>\n    25.13</p>\n\n<p>Find the tsa of curved surface:<br>\n    2 π rh<br>\n    2 &times π &times 2 &times 5<br>\n    2 &times π &times 10<br>\n    62.83</p>\n\n<p>Add the two surface areas together:<br>\n    25.13 + 62.83<br>\n    87.96</p>\n\n<h4>Therefore:</h4>\n\n<p>The total surface area of a cylinder with the radius of 2, and a height of 5 equals 87.96.</p>",
		"category": "Total Surface Area",
		"categorySlug": "total-surface-area"
	  },
	  {
		"name": "Rectangular Prism",
		"slug": "tsa-rectangular-prism",
		"description": "Solve: Total Surface Area of Rectangular Prism, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Height", "tag": "height" },
		  { "name": "Length", "tag": "length" },
		  { "name": "Width", "tag": "width" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where the height is 2, the length is 4, and the width is 3, find the total surface area\n    of a rectangular prism.</p>\n\n<h4>Answer:</h4>\n\n<p>The total surface area of a rectangular prism with the height of 2, a length of 4, and a\n    width of 3 is 52.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Find the Total Surface Area of a rectangular prism, with the equation:<br>\n    TSA = 2(WH + LW + LH)</p>\n\n<p>A rectangular prism has a surface tsa of six rectangles. Two sides of each: width &times height, length &times\n    width, and length &times height. That would be 2WH + 2LW + 2LH, which would be the same as 2(WH + LW + LH).</p>\n\n<p>Where H = 2, L = 4, W = 3:<br>\n    TSA = 2(3 × 2 + 4 × 3 + 4 × 2)<br>\n    TSA = 2(26)<br>\n    TSA = 52</p>\n\n<h4>Therefore:</h4>\n\n<p>The total surface area of a rectangular prism with the height of 2, a length of 4, and a\n    width of 3 is equal to 52.</p>",
		"category": "Total Surface Area",
		"categorySlug": "total-surface-area"
	  },
	  {
		"name": "Sphere",
		"slug": "tsa-sphere",
		"description": "Solve: Total Surface Area of Sphere, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Radius", "tag": "radius" }],
		"example": "<h4>Question:</h4>\n\n<p>Where the r (radius) is 1, find the total surface area of a sphere.</p>\n\n<h4>Answer:</h4>\n\n<p>The total surface area of a sphere with the radius of 1 equals 12.57.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Find the Total Surface Area of a sphere, with the equation:<br>\n    TSA = 4 π r<sup>2</sup></p>\n\n<p>r is the radius of the sphere.</p>\n\n<p>Where r = 1:<br>\n    TSA = 4 &times π &times 1<sup>2</sup><br>\n    TSA = 12.57</p>\n\n<h4>Therefore:</h4>\n\n<p>The total surface area of a sphere with the radius of 1 equals 12.57.</p>",
		"category": "Total Surface Area",
		"categorySlug": "total-surface-area"
	  },
	  {
		"name": "Square Based Triangle",
		"slug": "tsa-square-based-triangle",
		"description": "Solve: Total Surface Area of Square Based Triangle, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Base Length", "tag": "base_length" },
		  { "name": "Height", "tag": "height" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where the base length is 4, and the height is 6, find the total surface area of a square\n    based\n    triangle.</p>\n\n<h4>Answer:</h4>\n\n<p>The total surface area of a square based triangle with the base length of 4, and a height is\n    6 is 64.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Find the Total Surface Area of a square based triangle, with the equation:<br>\n    TSA = tsa of the square + tsa of 4 triangles<br>\n    TSA = b<sup>2</sup> + 4 &times (1/2bh)<br>\n    TSA = b<sup>2</sup> + 2bh</p>\n\n<p>b is the base length, and h is the height.</p>\n\n<p>Where b = 4, and h = 6:<br>\n    Area of the square:<br>\n    b<sup>2</sup><br>\n    4<sup>2</sup><br>\n    16</p>\n\n<p>Area of 4 triangles:<br>\n    2bh<br>\n    2 &times 4 &times 6<br>\n    48</p>\n\n<p>Add the two surface areas together:<br>\n    16 + 48<br>\n    64</p>\n\n<h4>Therefore:</h4>\n\n<p>The total surface area of a square based triangle with the base length of 4, and a height is\n    6 equals 64.</p>",
		"category": "Total Surface Area",
		"categorySlug": "total-surface-area"
	  }
	]
  },
  {
	"name": "Primes and Factors",
	"slug": "primes-and-factors",
	"description": "Calculations about numbers! Find Highest Common Factor, find Lowest Common Multiple, and figuring out Prime Numbers.",
	"calculations": [
	  {
		"name": "Find if Number is a Prime Number",
		"slug": "is-prime",
		"description": "Solve: Find if Number is a Prime Number, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Number", "tag": "number" }],
		"example": "<h4>Question:</h4>\n\n<p>Is 129 a prime number?</p>\n\n<h4>Answer:</h4>\n\n<p>129 is not a prime number, because it is divisible by 3.</p>\n\n<h4>Helpful Tips:</h4>\n\n<p>Prime numbers are numbers whose only whole-number factors are 1 and itself.</p>\n\n<p>To determine if a number (n) is a prime number, here are some rules:</p>\n\n<ul>\n    <li>Prime numbers are defined for whole-numbers greater than 1.</li>\n    <li>1 is not considered a prime number.</li>\n    <li>2 is considered a prime number because its only whole-number factors are 1 and itself.</li>\n    <li>Any number that is divisible by 2 is not a prime number.</li>\n    <li>After the steps above, start by dividing the number by 3, then 5, then 7, and so on, checking if the number\n        can be divided by those divisors. As we have determined it cannot be divided by 2, there is no need to divide by\n        4, 6, 8, and so on.\n    </li>\n    <li>The maximum divisor to look for is the square root of your number, because: n=a&timesb and if both values\n        were greater than the square root of n, a&timesb would be larger than n. Therefore at least one of those\n        factors must be less than or equal to the square root of n.\n    </li>\n</ul>\n",
		"category": "Primes and Factors",
		"categorySlug": "primes-and-factors"
	  },
	  {
		"name": "Highest Common Factor",
		"slug": "highest-common-factor",
		"description": "Solve: Highest Common Factor, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "First Number", "tag": "num_1" },
		  { "name": "Second Number", "tag": "num_2" }
		],
		"example": "<div>\n  <h4>Question:</h4>\n\n  <p>Find the highest common factor of 600 and 752.</p>\n\n  <h4>Answer:</h4>\n\n  <p>The highest common factor of 600 and 752 is 8.</p>\n\n  <h4>Here's the working out:</h4>\n\n  <p>Finding all prime factors of 600:</p>\n\n  <table class=\"pure-table pure-table-bordered\">\n    <tbody>\n      <tr>\n        <td>600 | 2</td>\n        <td>2 is a factor of 600</td>\n        <td>600 divided by 2 is 300</td>\n      </tr>\n      <tr>\n        <td>300 | 2</td>\n        <td>2 is a factor of 300</td>\n        <td>300 divided by 2 is 150</td>\n      </tr>\n      <tr>\n        <td>150 | 2</td>\n        <td>2 is a factor of 150</td>\n        <td>150 divided by 2 is 75</td>\n      </tr>\n      <tr>\n        <td>75 | 3</td>\n        <td>3 is a factor of 75</td>\n        <td>75 divided by 3 is 25</td>\n      </tr>\n      <tr>\n        <td>25 | 5</td>\n        <td>5 is a factor of 25</td>\n        <td>25 divided by 5 is 5</td>\n      </tr>\n      <tr>\n        <td>5 | 5</td>\n        <td>5 is a factor of 5</td>\n        <td>5 divided by 5 is 1</td>\n      </tr>\n      <tr>\n        <td>1     </td>\n        <td></td>\n        <td></td>\n      </tr>\n    </tbody>\n  </table>\n\n  <p>Finding all prime factors of 752:</p>\n\n  <table class=\"pure-table pure-table-bordered\">\n    <tbody>\n      <tr>\n        <td>752 | 2</td>\n        <td>2 is a factor of 752</td>\n        <td>752 divided by 2 is 376</td>\n      </tr>\n      <tr>\n        <td>376 | 2</td>\n        <td>2 is a factor of 376</td>\n        <td>376 divided by 2 is 188</td>\n      </tr>\n      <tr>\n        <td>188 | 2</td>\n        <td>2 is a factor of 188</td>\n        <td>188 divided by 2 is 94</td>\n      </tr>\n      <tr>\n        <td>94 | 2</td>\n        <td>2 is a factor of 94</td>\n        <td>94 divided by 2 is 47</td>\n      </tr>\n      <tr>\n        <td>47 | 47</td>\n        <td>47 is a factor of 47</td>\n        <td>47 divided by 47 is 1</td>\n      </tr>\n      <tr>\n        <td>1     </td>\n        <td></td>\n        <td></td>\n      </tr>\n    </tbody>\n  </table>\n\n  <p>\n    Prime factors for the first number are:<br />\n    2 × 2 × 2 × 3 × 5 × 5 = 600<br />2<sup>3</sup> × 3<sup>1</sup> × 5<sup>2</sup> = 600<br />\n  </p>\n\n  <p>\n    Prime factors for the second number are:<br />\n    2 × 2 × 2 × 2 × 47 = 752<br />2<sup>4</sup> × 47<sup>1</sup> = 752<br />\n  </p>\n\n  <p>\n    Find the primes that are shared between the two numbers:<br />\n    2, 2, 2\n  </p>\n\n  <p>\n    Take the shared primes and multiply them together:<br />\n    2 × 2 × 2 = 8\n  </p>\n\n  <h4>Therefore:</h4>\n\n  <p>The highest common factor of 600 and 752 is 8.</p>\n</div>",
		"category": "Primes and Factors",
		"categorySlug": "primes-and-factors"
	  },
	  {
		"name": "Lowest Common Multiple",
		"slug": "lowest-common-multiple",
		"description": "Solve: Lowest Common Multiple, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "First Number", "tag": "num_1" },
		  { "name": "Second Number", "tag": "num_2" }
		],
		"example": "<div>\n  <h4>Question:</h4>\n\n  <p>Find the lowest common multiple of 600 and 752.</p>\n\n  <h4>Answer:</h4>\n\n  <p>The lowest common multiple of 600 and 752 is 56400.</p>\n\n  <h4>Here's the working out:</h4>\n\n  <p>Finding all prime factors of 600:</p>\n\n  <table class=\"pure-table pure-table-bordered\">\n    <tbody>\n      <tr>\n        <td>600 | 2</td>\n        <td>2 is a factor of 600</td>\n        <td>600 divided by 2 is 300</td>\n      </tr>\n      <tr>\n        <td>300 | 2</td>\n        <td>2 is a factor of 300</td>\n        <td>300 divided by 2 is 150</td>\n      </tr>\n      <tr>\n        <td>150 | 2</td>\n        <td>2 is a factor of 150</td>\n        <td>150 divided by 2 is 75</td>\n      </tr>\n      <tr>\n        <td>75 | 3</td>\n        <td>3 is a factor of 75</td>\n        <td>75 divided by 3 is 25</td>\n      </tr>\n      <tr>\n        <td>25 | 5</td>\n        <td>5 is a factor of 25</td>\n        <td>25 divided by 5 is 5</td>\n      </tr>\n      <tr>\n        <td>5 | 5</td>\n        <td>5 is a factor of 5</td>\n        <td>5 divided by 5 is 1</td>\n      </tr>\n      <tr>\n        <td>1     </td>\n        <td></td>\n        <td></td>\n      </tr>\n    </tbody>\n  </table>\n\n  <p>Finding all prime factors of 752:</p>\n\n  <table class=\"pure-table pure-table-bordered\">\n    <tbody>\n      <tr>\n        <td>752 | 2</td>\n        <td>2 is a factor of 752</td>\n        <td>752 divided by 2 is 376</td>\n      </tr>\n      <tr>\n        <td>376 | 2</td>\n        <td>2 is a factor of 376</td>\n        <td>376 divided by 2 is 188</td>\n      </tr>\n      <tr>\n        <td>188 | 2</td>\n        <td>2 is a factor of 188</td>\n        <td>188 divided by 2 is 94</td>\n      </tr>\n      <tr>\n        <td>94 | 2</td>\n        <td>2 is a factor of 94</td>\n        <td>94 divided by 2 is 47</td>\n      </tr>\n      <tr>\n        <td>47 | 47</td>\n        <td>47 is a factor of 47</td>\n        <td>47 divided by 47 is 1</td>\n      </tr>\n      <tr>\n        <td>1     </td>\n        <td></td>\n        <td></td>\n      </tr>\n    </tbody>\n  </table>\n\n  <p>\n    Prime factors for the first number are:<br />\n    2 × 2 × 2 × 3 × 5 × 5 = 600<br />2<sup>3</sup> × 3<sup>1</sup> × 5<sup>2</sup> = 600<br />\n  </p>\n\n  <p>\n    Prime factors for the second number are:<br />\n    2 × 2 × 2 × 2 × 47 = 752<br />2<sup>4</sup> × 47<sup>1</sup> = 752<br />\n  </p>\n\n  <p>\n    Compare the primes of the first and second numbers and use the sets with the highest\n    exponent:<br />\n    47<sup>1</sup> × 2<sup>4</sup> × 3<sup>1</sup> × 5<sup>2</sup> = 56400\n  </p>\n\n  <h4>Therefore:</h4>\n\n  <p>The lowest common multiple of 600 and 752 is 56400.</p>\n</div>",
		"category": "Primes and Factors",
		"categorySlug": "primes-and-factors"
	  }
	]
  },
  {
	"name": "Networking",
	"slug": "networking",
	"description": "Calculations that you might use for Networking/Computer Science: Binary to Decimal, Binary to Hexadecimal, Decimal to Binary, Decimal to Hexadecimal, Hexadecimal to Binary, and Hexadecimal to Decimal.",
	"calculations": [
	  {
		"name": "Binary to Decimal",
		"slug": "binary-to-decimal",
		"description": "Solve: Binary to Decimal, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Binary", "tag": "binary" }],
		"example": "<h4>Question:</h4>\n\n<p>What is 10101 in decimal?</p>\n\n<h4>Answer:</h4>\n\n<p>The binary number <b>10101</b> in decimal is <b>48</b>.</p>\n\n<h4>Here's how to convert it:</h4>\n\n<p>Using the following continuing equation the accurate answer can be found: (b x 2<sup>n</sup>)…<br>\n    Where <strong>b</strong> is the binary number at the n<sup>th</sup> position and <strong>n</strong> is the\n    n<sup>th</sup> position.<br>\n    There are 5 numbers in this binary number, the n<sup>th</sup> position starts from 0, and therefore\n    the n<sup>th</sup> position goes up to 4 (5 - 1). Count down from 4\n    and insert the binary number at that position.</p>\n\n<p>Using the equation for the binary number <span class=\"word-break\">10101</span>, it would go like this:<br>\n    (b=1 x 2<sup>n=16</sup>) + (b=0 x 2<sup>n=16</sup>) + (b=1 x 2<sup>n=16</sup>) + (b=0 x 2<sup>n=16</sup>) + (b=1 x 2<sup>n=16</sup>)<br>(1 x 2<sup>16</sup>) + (0 x 2<sup>16</sup>) + (1 x 2<sup>16</sup>) + (0 x 2<sup>16</sup>) + (1 x 2<sup>16</sup>)<br>(1 x 16) + (0 x 16) + (1 x 16) + (0 x 16) + (1 x 16)<br>(16) + (0) + (16) + (0) + (16)<br>\n    <span class=\"word-break\">48</span></p>\n\n<h4>Therefore:</h4>\n\n<p class=\"word-break\">(10101)<sub>2</sub> = (48)<sub>10</sub></p>\n",
		"category": "Networking",
		"categorySlug": "networking"
	  },
	  {
		"name": "Binary to Hexadecimal",
		"slug": "binary-to-hexadecimal",
		"description": "Solve: Binary to Hexadecimal, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Binary", "tag": "binary" }],
		"example": "<h4>Question:</h4>\n\n<p>What is 10111 in hexadecimal?</p>\n\n<h4>Answer:</h4>\n\n<p>The binary number <b>10111</b> in hexadecimal is <b>17</b>.</p>\n\n<h4>Here's how to convert it:</h4>\n\n<p>Break the binary number <span class=\"word-break\">10111</span> into groups of 4 digits as such:<br>\n    00010111</p>\n\n<p>Convert those groups (base 2) to a decimal number (base 10), that way it is easy to convert them to a Hexadecimal\n    number (base 16):<br>\n    Refer to the Conversion Table <a href=\"/category/networking/conversion-table\" rel=\"noopener noreferrer\"\n                                     target=\"_blank\">here</a>.<br>\n    (0001)<sub>2</sub> = (1)<sub>10</sub> = (1)<sub>16</sub><br>(0111)<sub>2</sub> = (7)<sub>10</sub> = (7)<sub>16</sub><br></p>\n\n<h4>Therefore:</h4>\n\n<p class=\"word-break\">(10111)<sub>2</sub> = (17)<sub>16</sub></p>",
		"category": "Networking",
		"categorySlug": "networking"
	  },
	  {
		"name": "Decimal to Binary",
		"slug": "decimal-to-binary",
		"description": "Solve: Decimal to Binary, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Decimal", "tag": "decimal" }],
		"example": "<h4>Question:</h4>\n\n<p>What is 21 in binary?</p>\n\n<h4>Answer:</h4>\n\n<p>The decimal number <b>21</b> in binary is <b>10101</b>.\n</p>\n\n<h4>Here's how to convert it:</h4>\n\n<p>Divide the decimal number by 2 and record the remainder:<br>\n    21 &divide; 2 = 10, Remainder: 1<br>10 &divide; 2 = 5, Remainder: 0<br>5 &divide; 2 = 2, Remainder: 1<br>2 &divide; 2 = 1, Remainder: 0<br>1 &divide; 2 = 0, Remainder: 1<br></p>\n\n<p>Join the remainders from last to first:<br>\n    <span class=\"word-break\">10101</span></p>\n\n\n<h4>Therefore:</h4>\n\n<p class=\"word-break\">(21)<sub>10</sub> = (<span class=\"word-break\">10101</span>)<sub>2</sub>\n</p>\n",
		"category": "Networking",
		"categorySlug": "networking"
	  },
	  {
		"name": "Decimal to Hexadecimal",
		"slug": "decimal-to-hexadecimal",
		"description": "Solve: Decimal to Hexadecimal, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Decimal", "tag": "decimal" }],
		"example": "<h4>Question:</h4>\n\n<p>What is 92 in hexadecimal?</p>\n\n<h4>Answer:</h4>\n\n<p>The decimal number <b>92</b> in hexadecimal is <b>5C</b>.\n</p>\n\n<h4>Here's how to convert it:</h4>\n\n<p>Divide the decimal number by 16 and record the remainder:<br>\n    92 &divide; 16 = 5, Remainder: 12<br>5 &divide; 16 = 0, Remainder: 5<br></p>\n\n<p>Join the remainders from last to first:<br>\n    <span class=\"word-break\">5C</span></p>\n<p>Convert remainders to their hexadecimal equivalents:<br>\n    Refer to the Conversion Table <a href=\"/category/networking/conversion-table\" rel=\"noopener noreferrer\"\n                                     target=\"_blank\">here</a>.<br>\n    (12)<sub>10</sub> = (C)<sub>16</sub><br>(5)<sub>10</sub> = (5)<sub>16</sub><br></p>\n\n<h4>Therefore:</h4>\n\n<p class=\"word-break\">(92)<sub>10</sub> = (<span class=\"word-break\">5C</span>)<sub>16</sub>\n</p>\n",
		"category": "Networking",
		"categorySlug": "networking"
	  },
	  {
		"name": "Hexadecimal to Binary",
		"slug": "hexadecimal-to-binary",
		"description": "Solve: Hexadecimal to Binary, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Hexadecimal", "tag": "hexadecimal" }],
		"example": "<h4>Question:</h4>\n\n<p>What is 6BA in binary?</p>\n\n<h4>Answer:</h4>\n\n<p>The hexadecimal number <b>6BA</b> in binary is <b>01101011101</b>.</p>\n\n<h4>Here's how to convert it:</h4>\n\n<p>Convert each hexadecimal digit manually to its binary equivalent:<br>\n    Refer to the Conversion Table <a href=\"/category/networking/conversion-table\" rel=\"noopener noreferrer\"\n                                     target=\"_blank\">here</a>.<br>\n    (6)<sub>16</sub> = (110)<sub>2</sub><br>(B)<sub>16</sub> = (1011)<sub>2</sub><br>(A)<sub>16</sub> = (1010)<sub>2</sub><br></p>\n\n<p>Join the digit equivalents together from first to last:<br>\n    01101011101</p>\n\n<h4>Therefore:</h4>\n\n<p>(6BA)<sub>16</sub> = (01101011101)<sub>2</sub></p>\n",
		"category": "Networking",
		"categorySlug": "networking"
	  },
	  {
		"name": "Hexadecimal to Decimal",
		"slug": "hexadecimal-to-decimal",
		"description": "Solve: Hexadecimal to Decimal, showing mathematical proof and answer.",
		"inputInfo": [{ "name": "Hexadecimal", "tag": "hexadecimal" }],
		"example": "<h4>Question:</h4>\n\n<p>What is 6BA in decimal?</p>\n\n<h4>Answer:</h4>\n\n<p>The hexadecimal number <b>6BA</b> in decimal is <b>1722</b>.</p>\n\n<h4>Here's how to convert it:</h4>\n\n<p>Get the decimal equivalent of each hexadecimal digit:<br>\n    Refer to the Conversion Table <a href=\"/category/networking/conversion-table\" rel=\"noopener noreferrer\"\n                                     target=\"_blank\">here</a>.<br>\n    (6)<sub>16</sub> = (6)<sub>10</sub><br>(B)<sub>16</sub> = (11)<sub>10</sub><br>(A)<sub>16</sub> = (10)<sub>10</sub><br></p>\n\n<p>Multiply each hexadecimal digit with 16 power of digit position. Digit position starts from zero, right to\n    left:<br>\n    6BA = (6 x 16<sup>2</sup>) + (11 x 16<sup>1</sup>) + (10 x 16<sup>0</sup>)<br>6BA = (6 x 256) + (11 x 16) + (10 x 1)<br>6BA = 1536 + 176 + 10<br>6BA = 1722<br></p>\n\n<h4>Therefore:</h4>\n\n<p class=\"word-break\">(6BA)<sub>16</sub> = (1722)<sub>10</sub></p>",
		"category": "Networking",
		"categorySlug": "networking"
	  }
	]
  },
  {
	"name": "Percentages",
	"slug": "percentages",
	"description": "Calculations for percentages! Find the value from a percentage, find a percentage from a value, or find the percentage change between two values.",
	"calculations": [
	  {
		"name": "Change Number by Percentage",
		"slug": "change-number-by-percentage",
		"description": "Solve: Change Number by Percentage, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Number", "tag": "number" },
		  { "name": "Percentage", "tag": "percentage" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where the increase in percentage is 65% and the number is 900, find the resulting number.</p>\n\n<h4>Answer:</h4>\n\n<p>The number 900 increased by 65% is 1485.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Calculate using the equation shown below:<br>\n    (Number &#xf7; 100) &times Percentage + Number<br>\n    (900 &#xf7; 100) &times 65 + 900</p>\n\n<p>Find 1% of 900:<br>\n    Number &#xf7; 100<br>\n    900 &#xf7; 100<br>\n    9</p>\n\n<p>Multiply the one percent value by 65 to find 65% of 900:<br>\n    9 &times Percentage<br>\n    9 &times 65<br>\n    585</p>\n\n<p>The value 585 is 65% of 900, add it to the Number to get the final answer:<br>\n    585 + Number<br>\n    585 + 900<br>\n    1485</p>\n\n<h4>Therefore:</h4>\n\n<p>900 increased by 65% is 1485.</p>\n",
		"category": "Percentages",
		"categorySlug": "percentages"
	  },
	  {
		"name": "Get Number from a Percentage",
		"slug": "get-number-from-a-percentage",
		"description": "Solve: Get Number from a Percentage, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Percentage", "tag": "percentage" },
		  { "name": "Number", "tag": "number" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where the percentage is 600% and the number is 752, find the resulting number of the\n    percentage.</p>\n\n<h4>Answer:</h4>\n\n<p>600% of 752 is 4512.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Use the equation below:<br>\n    (Percentage &#xf7; 100) &times Number<br>\n    (600 &#xf7; 100) &times 752</p>\n\n<p>Find the decimalised version of the Percentage:<br>\n    Percentage &#xf7; 100<br>\n    600 &#xf7; 100<br>\n    6</p>\n\n<p>Now multiply it with the Number to get the final answer:<br>\n    (6) &times Number<br>\n    (6) &times 752<br>\n    4512</p>\n\n<h4>Therefore:</h4>\n\n<p>600% of 752 is 4512.</p>",
		"category": "Percentages",
		"categorySlug": "percentages"
	  },
	  {
		"name": "Find Percentage Difference of Two Numbers",
		"slug": "find-percentage-difference-of-two-numbers",
		"description": "Solve: Find Percentage Difference of Two Numbers, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Number", "tag": "number" },
		  { "name": "New Number", "tag": "new_number" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where original number is 400 and the new number is 540, find the difference between the numbers in\n    percentage.</p>\n\n<h4>Answer:</h4>\n\n<p>The difference in percentage between 400 and 540 is 35%.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Calculate using the equation shown below:<br>\n    ((New Number - Original Number) &#xf7; Original Number) &times 100<br>\n\n<p>Find the change in value between the two numbers:<br>\n    New Number - Original Number<br>\n    540 - 400<br>\n    140</p>\n\n<p>Find the decimalised version of the percentage change using the change in value:<br>\n    140 &#xf7; Original Number<br>\n    140 &#xf7; 400<br>\n    0.35</p>\n\n<p>Convert decimalised percentage to percentage to get the final answer:<br>\n    0.35 &times 100<br>\n    35</p>\n\n<h4>Therefore:</h4>\n\n<p>The percentage of change from 400 to 540 is 35%.</p>",
		"category": "Percentages",
		"categorySlug": "percentages"
	  },
	  {
		"name": "Find Percentage of a Number",
		"slug": "find-percentages-of-a-number",
		"description": "Solve: Find Percentage of a Number, showing mathematical proof and answer.",
		"inputInfo": [
		  { "name": "Number", "tag": "number" },
		  { "name": "Total Number", "tag": "total_number" }
		],
		"example": "<h4>Question:</h4>\n\n<p>Where the number is 585 and the total number is 900, find how much percentage 585 is of\n    900.</p>\n\n<h4>Answer:</h4>\n\n<p>585 is 65% of 900.</p>\n\n<h4>Here's how to calculate it:</h4>\n\n<p>Calculate using the equation shown below:<br>\n    (Number &#xf7; Total Number) &times 100</p>\n\n<p>Find the decimalised version of the percentage:<br>\n    Number &#xf7; Total Number<br>\n    585 &#xf7; 900<br>\n    0.65</p>\n\n<p>Convert decimalised percentage to percentage to get the answer:<br>\n    0.65 &times 100<br>\n    65</p>\n\n<h4>Therefore:</h4>\n\n<p>585 is 65% of 900.</p>",
		"category": "Percentages",
		"categorySlug": "percentages"
	  }
	]
  }
]`

func GetCategoriesJson() string {
	return categoriesJson
}

func FindMath(slug string) (Mathematics, error) {
	calculations := []CalculationShort{
		// Networking
		{
			Name:     "Binary to Decimal",
			Slug:     "binary-to-decimal",
			Math:     &BinaryToDecimalAPI{},
			Category: "Networking",
		},
		{
			Name:     "Binary to Hexadecimal",
			Slug:     "binary-to-hexadecimal",
			Math:     &BinaryToHexadecimalAPI{},
			Category: "Networking",
		},
		{
			Name:     "Decimal to Binary",
			Slug:     "decimal-to-binary",
			Math:     &DecimalToBinaryAPI{},
			Category: "Networking",
		},
		{
			"Decimal to Hexadecimal",
			"decimal-to-hexadecimal",
			&DecimalToHexadecimalAPI{},
			"Networking",
		},
		{
			"Hexadecimal to Binary",
			"hexadecimal-to-binary",
			&HexadecimalToBinaryAPI{},
			"Networking",
		},
		{
			"Hexadecimal to Decimal",
			"hexadecimal-to-decimal",
			&HexadecimalToDecimalAPI{},
			"Networking",
		},

		// Numbers
		{
			"Find if Number is a Prime Number",
			"is-prime",
			&IsPrimeAPI{},
			"Primes and Factors",
		},
		{
			"Highest Common Factor",
			"highest-common-factor",
			&HighestCommonFactorAPI{},
			"Primes and Factors",
		},
		{
			"Lowest Common Multiple",
			"lowest-common-multiple",
			&LowestCommonMultipleAPI{},
			"Primes and Factors",
		},

		// Percentages
		{
			"Change Number by Percentage",
			"change-number-by-percentage",
			&ChangeByPercentageAPI{},
			"Percentages",
		},
		{
			"Get Number from a Percentage",
			"get-number-from-a-percentage",
			&NumberFromPercentageAPI{},
			"Percentages",
		},
		{
			"Find Percentage Difference of Two Numbers",
			"find-percentage-difference-of-two-numbers",
			&PercentageChangeAPI{},
			"Percentages",
		},
		{
			"Find Percentage of a Number",
			"find-percentages-of-a-number",
			&PercentageFromNumberAPI{},
			"Percentages",
		},

		// Total Surface Area
		{
			"Pythagorean Theorem",
			"pythagorean-theorem",
			&TsaPythagoreanTheoremAPI{},
			"Total Surface Area",
		},
		{
			"Cone",
			"cone",
			&TsaConeAPI{},
			"Total Surface Area",
		},
		{
			"Cube",
			"cube",
			&TsaCubeAPI{},
			"Total Surface Area",
		},
		{
			"Cylinder",
			"cylinder",
			&TsaCylinderAPI{},
			"Total Surface Area",
		},
		{
			"Rectangular Prism",
			"rectangular-prism",
			&TsaRectangularPrismAPI{},
			"Total Surface Area",
		},
		{
			"Sphere",
			"sphere",
			&TsaSphereAPI{},
			"Total Surface Area",
		},
		{
			"Square Based Triangle",
			"square-based-triangle",
			&TsaSquareBaseTriangleAPI{},
			"Total Surface Area",
		},
	}

	for _, calc := range calculations {
		if calc.Slug == slug {
			return calc.Math, nil
		}
	}
	return nil, errors.New("slug not found")
}
