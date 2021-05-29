package model

// +--------+--------------+------+-----+---------+----------------+
// | Field  | Type         | Null | Key | Default | Extra          |
// +--------+--------------+------+-----+---------+----------------+
// | id     | int(11)      | NO   | PRI | NULL    | auto_increment |
// | name   | varchar(32)  | YES  |     | NULL    |                |
// | price  | float        | YES  |     | NULL    |                |
// | store  | int(11)      | YES  |     | NULL    |                |
// | detail | varchar(256) | YES  |     | NULL    |                |
// +--------+--------------+------+-----+---------+----------------+

type Product struct {
	ID        int     `gorm:"column:id;primary_key"`
	Name      string  `gorm:"column:name;type:varchar(32)"`
	Price     float64 `gorm:"column:price;type:float"`
	Store     int     `gorm:"column:store;type:int(11)"`
	Detail    string  `gorm:"column:name;type:varchar(256)"`
}
