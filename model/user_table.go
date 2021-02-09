package model

import (
	"context"
	"errors"
	"fmt"
	"log"
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

type UserInfoBind struct {
	User_id           int64     `json:"u_user_id" binding:"required"`
	Account_id        int64     `json:"u_account_id" binding:"required"`
	Contact_id        int64     `json:"u_contact_id" binding:"required"`
	Loyalty_id        int64     `json:"u_loyalty_id" binding:"required"`
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

func GetSingleUser(c *gin.Context) (*UserInfo, error) {
	var uinfo *UserInfo
	conn := db.Postgres

	user_id := c.Query("user_id")
	// account_id := c.Query("account_id")
	// contact_id := c.Query("contact_id")
	// loyalty_id := c.Query("loyalty_id")

	// fmt.Printf("user_id: %s; account_id: %s; contact_id: %s; loyalty_id: %s \n", user_id, account_id, contact_id, loyalty_id)
	fmt.Printf("user_id: %s\n", user_id)
	if user_id == "" {
		return nil, errors.New("all mandatory values NOT Passed")
	}

	rows, err := conn.Query(context.Background(), "SELECT * FROM user_info WHERE u_user_id=$1 LIMIT 1", user_id)
	if err != nil {
		// log.Fatal("error while executing query")
		return nil, err
	}

	count := 0
	// iterate through the rows
	for rows.Next() {
		count++
		values, err := rows.Values()
		if err != nil {
			// log.Fatal("error while iterating dataset")
			return nil, err
		}

		// convert DB types to Go types
		uinfo = &UserInfo{
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
		log.Println("[User_id:", uinfo.User_id, "]")
	}

	// return uinfo, nil
	if count > 0 {
		return uinfo, nil
	} else {
		return nil, errors.New("record NOT Found")
	}
}

func InsertSingleUser(c *gin.Context) error {
	var uinfo UserInfoBind
	conn := db.Postgres

	if err := c.ShouldBindJSON(&uinfo); err != nil {
		log.Printf("\n InPut Details: ERROR - shouldbindjson() : %v\n\n", err.Error())
		return err
	}

	log.Printf("\n InPut Details: %+v \n\n", uinfo)

	_, err := conn.Exec(context.Background(),
		"INSERT INTO user_info (\"u_user_id\", \"u_account_id\", \"u_contact_id\", \"u_loyalty_id\", \"u_is_active_id\", \"u_reference_id\", \"u_user_type\", \"u_account_type\", \"u_loyalty_type\", \"u_member_type\", \"u_brand_type\", \"u_create_rcd_at\", \"u_create_rcd_by_who\", \"u_create_rcd_by_app\", \"u_update_rcd_at\", \"u_update_rcd_by_who\", \"u_update_rcd_by_app\", \"u_data_source\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)",
		uinfo.User_id, uinfo.Account_id, uinfo.Contact_id, uinfo.Loyalty_id, uinfo.IsActive_id, uinfo.Reference_id, uinfo.User_type, uinfo.Account_type, uinfo.Loyalty_type, uinfo.Member_type, uinfo.Brand_type, uinfo.Create_rcd_at, uinfo.Create_rcd_by_who, uinfo.Create_rcd_by_app, uinfo.Update_rcd_at, uinfo.Update_rcd_by_who, uinfo.Update_rcd_by_app, uinfo.Data_source)

	if err != nil {
		log.Println("error while executing query: ", err.Error())
		// return errors.New("error while executing INSERT query")
		return err
	}

	log.Printf("\n INSERT SUCCESS \n\n")

	return nil
}

func UpdateSingleUser(c *gin.Context) error {
	eval, err := GetSingleUser(c)
	if err != nil {
		return err
	}
	if eval == nil {
		return errors.New("record NOT Found")
	}

	var uinfo UserInfoBind
	conn := db.Postgres

	if err := c.ShouldBindJSON(&uinfo); err != nil {
		log.Printf("\n InPut Details: ERROR - bindjson() : %v\n\n", err.Error())
		return err
	}

	log.Printf("\n InPut Details: %+v \n\n", uinfo)

	user_id := c.Query("user_id")
	// account_id := c.Query("account_id")
	// contact_id := c.Query("contact_id")
	// loyalty_id := c.Query("loyalty_id")

	// fmt.Printf("user_id: %s; account_id: %s; contact_id: %s; loyalty_id: %s \n", user_id, account_id, contact_id, loyalty_id)
	fmt.Printf("user_id: %s\n", user_id)
	if user_id == "" {
		return errors.New("all mandatory values NOT Passed")
	}

	_, err = conn.Exec(context.Background(),
		"UPDATE user_info SET \"u_user_id\"=$1, \"u_account_id\"=$2, \"u_contact_id\"=$3, \"u_loyalty_id\"=$4, \"u_is_active_id\"=$5, \"u_reference_id\"=$6, \"u_user_type\"=$7, \"u_account_type\"=$8, \"u_loyalty_type\"=$9, \"u_member_type\"=$10, \"u_brand_type\"=$11, \"u_create_rcd_at\"=$12, \"u_create_rcd_by_who\"=$13, \"u_create_rcd_by_app\"=$14, \"u_update_rcd_at\"=$15, \"u_update_rcd_by_who\"=$16, \"u_update_rcd_by_app\"=$17, \"u_data_source\"=$18 WHERE u_user_id=$19",
		uinfo.User_id, uinfo.Account_id, uinfo.Contact_id, uinfo.Loyalty_id, uinfo.IsActive_id, uinfo.Reference_id, uinfo.User_type, uinfo.Account_type, uinfo.Loyalty_type, uinfo.Member_type, uinfo.Brand_type, uinfo.Create_rcd_at, uinfo.Create_rcd_by_who, uinfo.Create_rcd_by_app, uinfo.Update_rcd_at, uinfo.Update_rcd_by_who, uinfo.Update_rcd_by_app, uinfo.Data_source, user_id)
	if err != nil {
		log.Println("error while executing query: ", err.Error())
		return errors.New("error while executing UPDATE query")
	}

	log.Printf("\n UPDATE SUCCESS \n\n")

	return nil
}

func DeleteSingleUser(c *gin.Context) error {
	conn := db.Postgres

	user_id := c.Query("user_id")
	// account_id := c.Query("account_id")
	// contact_id := c.Query("contact_id")
	// loyalty_id := c.Query("loyalty_id")

	// fmt.Printf("user_id: %s; account_id: %s; contact_id: %s; loyalty_id: %s \n", user_id, account_id, contact_id, loyalty_id)
	fmt.Printf("user_id: %s\n", user_id)

	if user_id == "" {
		return errors.New("all mandatory values NOT Passed")
	}

	_, err := conn.Exec(context.Background(), "DELETE FROM user_info WHERE u_user_id=$1", user_id)
	if err != nil {
		log.Println("error while executing query: ", err.Error())
		return errors.New("error while executing DELETE query")
	}

	return err
}
