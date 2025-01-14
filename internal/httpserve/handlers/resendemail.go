package handlers

import (
	"errors"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	echo "github.com/theopenlane/echox"

	"github.com/theopenlane/utils/rout"

	"github.com/theopenlane/iam/auth"

	ent "github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/privacy/token"
	"github.com/theopenlane/core/pkg/enums"
	"github.com/theopenlane/core/pkg/models"
)

// ResendEmail will resend an email verification email if the provided email exists
func (h *Handler) ResendEmail(ctx echo.Context) error {
	var in models.ResendRequest
	if err := ctx.Bind(&in); err != nil {
		return h.InvalidInput(ctx, err)
	}

	if err := in.Validate(); err != nil {
		return h.InvalidInput(ctx, err)
	}

	// set viewer context
	ctxWithToken := token.NewContextWithSignUpToken(ctx.Request().Context(), in.Email)

	out := &models.ResendReply{
		Reply:   rout.Reply{Success: true},
		Message: "We've received your request to be resent an email to complete verification. Please check your email.",
	}

	// email verifications only come to users that were created with username/password logins
	entUser, err := h.getUserByEmail(ctxWithToken, in.Email, enums.AuthProviderCredentials)
	if err != nil {
		if ent.IsNotFound(err) {
			// return a 200 response even if user is not found to avoid
			// exposing confidential information
			return h.Success(ctx, out)
		}

		h.Logger.Errorf("error retrieving user email", "error", err)

		return h.InternalServerError(ctx, ErrProcessingRequest)
	}

	// check to see if user is already confirmed
	if entUser.Edges.Setting.EmailConfirmed {
		out.Message = "email is already confirmed"

		return h.Success(ctx, out)
	}

	// setup user context
	userCtx := auth.AddAuthenticatedUserContext(ctx, &auth.AuthenticatedUser{
		SubjectID: entUser.ID,
	})

	// create email verification token
	user := &User{
		FirstName: entUser.FirstName,
		LastName:  entUser.LastName,
		Email:     entUser.Email,
		ID:        entUser.ID,
	}

	if _, err = h.storeAndSendEmailVerificationToken(userCtx, user); err != nil {
		h.Logger.Errorw("error storing token", "error", err)

		if errors.Is(err, ErrMaxAttempts) {
			return h.TooManyRequests(ctx, err)
		}

		return h.InternalServerError(ctx, ErrProcessingRequest)
	}

	return h.Success(ctx, out)
}

// BindResendEmailHandler binds the resend email verification endpoint to the OpenAPI schema
func (h *Handler) BindResendEmailHandler() *openapi3.Operation {
	resendEmail := openapi3.NewOperation()
	resendEmail.Description = "ResendEmail accepts an email address via a POST request and always returns a 200 Status OK response, no matter the input or result of the processing. This is to ensure that no secure information is leaked from this unauthenticated endpoint. If the email address belongs to a user who has not been verified, another verification email is sent. If the post request contains an orgID and the user is invited to that organization but hasn't accepted the invite, then the invite is resent."
	resendEmail.OperationID = "ResendEmail"
	resendEmail.Security = &openapi3.SecurityRequirements{}

	h.AddRequestBody("ResendEmail", models.ExampleResendEmailSuccessRequest, resendEmail)
	h.AddResponse("ResendReply", "success", models.ExampleResendEmailSuccessResponse, resendEmail, http.StatusOK)
	resendEmail.AddResponse(http.StatusInternalServerError, internalServerError())
	resendEmail.AddResponse(http.StatusBadRequest, badRequest())

	return resendEmail
}
