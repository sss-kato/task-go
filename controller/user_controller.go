package controller

import (
	"html"
	"net/http"
	"task-go/domain"
	"task-go/service"
	"task-go/util"

	"github.com/labstack/echo/v4"
)

type UserControllerIF interface {
	Signup(e echo.Context) error
}

type userController struct {
	us service.UserServiceIF
}

func NewUserController(us service.UserServiceIF) UserControllerIF {
	return &userController{us}
}

type UserRequest struct {
	Name       string
	Password   string
	MailAdress string
}

func (uc *userController) Signup(c echo.Context) error {

	userReq := &UserRequest{}

	if bindErr := c.Bind(userReq); bindErr != nil {
		// return err
		util.WriteErrLog(bindErr)
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	user := domain.NewUser(html.EscapeString(userReq.Name), userReq.Password, html.EscapeString(userReq.MailAdress))

	nameErr := user.ValidateName()
	if nameErr != nil {
		errMsg := &domain.Message{Message: nameErr.Error()}

		// errorLog := util.SetErrorInfo(nameErr)
		// util.WriteErrLog(errorLog)
		util.WriteErrLog(nameErr)

		return c.JSON(http.StatusBadRequest, errMsg)
	}

	pwErr := user.ValidatePassword()
	if pwErr != nil {
		errMsg := &domain.Message{Message: pwErr.Error()}

		// errorLog := util.SetErrorInfo(pwErr)
		util.WriteErrLog(pwErr)

		return c.JSON(http.StatusBadRequest, errMsg)
	}

	mailErr := user.ValidateMailAdress()
	if mailErr != nil {
		errMsg := &domain.Message{Message: mailErr.Error()}

		// errorLog := util.SetErrorInfo(mailErr)
		util.WriteErrLog(mailErr)

		return c.JSON(http.StatusBadRequest, errMsg)
	}

	userRes, signupErr := uc.us.Signup(user)
	if signupErr != nil {
		errMsg := &domain.Message{Message: "signup failed."}

		// errorLog := util.SetErrorInfo(signupErr)
		// util.WriteErrLog(errorLog)
		util.WriteErrLog(signupErr)

		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusCreated, userRes)

}

// func test(reqStruct interface{}) interface{} {
// 	reqStructType := reflect.TypeOf(reqStruct)
// 	// reqStructValue := reflect.ValueOf(&reqStruct).Elem()
// 	reqStructValue := reflect.ValueOf(reqStruct)
// 	fmt.Print(reqStructValue.Kind())

// 	for i := 0; i < reqStructType.NumField(); i++ {
// 		field := reqStructValue.Field(i)
// 		// サニタイズが必要かチェック
// 		if reqStructType.Field(i).Tag.Get("sanitized") == "true" {
// 			sanitizedValue := sanitizeField(field)
// 			test := "test"
// 			ptest := &test

// 			fmt.Print(sanitizedValue)
// 			// サニタイズされた値をセット
// 			// field.Set(reflect.ValueOf(sanitizedValue))
// 			a := reflect.ValueOf(ptest)
// 			pa := &a
// 			// field.Set(reflect.ValueOf(ptest))
// 			field.Set(a)

// 		}
// 	}
// 	fmt.Print(reqStruct)

// 	return nil
// }

// func sanitizeField(field reflect.Value) interface{} {
// 	switch field.Kind() {
// 	case reflect.String:
// 		// string型の場合はHTMLエスケープを実行
// 		return html.EscapeString(field.String())
// 	default:
// 		// その他の型の場合はそのまま返す
// 		return field.Interface()
// 	}
// }
