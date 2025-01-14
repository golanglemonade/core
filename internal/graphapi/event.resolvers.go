package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/generated/privacy"
)

// CreateEvent is the resolver for the createEvent field
func (r *mutationResolver) CreateEvent(ctx context.Context, input generated.CreateEventInput) (*EventCreatePayload, error) {
	t, err := withTransactionalMutation(ctx).Event.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			validationError := err.(*generated.ValidationError)

			r.logger.Debugw("validation error", "field", validationError.Name, "error", validationError.Error())

			return nil, validationError
		}

		if generated.IsConstraintError(err) {
			constraintError := err.(*generated.ConstraintError)

			r.logger.Debugw("constraint error", "error", constraintError.Error())

			return nil, constraintError
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionCreate, "event")
		}

		r.logger.Errorw("failed to create event", "error", err)

		return nil, err
	}

	return &EventCreatePayload{Event: t}, nil
}

// CreateBulkEvent is the resolver for the createBulkEvent field.
func (r *mutationResolver) CreateBulkEvent(ctx context.Context, input []*generated.CreateEventInput) (*EventBulkCreatePayload, error) {
	return r.bulkCreateEvent(ctx, input)
}

// CreateBulkCSVEvent is the resolver for the createBulkCSVEvent field.
func (r *mutationResolver) CreateBulkCSVEvent(ctx context.Context, input graphql.Upload) (*EventBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateEventInput](input)
	if err != nil {
		r.logger.Errorw("failed to unmarshal bulk data", "error", err)

		return nil, err
	}

	return r.bulkCreateEvent(ctx, data)
}

// UpdateEvent is the resolver for the updateEvent field
func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, input generated.UpdateEventInput) (*EventUpdatePayload, error) {
	event, err := withTransactionalMutation(ctx).Event.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get event on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "event")
		}

		r.logger.Errorw("failed to get event", "error", err)
		return nil, ErrInternalServerError
	}

	event, err = event.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update event", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "event")
		}

		r.logger.Errorw("failed to update event", "error", err)
		return nil, ErrInternalServerError
	}

	return &EventUpdatePayload{Event: event}, nil
}

// DeleteEvent is the resolver for the deleteEvent field
func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (*EventDeletePayload, error) {
	if err := withTransactionalMutation(ctx).Event.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "event")
		}

		r.logger.Errorw("failed to delete event", "error", err)
		return nil, err
	}

	if err := generated.EventEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &EventDeletePayload{DeletedID: id}, nil
}

// Event is the resolver for the event field
func (r *queryResolver) Event(ctx context.Context, id string) (*generated.Event, error) {
	event, err := withTransactionalMutation(ctx).Event.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get event", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "event")
		}

		return nil, ErrInternalServerError
	}

	return event, nil
}
