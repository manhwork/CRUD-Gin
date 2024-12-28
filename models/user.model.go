package models

// User là mô hình đại diện cho thông tin của người đi khám bệnh
type User struct {
	ID            uint   `json:"id"`
	CCCD          string `json:"cccd"`
	FullName      string `json:"full_name"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"` // Có thể là "Nam", "Nữ", "Khác"
	Nationality   string `json:"nationality"`
	PlaceOfBirth  string `json:"place_of_birth"` // Quê quán
	IssueDate     string `json:"issue_date"`     // Ngày cấp CCCD
	ExpiryDate    string `json:"expiry_date"`    // Ngày hết hạn CCCD
	OldCCCD       string `json:"old_cccd"`       // Số CCCD cũ
	Ethnicity     string `json:"ethnicity"`      // Dân tộc
	Religion      string `json:"religion"`       // Tôn giáo
	PersonalID    string `json:"personal_id"`    // Nhận dạng cá nhân
	PermanentAddr string `json:"permanent_addr"` // Địa chỉ thường trú
	FatherName    string `json:"father_name"`    // Tên cha
	MotherName    string `json:"mother_name"`    // Tên mẹ
	SpouseName    string `json:"spouse_name"`    // Tên vợ/chồng
}
