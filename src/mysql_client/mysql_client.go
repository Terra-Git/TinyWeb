package mysql_client

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnectIfo struct{
	Username_          string
	Password_          string
	Dp_ip_             string
	Dp_port_           string
	Dbname_            string
	Max_connect_       int 
}

// mysql 值类型
type MysqlValueType int32

const (
    MYSQL_INT         MysqlValueType = 0
    MYSQL_STRING      MysqlValueType = 1
)

// mysql 键值对数据
type MysqlColData struct{
	column_  string
	value_   string
	type_    MysqlValueType
}

// mysql 的行数据
type MysqlRowData struct{
	table_name_   string
	col_count_    int
	row_data_     []MysqlColData
}

// 实现增删改查
type MysqlClient struct{
	connect_     *sql.DB
	mysql_user_  MysqlConnectIfo
}

func (client *MysqlClient) Connet(username string,password string,ip string,port string,dbName string,max_connect int){
	client.connect_,_ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, ip, port, dbName))
	//设置上数据库最大闲置连接数
	client.connect_.SetMaxIdleConns(max_connect)
}

func (client *MysqlClient) Connect( user MysqlConnectIfo){
	client.mysql_user_ = user
	client.connect_,_= sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",user.Username_,user.Password_,user.Dp_ip_,user.Dp_port_,user.Dbname_))
	//设置数据库最大连接数
	client.connect_.SetMaxIdleConns(user.Max_connect_)
}

func (client *MysqlClient) Insert_record(table_name string, data []MysqlRowData){

}

func (client *MysqlClient) Delete_record(){

}

func (client *MysqlClient) Update_record(){

}

func (client *MysqlClient) Query_record(sql string){

}

// 组装一行的sql,注意增加空格
func build_insert_sql(table_name string, data []MysqlRowData) string {
	var s_columns_name string
	var s_columns_value string

	// 组装列
	s_columns_name += " ("
	for index,col := range data[0].row_data_{
		if index == 0 {
			s_columns_name += "`" + col.column_ + "`"
		}else{
			s_columns_name += ",`" + col.column_ + "`"
		}
	}
	s_columns_name += ")"

	s_columns_value += "("
	for i := 0; i < len(data); i++ {
		for index,col := range data[i].row_data_{
			if index != 0{
				s_columns_value += ","
			}
			if col.type_ == MYSQL_INT {
				s_columns_value += col.value_ 
			}else{
				s_columns_value += "'" + col.value_ + "'" 
			}
		}
		if i != len(data) - 1 {
			s_columns_value += ","
		}
	}
	s_columns_value += ")"

	var sql string = "insert into " + table_name + s_columns_name + " values " + s_columns_value + ";"
	return sql
}

func (client *MysqlClient) Test(){

	//验证连接
	if err := client.connect_.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

	var test1 MysqlColData = MysqlColData{"name","gyk",MYSQL_STRING}
	var test2 MysqlColData = MysqlColData{"phone","123",MYSQL_INT}
	var data MysqlRowData = MysqlRowData{"student",2,nil}
	data.row_data_ = append(data.row_data_ , test1)
	data.row_data_ = append(data.row_data_ , test2)
	var datas []MysqlRowData
	datas = append(datas,data)
	datas = append(datas,data)

	str := build_insert_sql("student",datas);
	print(str,"\n")

}