package presentation

// Code generated by 1_codegen_test.go DO NOT EDIT.

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"street/domain"
)

func ApiRoutes(fw *fiber.App, d *domain.Domain) {

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
		return in.ToFiberCtx(c, out)
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
		return in.ToFiberCtx(c, out)
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
		return in.ToFiberCtx(c, out)
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
		return in.ToFiberCtx(c, out)
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
		return in.ToFiberCtx(c, out)
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
		return in.ToFiberCtx(c, out)
	})

}

// Code generated by 1_codegen_test.go DO NOT EDIT.
