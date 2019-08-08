package main
import(
	_ "github.com/go-sql-driver/mysql"
		"database/sql"
)

func main(){
	db, err := sql.Open("mysql", "root:Scode1998.@tcp(124.156.206.94:3306)/people?charset=utf8")
	if err != nil {
		panic(err)
	}
	TestLink(db)
	Insert(db,42,"Gale")
	TestLink(db)
	Delete(db,21)
	TestLink(db)
	Update(db,23,"Karr")
	TestLink(db)
	Search(db,32)
	SearchInRange(db,21,32)
	CloseDataBase(db)
}
