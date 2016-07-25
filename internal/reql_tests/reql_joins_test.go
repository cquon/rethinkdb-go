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

// Tests that manipulation data in tables
func TestJoinsSuite(t *testing.T) {
	suite.Run(t, new(JoinsSuite ))
}

type JoinsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *JoinsSuite) SetupTest() {
	suite.T().Log("Setting up JoinsSuite")
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

	r.DB("test").TableDrop("messages").Exec(suite.session)
	err = r.DB("test").TableCreate("messages").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("messages").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("otbl").Exec(suite.session)
	err = r.DB("test").TableCreate("otbl").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("otbl").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("otbl2").Exec(suite.session)
	err = r.DB("test").TableCreate("otbl2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("otbl2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("receivers").Exec(suite.session)
	err = r.DB("test").TableCreate("receivers").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("receivers").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("senders").Exec(suite.session)
	err = r.DB("test").TableCreate("senders").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("senders").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("tbl").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("test").TableDrop("tbl2").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *JoinsSuite) TearDownSuite() {
	suite.T().Log("Tearing down JoinsSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		 r.DB("test").TableDrop("messages").Exec(suite.session)
		 r.DB("test").TableDrop("otbl").Exec(suite.session)
		 r.DB("test").TableDrop("otbl2").Exec(suite.session)
		 r.DB("test").TableDrop("receivers").Exec(suite.session)
		 r.DB("test").TableDrop("senders").Exec(suite.session)
		 r.DB("test").TableDrop("tbl").Exec(suite.session)
		 r.DB("test").TableDrop("tbl2").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *JoinsSuite) TestCases() {
	suite.T().Log("Running JoinsSuite: Tests that manipulation data in tables")

	messages := r.DB("test").Table("messages")
	_ = messages // Prevent any noused variable errors
	otbl := r.DB("test").Table("otbl")
	_ = otbl // Prevent any noused variable errors
	otbl2 := r.DB("test").Table("otbl2")
	_ = otbl2 // Prevent any noused variable errors
	receivers := r.DB("test").Table("receivers")
	_ = receivers // Prevent any noused variable errors
	senders := r.DB("test").Table("senders")
	_ = senders // Prevent any noused variable errors
	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors
	tbl2 := r.DB("test").Table("tbl2")
	_ = tbl2 // Prevent any noused variable errors


	{
		// joins.yaml line #7
		/* partial({'tables_created':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_created": 1, })
		/* r.db('test').table_create('test3', primary_key='foo') */

		suite.T().Log("About to run line #7: r.DB('test').TableCreate('test3').OptArgs(r.TableCreateOpts{PrimaryKey: 'foo', })")

		runAndAssert(suite.Suite, expected_, r.DB("test").TableCreate("test3").OptArgs(r.TableCreateOpts{PrimaryKey: "foo", }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #7")
	}

	// joins.yaml line #11
	// tbl3 = r.db('test').table('test3')
	suite.T().Log("Possibly executing: var tbl3 r.Term = r.DB('test').Table('test3')")

	tbl3 := r.DB("test").Table("test3")
	_ = tbl3 // Prevent any noused variable errors


	{
		// joins.yaml line #13
		/* partial({'errors':0, 'inserted':100}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 100, })
		/* tbl.insert(r.range(0, 100).map({'id':r.row, 'a':r.row % 4})) */

		suite.T().Log("About to run line #13: tbl.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{'id': r.Row, 'a': r.Row.Mod(4), }))")

		runAndAssert(suite.Suite, expected_, tbl.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{"id": r.Row, "a": r.Row.Mod(4), })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #13")
	}

	{
		// joins.yaml line #18
		/* partial({'errors':0, 'inserted':100}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 100, })
		/* tbl2.insert(r.range(0, 100).map({'id':r.row, 'b':r.row % 4})) */

		suite.T().Log("About to run line #18: tbl2.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{'id': r.Row, 'b': r.Row.Mod(4), }))")

		runAndAssert(suite.Suite, expected_, tbl2.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{"id": r.Row, "b": r.Row.Mod(4), })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #18")
	}

	{
		// joins.yaml line #23
		/* partial({'errors':0, 'inserted':100}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 100, })
		/* tbl3.insert(r.range(0, 100).map({'foo':r.row, 'b':r.row % 4})) */

		suite.T().Log("About to run line #23: tbl3.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{'foo': r.Row, 'b': r.Row.Mod(4), }))")

		runAndAssert(suite.Suite, expected_, tbl3.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{"foo": r.Row, "b": r.Row.Mod(4), })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// joins.yaml line #28
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* otbl.insert(r.range(1,100).map({'id': r.row, 'a': r.row})) */

		suite.T().Log("About to run line #28: otbl.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{'id': r.Row, 'a': r.Row, }))")

		runAndAssert(suite.Suite, expected_, otbl.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{"id": r.Row, "a": r.Row, })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// joins.yaml line #29
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* otbl2.insert(r.range(1,100).map({'id': r.row, 'b': 2 * r.row})) */

		suite.T().Log("About to run line #29: otbl2.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{'id': r.Row, 'b': r.Mul(2, r.Row), }))")

		runAndAssert(suite.Suite, expected_, otbl2.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{"id": r.Row, "b": r.Mul(2, r.Row), })), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #29")
	}

	// joins.yaml line #34
	// ij = tbl.inner_join(tbl2, lambda x,y:x['a'] == y['b']).zip()
	suite.T().Log("Possibly executing: var ij r.Term = tbl.InnerJoin(tbl2, func(x r.Term, y r.Term) interface{} { return x.AtIndex('a').Eq(y.AtIndex('b'))}).Zip()")

	ij := tbl.InnerJoin(tbl2, func(x r.Term, y r.Term) interface{} { return x.AtIndex("a").Eq(y.AtIndex("b"))}).Zip()
	_ = ij // Prevent any noused variable errors


	{
		// joins.yaml line #37
		/* 2500 */
		var expected_ int = 2500
		/* ij.count() */

		suite.T().Log("About to run line #37: ij.Count()")

		runAndAssert(suite.Suite, expected_, ij.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// joins.yaml line #39
		/* 0 */
		var expected_ int = 0
		/* ij.filter(lambda row:row['a'] != row['b']).count() */

		suite.T().Log("About to run line #39: ij.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Ne(row.AtIndex('b'))}).Count()")

		runAndAssert(suite.Suite, expected_, ij.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Ne(row.AtIndex("b"))}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #39")
	}

	// joins.yaml line #46
	// oj = tbl.outer_join(tbl2, lambda x,y:x['a'] == y['b']).zip()
	suite.T().Log("Possibly executing: var oj r.Term = tbl.OuterJoin(tbl2, func(x r.Term, y r.Term) interface{} { return x.AtIndex('a').Eq(y.AtIndex('b'))}).Zip()")

	oj := tbl.OuterJoin(tbl2, func(x r.Term, y r.Term) interface{} { return x.AtIndex("a").Eq(y.AtIndex("b"))}).Zip()
	_ = oj // Prevent any noused variable errors


	{
		// joins.yaml line #49
		/* 2500 */
		var expected_ int = 2500
		/* oj.count() */

		suite.T().Log("About to run line #49: oj.Count()")

		runAndAssert(suite.Suite, expected_, oj.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #49")
	}

	{
		// joins.yaml line #51
		/* 0 */
		var expected_ int = 0
		/* oj.filter(lambda row:row['a'] != row['b']).count() */

		suite.T().Log("About to run line #51: oj.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Ne(row.AtIndex('b'))}).Count()")

		runAndAssert(suite.Suite, expected_, oj.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Ne(row.AtIndex("b"))}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #51")
	}

	// joins.yaml line #57
	// blah = otbl.order_by("id").eq_join(r.row['id'], otbl2, ordered=True).zip()
	suite.T().Log("Possibly executing: var blah r.Term = otbl.OrderBy('id').EqJoin(r.Row.AtIndex('id'), otbl2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip()")

	blah := maybeRun(otbl.OrderBy("id").EqJoin(r.Row.AtIndex("id"), otbl2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip(), suite.session, r.RunOpts{
	});
	_ = blah // Prevent any noused variable errors


	// joins.yaml line #59
	// blah = otbl.order_by(r.desc("id")).eq_join(r.row['id'], otbl2, ordered=True).zip()
	suite.T().Log("Possibly executing: var blah r.Term = otbl.OrderBy(r.Desc('id')).EqJoin(r.Row.AtIndex('id'), otbl2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip()")

	blah = maybeRun(otbl.OrderBy(r.Desc("id")).EqJoin(r.Row.AtIndex("id"), otbl2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip(), suite.session, r.RunOpts{
	});


	// joins.yaml line #61
	// blah = otbl.order_by("id").eq_join(r.row['a'], otbl2, ordered=True).zip()
	suite.T().Log("Possibly executing: var blah r.Term = otbl.OrderBy('id').EqJoin(r.Row.AtIndex('a'), otbl2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip()")

	blah = maybeRun(otbl.OrderBy("id").EqJoin(r.Row.AtIndex("a"), otbl2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip(), suite.session, r.RunOpts{
	});


	{
		// joins.yaml line #65
		/* 100 */
		var expected_ int = 100
		/* tbl.eq_join('a', tbl2).zip().count() */

		suite.T().Log("About to run line #65: tbl.EqJoin('a', tbl2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin("a", tbl2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #65")
	}

	{
		// joins.yaml line #68
		/* 0 */
		var expected_ int = 0
		/* tbl.eq_join('fake', tbl2).zip().count() */

		suite.T().Log("About to run line #68: tbl.EqJoin('fake', tbl2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin("fake", tbl2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #68")
	}

	{
		// joins.yaml line #71
		/* 100 */
		var expected_ int = 100
		/* tbl.eq_join(lambda x:x['a'], tbl2).zip().count() */

		suite.T().Log("About to run line #71: tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex('a')}, tbl2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex("a")}, tbl2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #71")
	}

	{
		// joins.yaml line #76
		/* 0 */
		var expected_ int = 0
		/* tbl.eq_join(lambda x:x['fake'], tbl2).zip().count() */

		suite.T().Log("About to run line #76: tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex('fake')}, tbl2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex("fake")}, tbl2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #76")
	}

	{
		// joins.yaml line #81
		/* 0 */
		var expected_ int = 0
		/* tbl.eq_join(lambda x:null, tbl2).zip().count() */

		suite.T().Log("About to run line #81: tbl.EqJoin(func(x r.Term) interface{} { return nil}, tbl2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin(func(x r.Term) interface{} { return nil}, tbl2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #81")
	}

	{
		// joins.yaml line #86
		/* 100 */
		var expected_ int = 100
		/* tbl.eq_join(lambda x:x['a'], tbl2).count() */

		suite.T().Log("About to run line #86: tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex('a')}, tbl2).Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex("a")}, tbl2).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #86")
	}

	{
		// joins.yaml line #92
		/* 100 */
		var expected_ int = 100
		/* tbl.eq_join('a', tbl3).zip().count() */

		suite.T().Log("About to run line #92: tbl.EqJoin('a', tbl3).Zip().Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin("a", tbl3).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #92")
	}

	{
		// joins.yaml line #95
		/* 100 */
		var expected_ int = 100
		/* tbl.eq_join(lambda x:x['a'], tbl3).count() */

		suite.T().Log("About to run line #95: tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex('a')}, tbl3).Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin(func(x r.Term) interface{} { return x.AtIndex("a")}, tbl3).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #95")
	}

	{
		// joins.yaml line #101
		/* 100 */
		var expected_ int = 100
		/* tbl.eq_join(r.row['a'], tbl2).count() */

		suite.T().Log("About to run line #101: tbl.EqJoin(r.Row.AtIndex('a'), tbl2).Count()")

		runAndAssert(suite.Suite, expected_, tbl.EqJoin(r.Row.AtIndex("a"), tbl2).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #101")
	}

	// joins.yaml line #106
	// left = r.expr([{'a':1},{'a':2},{'a':3}])
	suite.T().Log("Possibly executing: var left r.Term = r.Expr([]interface{}{map[interface{}]interface{}{'a': 1, }, map[interface{}]interface{}{'a': 2, }, map[interface{}]interface{}{'a': 3, }})")

	left := r.Expr([]interface{}{map[interface{}]interface{}{"a": 1, }, map[interface{}]interface{}{"a": 2, }, map[interface{}]interface{}{"a": 3, }})
	_ = left // Prevent any noused variable errors


	// joins.yaml line #107
	// right = r.expr([{'b':2},{'b':3}])
	suite.T().Log("Possibly executing: var right r.Term = r.Expr([]interface{}{map[interface{}]interface{}{'b': 2, }, map[interface{}]interface{}{'b': 3, }})")

	right := r.Expr([]interface{}{map[interface{}]interface{}{"b": 2, }, map[interface{}]interface{}{"b": 3, }})
	_ = right // Prevent any noused variable errors


	{
		// joins.yaml line #109
		/* [{'a':2,'b':2},{'a':3,'b':3}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": 2, "b": 2, }, map[interface{}]interface{}{"a": 3, "b": 3, }}
		/* left.inner_join(right, lambda l, r:l['a'] == r['b']).zip() */

		suite.T().Log("About to run line #109: left.InnerJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex('a').Eq(r.AtIndex('b'))}).Zip()")

		runAndAssert(suite.Suite, expected_, left.InnerJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex("a").Eq(r.AtIndex("b"))}).Zip(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #109")
	}

	{
		// joins.yaml line #115
		/* [{'a':1},{'a':2,'b':2},{'a':3,'b':3}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": 1, }, map[interface{}]interface{}{"a": 2, "b": 2, }, map[interface{}]interface{}{"a": 3, "b": 3, }}
		/* left.outer_join(right, lambda l, r:l['a'] == r['b']).zip() */

		suite.T().Log("About to run line #115: left.OuterJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex('a').Eq(r.AtIndex('b'))}).Zip()")

		runAndAssert(suite.Suite, expected_, left.OuterJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex("a").Eq(r.AtIndex("b"))}).Zip(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #115")
	}

	{
		// joins.yaml line #132
		/* partial({'tables_dropped':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_dropped": 1, })
		/* r.db('test').table_drop('test3') */

		suite.T().Log("About to run line #132: r.DB('test').TableDrop('test3')")

		runAndAssert(suite.Suite, expected_, r.DB("test").TableDrop("test3"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat: "map",
		})
		suite.T().Log("Finished running line #132")
	}
}
