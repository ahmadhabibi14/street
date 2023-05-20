package presentation

// Code generated by 1_codegen_test.go DO NOT EDIT.

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"street/domain"
)

func ApiRoutes(fw *fiber.App, d *domain.Domain) {

	// GuestDebug
	fw.Post("/"+domain.GuestDebugAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.GuestDebugIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestDebugAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.GuestDebug(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// GuestForgotPassword
	fw.Post("/"+domain.GuestForgotPasswordAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.GuestForgotPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestForgotPasswordAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.GuestForgotPassword(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// GuestLogin
	fw.Post("/"+domain.GuestLoginAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.GuestLoginIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestLoginAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.GuestLogin(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// GuestRegister
	fw.Post("/"+domain.GuestRegisterAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.GuestRegisterIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestRegisterAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.GuestRegister(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// GuestResendVerificationEmail
	fw.Post("/"+domain.GuestResendVerificationEmailAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.GuestResendVerificationEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResendVerificationEmailAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.GuestResendVerificationEmail(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// GuestResetPassword
	fw.Post("/"+domain.GuestResetPasswordAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.GuestResetPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResetPasswordAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.GuestResetPassword(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// GuestVerifyEmail
	fw.Post("/"+domain.GuestVerifyEmailAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.GuestVerifyEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestVerifyEmailAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.GuestVerifyEmail(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// UserLogout
	fw.Post("/"+domain.UserLogoutAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.UserLogoutIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserLogoutAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.UserLogout(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

	// UserProfile
	fw.Post("/"+domain.UserProfileAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.UserProfileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := d.UserProfile(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon)
	})

}

// Code generated by 1_codegen_test.go DO NOT EDIT.
