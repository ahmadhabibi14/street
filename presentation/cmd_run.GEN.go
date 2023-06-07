package presentation

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"

	"street/domain"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


func cmdRun(b *domain.Domain, action string, payload []byte) {
	switch action {

	case domain.AdminUserListAction:
		in := domain.AdminUserListIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.AdminUserList(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestDebugAction:
		in := domain.GuestDebugIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestDebug(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestExternalAuthAction:
		in := domain.GuestExternalAuthIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestExternalAuth(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestForgotPasswordAction:
		in := domain.GuestForgotPasswordIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestForgotPassword(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestLoginAction:
		in := domain.GuestLoginIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestLogin(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestOauthCallbackAction:
		in := domain.GuestOauthCallbackIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestOauthCallback(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestRegisterAction:
		in := domain.GuestRegisterIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestRegister(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestResendVerificationEmailAction:
		in := domain.GuestResendVerificationEmailIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestResendVerificationEmail(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestResetPasswordAction:
		in := domain.GuestResetPasswordIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestResetPassword(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestVerifyEmailAction:
		in := domain.GuestVerifyEmailIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestVerifyEmail(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.UserChangePasswordAction:
		in := domain.UserChangePasswordIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.UserChangePassword(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.UserDeactivateAction:
		in := domain.UserDeactivateIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.UserDeactivate(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.UserLogoutAction:
		in := domain.UserLogoutIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.UserLogout(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.UserProfileAction:
		in := domain.UserProfileIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.UserProfile(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.UserUpdateProfileAction:
		in := domain.UserUpdateProfileIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.UserUpdateProfile(&in)
		fmt.Println(X.ToJsonPretty(out))

	}
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
