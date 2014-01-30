package db

import (
    "fmt"
    "github.com/public0821/dnserver/errors"
)

type User struct {
    Name string
    Pwd  string
}

func (m *User) Query(dbm *DBManager, start, offset int) (records []interface{}, err error) {
    if start < 0 || offset < 0 {
        err = errors.New("invalid arguments")
    }
    sql := "select name, pwd from user "
    if len(m.Name) != 0 {
        sql += fmt.Sprintf(" where name='%s'", m.Name)
    }
    rows, err := dbm.db.Query(sql)
    if err != nil {
        return
    }
    defer rows.Close()
    for rows.Next() {
        var user User
        rows.Scan(&user.Name, &user.Pwd)
        records = append(records, user)
    }

    return
}

func (m *User) Modify(dbm *DBManager) (err error) {
    sql := fmt.Sprintf("update user set pwd='%s' where name='%s'",
        m.Pwd, m.Name)
    _, err = dbm.db.Exec(sql)
    return
}

func (m *User) DeleteAll(dbm *DBManager) (err error) {
    err = errors.New("unimplemented")
    return
}

func (m *User) Delete(dbm *DBManager) (err error) {
    err = errors.New("unimplemented")
    return
}

func (m *User) Count(dbm *DBManager) (count int, err error) {
    sql := "select count(*) from user "
    rows, err := dbm.db.Query(sql)
    if err != nil {
        return
    }
    defer rows.Close()
    for rows.Next() {
        rows.Scan(&count)
        return
    }
    return
}

func (m *User) Add(dbm *DBManager) (err error) {
    tx, err := dbm.db.Begin()
    if err != nil {
        return
    }
    stmt, err := tx.Prepare(`insert into user(name,  pwd)
        values (?,?)`)
    if err != nil {
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec(m.Name, m.Pwd)
    if err != nil {
        return
    }
    tx.Commit()

    return
}
