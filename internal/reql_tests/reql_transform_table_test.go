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

// Tests manipulation operations on tables
func TestTransformTableSuite(t *testing.T) {
	suite.Run(t, new(TransformTableSuite ))
}

type TransformTableSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TransformTableSuite) SetupTest() {
	suite.T().Log("Setting up TransformTableSuite")
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

	r.DB("test").TableDrop("tbl").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *TransformTableSuite) TearDownSuite() {
	suite.T().Log("Tearing down TransformTableSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		 r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TransformTableSuite) TestCases() {
	suite.T().Log("Running TransformTableSuite: Tests manipulation operations on tables")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors


	{
		// transform/table.yaml line #5
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* tbl.insert([{"a":["k1","v1"]},{"a":["k2","v2"]}]) */

		suite.T().Log("About to run line #5: tbl.Insert([]interface{}{map[interface{}]interface{}{'a': []interface{}{'k1', 'v1'}, }, map[interface{}]interface{}{'a': []interface{}{'k2', 'v2'}, }})")

		runAndAssert(suite.Suite, expected_, tbl.Insert([]interface{}{map[interface{}]interface{}{"a": []interface{}{"k1", "v1"}, }, map[interface{}]interface{}{"a": []interface{}{"k2", "v2"}, }}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// transform/table.yaml line #10
		/* {"k1":"v1","k2":"v2"} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"k1": "v1", "k2": "v2", }
		/* tbl.map(r.row["a"]).coerce_to("object") */

		suite.T().Log("About to run line #10: tbl.Map(r.Row.AtIndex('a')).CoerceTo('object')")

		runAndAssert(suite.Suite, expected_, tbl.Map(r.Row.AtIndex("a")).CoerceTo("object"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// transform/table.yaml line #14
		/* "SELECTION<STREAM>" */
		var expected_ string = "SELECTION<STREAM>"
		/* tbl.limit(1).type_of() */

		suite.T().Log("About to run line #14: tbl.Limit(1).TypeOf()")

		runAndAssert(suite.Suite, expected_, tbl.Limit(1).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #14")
	}

	{
		// transform/table.yaml line #17
		/* "ARRAY" */
		var expected_ string = "ARRAY"
		/* tbl.limit(1).coerce_to('array').type_of() */

		suite.T().Log("About to run line #17: tbl.Limit(1).CoerceTo('array').TypeOf()")

		runAndAssert(suite.Suite, expected_, tbl.Limit(1).CoerceTo("array").TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #17")
	}
}
