package repository

import (
	// "database/sql"
	"checkapp/db"
	"checkapp/model"
)

func GetAllMemos()([]model.Memo,error){
	rows,err:= db.Conn.Query("SELECT id,content FROM memos")
	if err != nil{
		return nil,err
	}
	defer rows.Close()

  var memos []model.Memo
	for rows.Next(){
		var m model.Memo
		if err := rows.Scan(&m.ID,&m.Content);err != nil{
			return nil,err
		}
		memos = append(memos,m)
	 }
	 return memos,nil
}

func InsertMemo(m model.Memo) error {
_,err := db.Conn.Exec("INSERT INTO memos (content) VALUES (?)",m.Content)
return err
}

func DeleteMemo(id string) error{
	_,err := db.Conn.Exec("DELETE FROM memos WHERE id=?",id)
	return err
}