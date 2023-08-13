package presentation

import (
	"os"

	"street/domain"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


func cmdRun(b *domain.Domain, action string, payload []byte) {
	switch action {

	case domain.AdminDashboardAction:
		in := domain.AdminDashboardIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.AdminDashboard(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.AdminFilesAction:
		in := domain.AdminFilesIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.AdminFiles(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.AdminPropHistoriesAction:
		in := domain.AdminPropHistoriesIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.AdminPropHistories(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.AdminPropertiesAction:
		in := domain.AdminPropertiesIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.AdminProperties(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.AdminUsersAction:
		in := domain.AdminUsersIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.AdminUsers(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestDebugAction:
		in := domain.GuestDebugIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestDebug(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestExternalAuthAction:
		in := domain.GuestExternalAuthIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestExternalAuth(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestFilesAction:
		in := domain.GuestFilesIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestFiles(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestForgotPasswordAction:
		in := domain.GuestForgotPasswordIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestForgotPassword(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestLoginAction:
		in := domain.GuestLoginIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestLogin(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestOauthCallbackAction:
		in := domain.GuestOauthCallbackIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestOauthCallback(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestRegisterAction:
		in := domain.GuestRegisterIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestRegister(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestResendVerificationEmailAction:
		in := domain.GuestResendVerificationEmailIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestResendVerificationEmail(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestResetPasswordAction:
		in := domain.GuestResetPasswordIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestResetPassword(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.GuestVerifyEmailAction:
		in := domain.GuestVerifyEmailIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestVerifyEmail(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.RealtorOwnedPropertiesAction:
		in := domain.RealtorOwnedPropertiesIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.RealtorOwnedProperties(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.RealtorUpsertPropertyAction:
		in := domain.RealtorUpsertPropertyIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.RealtorUpsertProperty(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserChangePasswordAction:
		in := domain.UserChangePasswordIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserChangePassword(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserDeactivateAction:
		in := domain.UserDeactivateIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserDeactivate(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserLogoutAction:
		in := domain.UserLogoutIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserLogout(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserProfileAction:
		in := domain.UserProfileIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserProfile(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserPropHistoryAction:
		in := domain.UserPropHistoryIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserPropHistory(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserSearchPropAction:
		in := domain.UserSearchPropIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserSearchProp(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserUpdateProfileAction:
		in := domain.UserUpdateProfileIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserUpdateProfile(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)


	case domain.UserUploadFileAction:
		in := domain.UserUploadFileIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserUploadFile(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	}
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
