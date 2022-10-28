// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	token "orginone/app/user/cmd/api/internal/handler/token"
	user "orginone/app/user/cmd/api/internal/handler/user"
	"orginone/app/user/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/token", // TEST OK
				Handler: token.BladeauthtokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/token/user/v1/user/reReg", // TEST OK
				Handler: token.UseruserreRegHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/token/getSimplePwd", // TEST OK
				Handler: token.GetSimplePwdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/apptoken", // TEST OK
				Handler: token.BladeauthapptokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/app-info", // TEST OK
				Handler: token.BladeauthappinfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/asset-user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/switch-tenant", // TEST OK
				Handler: token.SwitchHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/retrieve/user",
				Handler: token.UserretrieveuserHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/update/pwd",
				Handler: token.UserupdatepwdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-user/token"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add/tenant", // TEST OK
				Handler: user.AddtenantHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add/user/v2",
				Handler: user.Adduserv2Handler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/authority/transfer",
				Handler: user.AuthoritytransferHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/check/token",
				Handler: user.ChecktokenHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/createPassword",
				Handler: user.CreatePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete/by/userIds",
				Handler: user.DeletebyuserIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete/person",
				Handler: user.DeletepersonHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete/userInfo", // TEST OK
				Handler: user.DeleteuserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: user.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/all/usable/user",
				Handler: user.GetallusableuserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/get/test/user",
				Handler: user.V1gettestuserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/all/user",
				Handler: user.GetalluserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/all/users",
				Handler: user.GetallusersHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/online/user/count",
				Handler: user.GetonlineusercountHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/person/tenantCode/v2",
				Handler: user.GetpersontenantCodev2Handler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/user/by/phone",
				Handler: user.GetuserbyphoneHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get/user/tenantCode",
				Handler: user.GetusertenantCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getManagerUserIdList",
				Handler: user.GetManagerUserIdListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: user.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listPage", // TEST OK
				Handler: user.ListPageHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/remove",
				Handler: user.RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/removeByIds", // TEST OK
				Handler: user.RemoveByIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/resetPassword", // TEST OK
				Handler: user.ResetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/submit", // TEST OK
				Handler: user.SubmitHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/tenant/user", // TEST OK
				Handler: user.TenantuserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/token/getTenantUserId",
				Handler: user.TokengetTenantUserIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update/login/tenant",
				Handler: user.UpdatelogintenantHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update/password", // TEST OK
				Handler: user.UpdatepasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/reset-password",
				Handler: user.ResetpasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update/user/info", // TEST OK
				Handler: user.UpdateuserinfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/update",
				Handler: user.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/quit/tenant",
				Handler: user.UserquittenantHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/v1/reView/user", // TEST OK
				Handler: user.ReViewuserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/withdraw/application", // TEST OK
				Handler: user.WithdrawapplicationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/asset-user/user"),
	)
}