package model

import (
	"time"
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
