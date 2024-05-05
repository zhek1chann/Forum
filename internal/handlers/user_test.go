package handlers

import (
	mock "forum/internal/repo/mocks"
	"net/http"
	"net/url"
	"testing"
)

func TestSignUp(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Close()

	const (
		validUsername = "max"
		validEmail    = "max@gmail.com"
		validPassword = "max@gmail.com"
	)

	tests := []struct {
		name          string
		username      string
		email         string
		password      string
		passwordAgain string
		wantCode      int
		wantBody      string
	}{
		{
			name:          "Valid signup",
			username:      validUsername,
			email:         validEmail,
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Blank username",
			username:      "",
			email:         validEmail,
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Too long username",
			username:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.",
			email:         validEmail,
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Invalid username (non-ascii)",
			username:      "нееееееееееее",
			email:         validEmail,
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Duplicate username",
			username:      "zubenko",
			email:         validEmail,
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Blank email",
			username:      validUsername,
			email:         "",
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Too long email",
			username:      validUsername,
			email:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.",
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Invalid email",
			username:      validUsername,
			email:         "some@invalid.email.shoud.be.here@.",
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Invalid email (non-ascii)",
			username:      validUsername,
			email:         "недействительный@мейл.ру",
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Duplicate email",
			username:      validUsername,
			email:         "zubenko@gmail.com",
			password:      validPassword,
			passwordAgain: validPassword,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Blank password",
			username:      validUsername,
			email:         validEmail,
			password:      "",
			passwordAgain: "",
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Short password",
			username:      validUsername,
			email:         validEmail,
			password:      "a",
			passwordAgain: "a",
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Too long password",
			username:      validUsername,
			email:         validEmail,
			password:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.",
			passwordAgain: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam porttitor laoreet nisi eget molestie. Morbi vestibulum enim nec pharetra mattis. Etiam iaculis consequat risus, et facilisis elit venenatis ac. Suspendisse at consectetur nibh, quis interdum leo. Ut convallis eget justo vitae condimentum. Vivamus justo mauris, iaculis vitae ex nec, vehicula blandit est. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Praesent aliquet fermentum turpis nec rutrum.",
			wantCode:      http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("username-signup", tt.username)
			form.Add("email-signup", tt.email)
			form.Add("password-signup", tt.password)
			form.Add("password-again", tt.passwordAgain)

			code, _, _ := ts.postForm(t, "/signup", form)

			mock.Equal(t, code, tt.wantCode)
		})
	}
}

func TestUserLoginPost(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Close()

	const (
		validEmail    = "max@gmail.com"
		validPassword = "max@gmail.com"
	)

	tests := []struct {
		name     string
		email    string
		password string
		wantCode int
		wantBody string
	}{

		{
			name:     "Valid login",
			email:    validEmail,
			password: validPassword,
			wantCode: http.StatusUnprocessableEntity,
		},

		{
			name:     "Incorrect email",
			email:    "naaaaaaaah@gmail.com@gmail.com",
			password: validPassword,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name:     "Invalid identifier (non-existent email)",
			email:    "dontexist@gmail.com",
			password: validPassword,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name:     "Blank password",
			email:    validEmail,
			password: "",
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name:     "Blank email",
			email:    "",
			password: validPassword,
			wantCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("email-login", tt.email)
			form.Add("password-login", tt.password)

			code, _, _ := ts.postForm(t, "/login", form)

			mock.Equal(t, code, tt.wantCode)
		})
	}
}
