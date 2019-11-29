package model

// 構造体
// メタ情報でカラム名を指定する事でsqlxがSQLと構造体を紐付ける
type coin struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Price int    `db:"price"`
}
