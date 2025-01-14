package request

// 비밀번호, 이메일, 생년월일, 성별 정도만 정보를 받는다.
type ReqSignup struct {
	Password string `json:"password" validate:"required,min=6" example:"6글자 이상"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required" example:"홍길동"`
	Sex      string `json:"sex" validate:"required,oneof=male female" example:"male / female"`
	Birth    string `json:"birth" validate:"required,datetime=2006-01-02" example:"1990-01-01"`
	AuthCode string `json:"authCode" validate:"required" example:"인증코드"`
}
