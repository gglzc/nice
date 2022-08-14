package account

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct{
	ID        		primitive.ObjectID     `bson:"id"`
	UserID			*string				   `json:"userid"    validate:"required,min=6"` 
	Password  		*string                `json:"password"  validate:"required,min=6"` 	
	Lastname  		*string 			   `json:"lastname"  validate:"required min=2,max=50"`	
	Firstname 		*string 			   `json:"firstname" validate:"required min=2,max=50"`
	Gender          *string 			   `json:"gender"    validate:"required"`   
	Birth	  		*string                `json:"birth"     validate:"required"`   
	Email     		*string                `json:"email"     validate:"email,required"`
	Address   		*string                `json:"address"`
	User_Type 		*string                `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`	 
	Token     		*string                `json:"token"`
	Refresh_Token   *string                `json:refresh_token`
	CreateTime      time.Time
	UpdateTime      time.Time
}

//存User資料
var user_info []User

