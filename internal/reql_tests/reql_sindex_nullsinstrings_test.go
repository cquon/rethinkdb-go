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

// sindex nulls in strings
func TestSindexNullsinstringsSuite(t *testing.T) {
	suite.Run(t, new(SindexNullsinstringsSuite ))
}

type SindexNullsinstringsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *SindexNullsinstringsSuite) SetupTest() {
	suite.T().Log("Setting up SindexNullsinstringsSuite")
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

func (suite *SindexNullsinstringsSuite) TearDownSuite() {
	suite.T().Log("Tearing down SindexNullsinstringsSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		 r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *SindexNullsinstringsSuite) TestCases() {
	suite.T().Log("Running SindexNullsinstringsSuite: sindex nulls in strings")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors


	{
		// sindex/nullsinstrings.yaml line #4
		/* ({"created":1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1, }
		/* tbl.index_create("key") */

		suite.T().Log("About to run line #4: tbl.IndexCreate('key')")

		runAndAssert(suite.Suite, expected_, tbl.IndexCreate("key"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #4")
	}

	{
		// sindex/nullsinstrings.yaml line #6
		/* ([{"ready":true}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"ready": true, }}
		/* tbl.index_wait().pluck("ready") */

		suite.T().Log("About to run line #6: tbl.IndexWait().Pluck('ready')")

		runAndAssert(suite.Suite, expected_, tbl.IndexWait().Pluck("ready"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// sindex/nullsinstrings.yaml line #10
		/* ({"inserted":2}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"inserted": 2, }
		/* tbl.insert([{"id":1,"key":["a","b"]},{"id":2,"key":["a\u0000Sb"]}]).pluck("inserted") */

		suite.T().Log("About to run line #10: tbl.Insert([]interface{}{map[interface{}]interface{}{'id': 1, 'key': []interface{}{'a', 'b'}, }, map[interface{}]interface{}{'id': 2, 'key': []interface{}{'a\\u0000Sb'}, }}).Pluck('inserted')")

		runAndAssert(suite.Suite, expected_, tbl.Insert([]interface{}{map[interface{}]interface{}{"id": 1, "key": []interface{}{"a", "b"}, }, map[interface{}]interface{}{"id": 2, "key": []interface{}{"a\u0000Sb"}, }}).Pluck("inserted"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// sindex/nullsinstrings.yaml line #13
		/* ([{"id":2}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 2, }}
		/* tbl.get_all(["a\u0000Sb"], index="key").pluck("id") */

		suite.T().Log("About to run line #13: tbl.GetAll([]interface{}{'a\\u0000Sb'}).OptArgs(r.GetAllOpts{Index: 'key', }).Pluck('id')")

		runAndAssert(suite.Suite, expected_, tbl.GetAll([]interface{}{"a\u0000Sb"}).OptArgs(r.GetAllOpts{Index: "key", }).Pluck("id"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #13")
	}

	{
		// sindex/nullsinstrings.yaml line #18
		/* ([{"id":1}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"id": 1, }}
		/* tbl.get_all(["a","b"], index="key").pluck("id") */

		suite.T().Log("About to run line #18: tbl.GetAll([]interface{}{'a', 'b'}).OptArgs(r.GetAllOpts{Index: 'key', }).Pluck('id')")

		runAndAssert(suite.Suite, expected_, tbl.GetAll([]interface{}{"a", "b"}).OptArgs(r.GetAllOpts{Index: "key", }).Pluck("id"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #18")
	}
}
