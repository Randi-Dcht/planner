package main

//list of import
import (
	"github.com/gin-gonic/gin"

	"net/http"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

type User struct {
	gorm.Model
	ID    string
	Name  string
	Email string
	Phone string
}
type PlanTime struct {
	gorm.Model
	ID        string
	UserID    string
	Date      string
	TimeStart string
	TimeEnd   string
}
type PlanSchedule struct {
	gorm.Model
	ID       string
	Date     string
	DayStart string
	DayEnd   string
	Remarque string
}
type HolidayUser struct {
	gorm.Model
	ID          string
	UserID      string
	Date        string
	HolidayName string
}

var db *gorm.DB
var err error

func getUser(c *gin.Context) {
	var user []User
	db.Find(&user)
	c.IndentedJSON(http.StatusOK, user)
}

func postUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error to add user"})
		return
	}
	db.Create(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&User{}, id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func main() {
	// Create connection
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&PlanTime{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&PlanSchedule{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&HolidayUser{})
	if err != nil {
		return
	}

	route := gin.Default()
	route.GET("/user", getUser)
	route.POST("/user", postUser)
	route.DELETE("/user/:id", deleteUser)

	route.Run(":8080")
}
