// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../../gen_tests/template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
    r "gopkg.in/dancannon/gorethink.v2"
	"gopkg.in/dancannon/gorethink.v2/internal/compare"
)

// Tests of nested arithmetic expressions
func TestMathLogicMathSuite(t *testing.T) {
	suite.Run(t, new(MathLogicMathSuite ))
}

type MathLogicMathSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicMathSuite) SetupTest() {
	suite.T().Log("Setting up MathLogicMathSuite")
	// Use imports to prevent errors
	_ = time.Time{}
    _ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MathLogicMathSuite) TearDownSuite() {
	suite.T().Log("Tearing down MathLogicMathSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MathLogicMathSuite) TestCases() {
	suite.T().Log("Running MathLogicMathSuite: Tests of nested arithmetic expressions")



	{
		// math_logic/math.yaml line #4
		/* 1 */
		var expected_ int = 1
		/* (((4 + 2 * (r.expr(26) % 18)) / 5) - 3) */

		suite.T().Log("About to run line #4: r.Add(4, r.Mul(2, r.Expr(26).Mod(18))).Div(5).Sub(3)")

		runAndAssert(suite.Suite, expected_, r.Add(4, r.Mul(2, r.Expr(26).Mod(18))).Div(5).Sub(3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #4")
	}
}
