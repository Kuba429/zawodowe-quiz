package zdjecie

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func Usun(db *sql.DB, id int, kat string) {
	query := fmt.Sprintf("SELECT Obrazek from %s WHERE Id=%d;", kat, id)
	row := db.QueryRow(query)
	var res string
	row.Scan(&res)

	split := strings.Split(res, "/")
	split = split[1:]
	if len(split) != 3 || split[0] != FolderOgolny || split[1] != FolderUzytkPrefix || len(split[2]) < 1 {
		return
	}
	os.Remove(fmt.Sprintf("%s/%s", Folder, split[2]))
}
