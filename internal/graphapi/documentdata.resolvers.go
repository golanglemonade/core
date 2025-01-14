package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/utils/rout"
)

// CreateDocumentData is the resolver for the createDocumentData field.
func (r *mutationResolver) CreateDocumentData(ctx context.Context, input generated.CreateDocumentDataInput) (*DocumentDataCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	res, err := withTransactionalMutation(ctx).DocumentData.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "documentdata"}, r.logger)
	}

	return &DocumentDataCreatePayload{
		DocumentData: res,
	}, nil
}

// CreateBulkDocumentData is the resolver for the createBulkDocumentData field.
func (r *mutationResolver) CreateBulkDocumentData(ctx context.Context, input []*generated.CreateDocumentDataInput) (*DocumentDataBulkCreatePayload, error) {
	return r.bulkCreateDocumentData(ctx, input)
}

// CreateBulkCSVDocumentData is the resolver for the createBulkCSVDocumentData field.
func (r *mutationResolver) CreateBulkCSVDocumentData(ctx context.Context, input graphql.Upload) (*DocumentDataBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateDocumentDataInput](input)
	if err != nil {
		r.logger.Errorw("failed to unmarshal bulk data", "error", err)

		return nil, err
	}

	return r.bulkCreateDocumentData(ctx, data)
}

// UpdateDocumentData is the resolver for the updateDocumentData field.
func (r *mutationResolver) UpdateDocumentData(ctx context.Context, id string, input generated.UpdateDocumentDataInput) (*DocumentDataUpdatePayload, error) {
	res, err := withTransactionalMutation(ctx).DocumentData.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "documentdata"}, r.logger)
	}
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &res.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, ErrPermissionDenied
	}

	// setup update request
	req := res.Update().SetInput(input).AppendTags(input.AppendTags)

	res, err = req.Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "documentdata"}, r.logger)
	}

	return &DocumentDataUpdatePayload{
		DocumentData: res,
	}, nil
}

// DeleteDocumentData is the resolver for the deleteDocumentData field.
func (r *mutationResolver) DeleteDocumentData(ctx context.Context, id string) (*DocumentDataDeletePayload, error) {
	if err := withTransactionalMutation(ctx).DocumentData.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "documentdata"}, r.logger)
	}

	if err := generated.DocumentDataEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &DocumentDataDeletePayload{
		DeletedID: id,
	}, nil
}

// DocumentData is the resolver for the documentData field.
func (r *queryResolver) DocumentData(ctx context.Context, id string) (*generated.DocumentData, error) {
	res, err := withTransactionalMutation(ctx).DocumentData.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "documentdata"}, r.logger)
	}

	return res, nil
}
