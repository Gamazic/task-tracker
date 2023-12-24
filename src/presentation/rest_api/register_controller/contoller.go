package register_controller

import (
	"errors"
	"io"
	"net/http"
	"tracker_backend/src/application"
	"tracker_backend/src/presentation/rest_api/microframework"
)

type RegisterController struct {
	RegisterFactory AbsRegisterFactory
	Logger          microframework.Logger
}

func (u RegisterController) Post(w http.ResponseWriter, r *http.Request) {
	var body RegisterRequestModel
	err := microframework.ReadValidate(r.Body, &body)
	if errors.Is(err, io.EOF) {
		microframework.SendValidationError(w, errors.New("empty body"))
		return
	}
	if errors.Is(err, microframework.ValidationErr) {
		microframework.SendValidationError(w, err)
		return
	}
	if err != nil {
		u.Logger.Errorf("register post parsing: %s", err)
		microframework.SendValidationError(w, errors.New("bad body"))
		return
	}
	ctx := r.Context()
	register, err := u.RegisterFactory.Build(RegisterDeps{
		Ctx:      ctx,
		Username: body.Username,
		Password: body.Password,
	})
	if err != nil {
		u.Logger.Errorf("register post building: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	err = register.Register()
	if errors.Is(err, application.ErrIdentityAlreadyExist) {
		u.Logger.LogIfErr(microframework.NewResponseBuilder(w).
			BuildStatus(http.StatusBadRequest).
			BuildBodyNestedMsg("user already exist").
			Send())
		return
	}
	if err != nil {
		u.Logger.Errorf("register post usecase call: %s", err)
		microframework.SendInternalServerError(w)
		return
	}
	u.Logger.LogIfErr(microframework.NewResponseBuilder(w).
		BuildStatus(http.StatusCreated).
		BuildBody(RegisterResponseModel{Username: body.Username}).
		Send())
}
