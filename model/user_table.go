package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pksingh/gin-curd-demo/startup/db"
)

type UserInfo struct {
	User_id           int64     `json:"u_user_id"`
	Account_id        int64     `json:"u_account_id"`
	Contact_id        int64     `json:"u_contact_id"`
	Loyalty_id        int64     `json:"u_loyalty_id"`
	IsActive_id       bool      `json:"u_is_active_id"`
	Reference_id      int64     `json:"u_reference_id"`
	User_type         string    `json:"u_user_type"`
	Account_type      string    `json:"u_account_type"`
	Loyalty_type      string    `json:"u_loyalty_type"`
	Member_type       string    `json:"u_member_type"`
	Brand_type        string    `json:"u_brand_type"`
	Create_rcd_at     time.Time `json:"u_create_rcd_at"`
	Create_rcd_by_who string    `json:"u_create_rcd_by_who"`
	Create_rcd_by_app string    `json:"u_create_rcd_by_app"`
	Update_rcd_at     time.Time `json:"u_update_rcd_at"`
	Update_rcd_by_who string    `json:"u_update_rcd_by_who"`
	Update_rcd_by_app string    `json:"u_update_rcd_by_app"`
	Data_source       string    `json:"u_data_source"`
}

func GetAllUsers(c *gin.Context) ([]UserInfo, error) {
	var uinfos []UserInfo
	conn := db.Postgres

	query := "SELECT * FROM user_info"
	fmt.Println("Query: ", query)
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	count := 0
	// iterate through the rows
	for rows.Next() {
		count++
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		// convert DB types to Go types
		uinfo := UserInfo{
			User_id:           values[0].(int64),
			Account_id:        values[1].(int64),
			Contact_id:        values[2].(int64),
			Loyalty_id:        values[3].(int64),
			IsActive_id:       values[4].(bool),
			Reference_id:      values[5].(int64),
			User_type:         values[6].(string),
			Account_type:      values[7].(string),
			Loyalty_type:      values[8].(string),
			Member_type:       values[9].(string),
			Brand_type:        values[10].(string),
			Create_rcd_at:     values[11].(time.Time),
			Create_rcd_by_who: values[12].(string),
			Create_rcd_by_app: values[13].(string),
			Update_rcd_at:     values[14].(time.Time),
			Update_rcd_by_who: values[15].(string),
			Update_rcd_by_app: values[16].(string),
			Data_source:       values[17].(string),
		}

		uinfos = append(uinfos, uinfo)
	}

	if count > 0 {
		return uinfos, nil
	} else {
		return nil, errors.New("record NOT Found")
	}
}
