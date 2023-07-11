// Code generated by sqlc. DO NOT EDIT.
// versions:
//
//	sqlc v1.19.0
package users

import (
	"database/sql"
	"time"
)

type UsersBusinessEntity struct {
	EntityID int32 `db:"entity_id" json:"entityId"`
}

type UsersPhoneNumberType struct {
	PontyCode         string       `db:"ponty_code" json:"pontyCode"`
	PontyModifiedDate time.Time `db:"ponty_modified_date" json:"pontyModifiedDate"`
}

type UsersRole struct {
	RoleID           int32          `db:"role_id" json:"roleId"`
	RoleName         string `db:"role_name" json:"roleName"`
	RoleType         string `db:"role_type" json:"roleType"`
	RoleModifiedDate time.Time   `db:"role_modified_date" json:"roleModifiedDate"`
}

type UsersUser struct {
	UserEntityID       int32                 `db:"user_entity_id" json:"userEntityId"`
	UserName           string        `db:"user_name" json:"userName"`
	UserPassword       string        `db:"user_password" json:"userPassword"`
	UserFirstName      string        `db:"user_first_name" json:"userFirstName"`
	UserLastName       string        `db:"user_last_name" json:"userLastName"`
	UserBirthDate      time.Time          `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion int64         `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    sql.NullString `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   time.Time          `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          string        `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    int64         `db:"user_current_role" json:"userCurrentRole"`
}

type UsersUsersAddress struct {
	EtadAddrID       int32         `db:"etad_addr_id" json:"etadAddrId"`
	EtadModifiedDate time.Time  `db:"etad_modified_date" json:"etadModifiedDate"`
	EtadEntityID     int64 `db:"etad_entity_id" json:"etadEntityId"`
	EtadAdtyID       int64 `db:"etad_adty_id" json:"etadAdtyId"`
}

type UsersUsersEducation struct {
	UsduID           int32          `db:"usdu_id" json:"usduId"`
	UsduEntityID     int32          `db:"usdu_entity_id" json:"usduEntityId"`
	UsduSchool       string `db:"usdu_school" json:"usduSchool"`
	UsduDegree       string `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   string `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGraduateYear string `db:"usdu_graduate_year" json:"usduGraduateYear"`
	UsduStartDate    time.Time   `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      time.Time   `db:"usdu_end_date" json:"usduEndDate"`
	UsduGrade        string `db:"usdu_grade" json:"usduGrade"`
	UsduActivities   string `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  string `db:"usdu_description" json:"usduDescription"`
	UsduModifiedDate time.Time   `db:"usdu_modified_date" json:"usduModifiedDate"`
}

type UsersUsersEmail struct {
	PmailEntityID     int32          `db:"pmail_entity_id" json:"pmailEntityId"`
	PmailID           int32          `db:"pmail_id" json:"pmailId"`
	PmailAddress      string `db:"pmail_address" json:"pmailAddress"`
	PmailModifiedDate time.Time   `db:"pmail_modified_date" json:"pmailModifiedDate"`
}

type UsersUsersExperience struct {
	UsexID              int32          `db:"usex_id" json:"usexId"`
	UsexEntityID        int32          `db:"usex_entity_id" json:"usexEntityId"`
	UsexTitle           string `db:"usex_title" json:"usexTitle"`
	UsexProfileHeadline string `db:"usex_profile_headline" json:"usexProfileHeadline"`
	UsexEmploymentType  string `db:"usex_employment_type" json:"usexEmploymentType"`
	UsexCompanyName     string `db:"usex_company_name" json:"usexCompanyName"`
	UsexIsCurrent       string `db:"usex_is_current" json:"usexIsCurrent"`
	UsexStartDate       time.Time   `db:"usex_start_date" json:"usexStartDate"`
	UsexEndDate         time.Time   `db:"usex_end_date" json:"usexEndDate"`
	UsexIndustry        string `db:"usex_industry" json:"usexIndustry"`
	UsexDescription     string `db:"usex_description" json:"usexDescription"`
	UsexExperienceType  string `db:"usex_experience_type" json:"usexExperienceType"`
	UsexCityID          int64  `db:"usex_city_id" json:"usexCityId"`
}

type UsersUsersExperiencesSkill struct {
	UeskUsexID int32 `db:"uesk_usex_id" json:"ueskUsexId"`
	UeskUskiID int32 `db:"uesk_uski_id" json:"ueskUskiId"`
}

type UsersUsersLicense struct {
	UsliID           int32          `db:"usli_id" json:"usliId"`
	UsliLicenseCode  string `db:"usli_license_code" json:"usliLicenseCode"`
	UsliModifiedDate time.Time   `db:"usli_modified_date" json:"usliModifiedDate"`
	UsliStatus       string `db:"usli_status" json:"usliStatus"`
	UsliEntityID     int32          `db:"usli_entity_id" json:"usliEntityId"`
}

type UsersUsersMedium struct {
	UsmeID           int32          `db:"usme_id" json:"usmeId"`
	UsmeEntityID     int32          `db:"usme_entity_id" json:"usmeEntityId"`
	UsmeFileLink     string `db:"usme_file_link" json:"usmeFileLink"`
	UsmeFilename     string `db:"usme_filename" json:"usmeFilename"`
	UsmeFilesize     int64  `db:"usme_filesize" json:"usmeFilesize"`
	UsmeFiletype     string `db:"usme_filetype" json:"usmeFiletype"`
	UsmeNote         string `db:"usme_note" json:"usmeNote"`
	UsmeModifiedDate time.Time   `db:"usme_modified_date" json:"usmeModifiedDate"`
}

type UsersUsersPhone struct {
	UspoEntityID     int32          `db:"uspo_entity_id" json:"uspoEntityId"`
	UspoNumber       string         `db:"uspo_number" json:"uspoNumber"`
	UspoModifiedDate time.Time   `db:"uspo_modified_date" json:"uspoModifiedDate"`
	UspoPontyCode    string `db:"uspo_ponty_code" json:"uspoPontyCode"`
}

type UsersUsersRole struct {
	UsroEntityID     int32        `db:"usro_entity_id" json:"usroEntityId"`
	UsroRoleID       int32        `db:"usro_role_id" json:"usroRoleId"`
	UsroModifiedDate time.Time `db:"usro_modified_date" json:"usroModifiedDate"`
}

type UsersUsersSkill struct {
	UskiID           int32          `db:"uski_id" json:"uskiId"`
	UskiEntityID     int32          `db:"uski_entity_id" json:"uskiEntityId"`
	UskiModifiedDate time.Time   `db:"uski_modified_date" json:"uskiModifiedDate"`
	UskiSktyName     string `db:"uski_skty_name" json:"uskiSktyName"`
}
