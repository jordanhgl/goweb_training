package dbtest

import (
  "testing"
)

func TestConnDb(t *testing.T) {
  ConnDb()
  t.Log("test over")

  affectrow := InsertRecord(Db, User{3, "h&k", 19})
  t.Log("affectrow = ", affectrow)

}
