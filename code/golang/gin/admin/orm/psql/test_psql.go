package psql

import (
	"database/sql"
	"fmt"
	_ "go.uber.org/automaxprocs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.7.2
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /

type User struct {
	gorm.Model
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     time.Time      // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

func connection() (db *gorm.DB, error error) {

	dsn := "host=localhost user=postgres password=123456 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败")
		panic(err)
	}
	fmt.Println(db)

	return db, nil
}

func table_create() {
	// init table creating

	db, err := connection()
	if err != nil {
		log.Fatal("连接失败")
	}
	db.AutoMigrate(&User{}, &Blog{}, &Author{})
}

func save() {
	db, err := connection()
	if err != nil {
		log.Fatal("连接失败")
	}
	user := User{Name: "laios", Age: 23, Birthday: time.Now()}
	res := db.Create(&user)

	fmt.Println(user.ID)
	fmt.Println(res.RowsAffected)
	fmt.Println(res.Error)

	user.Age = 24
	db.Save(&user)
	fmt.Println(user)

	// batch save
	users := []*User{
		{Name: "laios1"},
		{Name: "laios2"},
	}
	db.Save(&users)

	// batch save specify batchSize
	db.CreateInBatches(users, 100)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	fmt.Println("test hook")
	return
}
func (u *User) String() string {
	str := fmt.Sprintf(" %s  %s %s ", u.ID, u.Name, u.Age)
	return str
}

func createByMap() {
	db, err := connection()
	if err != nil {
		log.Fatal("连接失败")
	}

	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "laios_by_map",
	})

	res := db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
	fmt.Println(res.Error)
}

func byId(db gorm.DB) {
	var user User
	var users []User

	db.Find(&user, "1")
	db.Find(&users, []int{1, 2, 3})
	fmt.Println(user)
	fmt.Println(users)

}

func where(db gorm.DB) {
	var user User
	var users []User

	db.Where("name = ?", "laios").First(&user)
	fmt.Println(user, " \n\n")

	db.Where("id in ?", []int{1, 2, 3, 4}).Find(&users)
	fmt.Println(users)

}

func update(db gorm.DB) {
	db.Model(&User{}).Update("name", "laios")

	db.Model(&User{}).Updates(User{Name: "laios", Age: 23})
}

func main() {

	//table_create()
	db, err := connection()
	if err != nil {
		panic(err)
	}
	//save()
	//createByMap()
	//byId(*db)
	//where(*db)
	update(*db)
	//global.GVA_VP = core.Viper() // 初始化Viper
	//initialize.OtherInit()
	//global.GVA_LOG = core.Zap() // 初始化zap日志库
	//zap.ReplaceGlobals(global.GVA_LOG)
	//global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//initialize.Timer()
	//initialize.DBList()
	//if global.GVA_DB != nil {
	//	initialize.RegisterTables() // 初始化表
	//	// 程序结束前关闭数据库链接
	//	db, _ := global.GVA_DB.DB()
	//	defer db.Close()
	//}
	//core.RunWindowsServer()

}
