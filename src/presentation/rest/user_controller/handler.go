package user_controller

import (
	"errors"
	"net/http"
	userUsecase "tracker_backend/src/application/user"
	userFactory "tracker_backend/src/factory/user"
	"tracker_backend/src/infrastructure"
	"tracker_backend/src/presentation/rest/microframework"
)

type UserHandler struct {
	CreateUserFactory userFactory.AbsCreateUserFactory
	Logger            infrastructure.Logger
}

func (u UserHandler) Post(w http.ResponseWriter, r *http.Request) {
	var body UserRequestModel
	err := microframework.ReadValidate(r.Body, &body)
	if errors.Is(err, microframework.ValidationErr) {
		microframework.SendValidationError(w, err)
		return
	}
	if err != nil {
		u.Logger.Errorf("user post parsing: %s", err)
		microframework.SendValidationError(w, errors.New("bad body"))
		return
	}

	ctx := r.Context()
	createUsecase, err := u.CreateUserFactory.Build(userFactory.CreateUserDeps{Ctx: ctx})
	if err != nil {
		u.Logger.Errorf("user post building: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	user := userUsecase.UserInCreate{
		Username: body.Username,
	}
	err = createUsecase.Execute(user)
	if errors.Is(err, userUsecase.ErrUserAlreadyExist) {
		u.Logger.LogIfErr(microframework.NewResponseBuilder(w).
			BuildStatus(http.StatusBadRequest).
			BuildBodyNestedMsg("user already exist").
			Send())
		return
	}
	if err != nil {
		u.Logger.Errorf("user post usecase call: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	u.Logger.LogIfErr(microframework.NewResponseBuilder(w).
		BuildStatus(http.StatusCreated).
		BuildBody(UserResponseModel{Username: user.Username}).
		Send())
}
