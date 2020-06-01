type auth_user struct {
	TableName() string 
	id 			int64
	username 	string 		`xorm:"varchar(64)"`
	password 	string 		`xorm:"varchar(255)"`
	user_type 	int64
	real_name 	string 		`xorm:"varchar(64)"`
	mail 		string 		`xorm:"varchar(64)"`
	mobile 		string 		`xorm:"varchar(16)"`
	status 		int64
	description string 		`xorm:"varchar(255)"`
	creator 	string 		`xorm:"varchar(64)"`
	create_time time.Time 	`xorm:"created"`
	update_time time.Time 	`xorm:"updated"`
}