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
	Column_  string
	Value_   string
	Type_    MysqlValueType
}

// 实现增删改查
type MysqlClient struct{
	connect_     *sql.DB
	mysql_user_  MysqlConnectIfo
}

func (this *MysqlClient) Connet(username string,password string,ip string,port string,dbName string,max_connect int){
	this.mysql_user_ = MysqlConnectIfo{username,password,ip,port,dbName,max_connect}
	this.connect_,_ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, ip, port, dbName))
	//设置上数据库最大闲置连接数
	this.connect_.SetMaxIdleConns(max_connect)
}

func (this *MysqlClient) Connect( user MysqlConnectIfo){
	this.mysql_user_ = user
	this.connect_,_ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",user.Username_,user.Password_,user.Dp_ip_,user.Dp_port_,user.Dbname_))
	//设置数据库最大连接数
	this.connect_.SetMaxIdleConns(user.Max_connect_)
}

func (this *MysqlClient) Insert_record(table_name string, data []MysqlColData) int {
	sql := this.build_insert_sql(table_name,data)
	return this.execute(sql)
}

func (this *MysqlClient) Insert_batch_record(table_name string, data [][]MysqlColData) int {
	sql := this.build_batch_insert_sql(table_name,data)
	return this.execute(sql)
}

func (this *MysqlClient) Delete_record(table_name string, colunm_name string, column_value string ,value_type MysqlValueType) int {
	var sql string = this.build_delete_sql(table_name,colunm_name,column_value,value_type)
	return this.execute(sql)
}

func (this *MysqlClient) Update_record(){

}

func (this *MysqlClient) Query_record(sql string){

}


// 组装一行的sql,注意增加空格
func (this *MysqlClient) build_insert_sql(table_name string, data []MysqlColData) string {
	var s_columns_name string
	var s_columns_value string

	s_columns_name += " ("
	for index,col := range data {
		if index == 0 {
			s_columns_name += "`" + col.Column_ + "`"
		}else{
			s_columns_name += ",`" + col.Column_ + "`"
		}
	}
	s_columns_name += ")"

	s_columns_value += "("
	for index,col := range data {
		if index != 0{
			s_columns_value += ","
		}
		if col.Type_ == MYSQL_INT {
			s_columns_value += col.Value_ 
		}else{
			s_columns_value += "'" + col.Value_ + "'" 
		}
	}
	s_columns_value += ")"

	var sql string = "insert into " + table_name + s_columns_name + " values " + s_columns_value + ";"
	return sql
}

// 组装一行的sql,注意增加空格
func (this *MysqlClient) build_batch_insert_sql(table_name string, data [][]MysqlColData) string {
	var s_columns_name string
	var s_columns_value string

	// 组装列
	s_columns_name += " ("
	for index,col := range data[0] {
		if index == 0 {
			s_columns_name += "`" + col.Column_ + "`"
		}else{
			s_columns_name += ",`" + col.Column_ + "`"
		}
	}
	s_columns_name += ")"

	for i := 0; i < len(data); i++ {
		s_columns_value += "("
		for index,col := range data[i] {
			if index != 0{
				s_columns_value += ","
			}
			if col.Type_ == MYSQL_INT {
				s_columns_value += col.Value_ 
			}else{
				s_columns_value += "'" + col.Value_ + "'" 
			}
		}
		s_columns_value += ")"
		if i != len(data) - 1 {
			s_columns_value += ","
		}
	}

	var sql string = "insert into " + table_name + s_columns_name + " values " + s_columns_value + ";"
	return sql
}

func (this *MysqlClient) build_delete_sql(table_name string, colunm_name string, column_value string ,value_type MysqlValueType) string {
	var value string
	if( value_type == MYSQL_INT){
		value = column_value
	}else{
		value = "'" + column_value + "'"
	}
	var sql string = "delete from `" + table_name + "` where `" + colunm_name + "` =" + value + ";"
	return sql;
}

// 执行插入语句，返回 0 为成功
func (this *MysqlClient) execute(sql string) int {

	r, err := this.connect_.Exec(sql)
	print(r)
	if err != nil {
		fmt.Println("exec failed,", err)
		return 1
	}
	fmt.Println("insert succ")
	return 0
}

func (this *MysqlClient) Test_insert(){

	//验证连接
	if err := this.connect_.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

	var test1 MysqlColData = MysqlColData{"name","gyk",MYSQL_STRING}
	var test2 MysqlColData = MysqlColData{"phone","123",MYSQL_INT}
	var test3 MysqlColData = MysqlColData{"address","hunan",MYSQL_STRING}

	var data []MysqlColData

	data = append(data , test1)
	data = append(data , test2)
	data = append(data , test3)

	var datas [][]MysqlColData

	datas = append(datas,data)
	datas = append(datas,data)

	// this.Insert_batch_record("t_note_user",datas);

	this.Insert_record("t_note_user",data)

	// str := build_batch_insert_sql("student",datas);
	// print(str,"\n")
	// str = build_insert_sql("student",data)
	// print(str,"\n")
}