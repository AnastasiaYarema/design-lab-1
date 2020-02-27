package implementation_testing

import (
	"testing"
	. "gopkg.in/check.v1"
	. "github.com/AnastasiaYarema/design-lab-1/implementation"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct {}

var _ = Suite(&TestSuite{})

func (s * TestSuite) TestPositivePostfixExpressionWithTwoOperands(c *C) {
	result, err := CalculatePostfixExpression("10 11.5 +")

	c.Assert(err, IsNil)
	c.Assert(result, Equals, 21.5)
}

func (s * TestSuite) TestPositivePostfixExpressionWithTwoOperandsNegativeNumbers(c *C) {
	result, err := CalculatePostfixExpression("-10 -11.5 +")

	c.Assert(err, IsNil)
	c.Assert(result, Equals, -21.5)
}

func (s * TestSuite) TestPositivePostfixExpressionWithManyWhitespaces(c *C) {
	result, err := CalculatePostfixExpression("10    11.5    +")

	c.Assert(err, IsNil)
	c.Assert(result, Equals, 21.5)
}
  
func (s * TestSuite) TestPositivePostfixExpressionWithThreeOperands(c *C) {
	result, err := CalculatePostfixExpression("10.0 11 + 2 ^")

	c.Assert(err, IsNil)
	c.Assert(result, Equals, float64(21 * 21))
}

func (s * TestSuite) TestPositivePostfixExpressionWithSevenOperands(c *C) {
	result, err := CalculatePostfixExpression("4 2 - 3.0 * 5 + 10 / 10 * 5 +")

	c.Assert(err, IsNil)
	c.Assert(result, Equals, 16.0)
}

func (s * TestSuite) TestPositivePostfixExpressionWithTenOperands(c *C) {
	result, err := CalculatePostfixExpression("4 2 - 3 * 5 + 10 / 10 * 5 + 2 ^ 16 / 100500 +")

	c.Assert(err, IsNil)
	c.Assert(result, Equals, 100516.0)
}

func (s * TestSuite) TestNegativePostfixExpressionWithEmptyString(c *C) {
	_, err := CalculatePostfixExpression("")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Matches, "Postfix expression string cannot be empty")
}

func (s * TestSuite) TestNegativePostfixExpressionWithWrongCharacters(c *C) {
	_, err := CalculatePostfixExpression("10.0 11 + 2 (")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Matches, "Postfix expression contains invalid token")
}

func (s * TestSuite) TestNegativePostfixExpressionWithWrongNumberOfOperands(c *C) {
	_, err := CalculatePostfixExpression("10.0 11 + 2")

	c.Assert(err, NotNil)
	c.Assert(err.Error(), Matches, "Wrong postfix expression")
}