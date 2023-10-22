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

// function for user
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

// function for plan schedule
func getPlanSchedule(c *gin.Context) {
	var planSchedule []PlanSchedule
	db.Find(&planSchedule)
	c.IndentedJSON(http.StatusOK, planSchedule)
}
func postPlanSchedule(c *gin.Context) {
	var newPlanSchedule PlanSchedule
	if err := c.BindJSON(&newPlanSchedule); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error to add plan schedule"})
		return
	}
	db.Create(&newPlanSchedule)
	c.IndentedJSON(http.StatusCreated, newPlanSchedule)
}
func getPlanScheduleByDate(c *gin.Context) {
	dateStart := c.Param("start")
	dateEnd := c.Param("end")
	var planSchedule []PlanSchedule
	db.Where("date >= ? AND date <= ?", dateStart, dateEnd).Find(&planSchedule)
	c.IndentedJSON(http.StatusOK, planSchedule)
}
func updatePlanSchedule(c *gin.Context) {
	var planSchedule PlanSchedule
	if err := c.BindJSON(&planSchedule); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error to update plan schedule"})
		return
	}
	db.Save(&planSchedule)
	c.IndentedJSON(http.StatusOK, planSchedule)
}

// function for holiday user
func getHolidayUser(c *gin.Context) {
	var holidayUser []HolidayUser
	db.Find(&holidayUser)
	c.IndentedJSON(http.StatusOK, holidayUser)
}
func postHolidayUser(c *gin.Context) {
	var newHolidayUser HolidayUser
	if err := c.BindJSON(&newHolidayUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error to add holiday user"})
		return
	}
	db.Create(&newHolidayUser)
	c.IndentedJSON(http.StatusCreated, newHolidayUser)
}
func getHolidayByUser(c *gin.Context) {
	id := c.Param("id")
	var holidayUser []HolidayUser
	db.Where("user_id = ?", id).Find(&holidayUser)
	c.IndentedJSON(http.StatusOK, holidayUser)
}
func removeHolidayUser(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&HolidayUser{}, id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "holiday deleted"})
}

// function for plan time
func getPlanTime(c *gin.Context) {
	var planTime []PlanTime
	db.Find(&planTime)
	c.IndentedJSON(http.StatusOK, planTime)
}
func postPlanTime(c *gin.Context) {
	var newPlanTime PlanTime
	if err := c.BindJSON(&newPlanTime); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error to add plan time"})
		return
	}
	db.Create(&newPlanTime)
	c.IndentedJSON(http.StatusCreated, newPlanTime)
}
func updatePlanTime(c *gin.Context) {
	var planTime PlanTime
	if err := c.BindJSON(&planTime); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error to update plan time"})
		return
	}
	db.Save(&planTime)
	c.IndentedJSON(http.StatusOK, planTime)
}
func getPlanTimeByUser(c *gin.Context) {
	id := c.Param("id")
	var planTime []PlanTime
	db.Where("user_id = ?", id).Find(&planTime)
	c.IndentedJSON(http.StatusOK, planTime)
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

	route.GET("/planSchedule", getPlanSchedule)
	route.POST("/planSchedule", postPlanSchedule)
	route.GET("/planSchedule/:start/:end", getPlanScheduleByDate)
	route.PUT("/planSchedule", updatePlanSchedule)

	route.GET("/holidayUser", getHolidayUser)
	route.POST("/holidayUser", postHolidayUser)
	route.GET("/holidayUser/:id", getHolidayByUser)
	route.DELETE("/holidayUser/:id", removeHolidayUser)

	route.GET("/planTime", getPlanTime)
	route.POST("/planTime", postPlanTime)
	route.PUT("/planTime", updatePlanTime)
	route.GET("/planTime/:id", getPlanTimeByUser)

	err = route.Run(":8080")
	if err != nil {
		return
	}
}
