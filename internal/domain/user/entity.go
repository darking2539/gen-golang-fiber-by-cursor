package user

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"uniqueIndex" json:"username"`
	Password  string `json:"-"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Bio       string `json:"bio"`
	Title     string `json:"title"`
	Location  string `json:"location"`
	Website   string `json:"website"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Posts     int    `json:"posts"`
}
