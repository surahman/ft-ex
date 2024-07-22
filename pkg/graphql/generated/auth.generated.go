// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_generated

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/surahman/FTeX/pkg/models"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _JWTAuthResponse_token(ctx context.Context, field graphql.CollectedField, obj *models.JWTAuthResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_JWTAuthResponse_token(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Token, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_JWTAuthResponse_token(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "JWTAuthResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _JWTAuthResponse_expires(ctx context.Context, field graphql.CollectedField, obj *models.JWTAuthResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_JWTAuthResponse_expires(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Expires, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int64)
	fc.Result = res
	return ec.marshalNInt642int64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_JWTAuthResponse_expires(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "JWTAuthResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int64 does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _JWTAuthResponse_threshold(ctx context.Context, field graphql.CollectedField, obj *models.JWTAuthResponse) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_JWTAuthResponse_threshold(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Threshold, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int64)
	fc.Result = res
	return ec.marshalNInt642int64(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_JWTAuthResponse_threshold(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "JWTAuthResponse",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int64 does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var jWTAuthResponseImplementors = []string{"JWTAuthResponse"}

func (ec *executionContext) _JWTAuthResponse(ctx context.Context, sel ast.SelectionSet, obj *models.JWTAuthResponse) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, jWTAuthResponseImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("JWTAuthResponse")
		case "token":
			out.Values[i] = ec._JWTAuthResponse_token(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "expires":
			out.Values[i] = ec._JWTAuthResponse_expires(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "threshold":
			out.Values[i] = ec._JWTAuthResponse_threshold(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNJWTAuthResponse2githubᚗcomᚋsurahmanᚋFTeXᚋpkgᚋmodelsᚐJWTAuthResponse(ctx context.Context, sel ast.SelectionSet, v models.JWTAuthResponse) graphql.Marshaler {
	return ec._JWTAuthResponse(ctx, sel, &v)
}

func (ec *executionContext) marshalNJWTAuthResponse2ᚖgithubᚗcomᚋsurahmanᚋFTeXᚋpkgᚋmodelsᚐJWTAuthResponse(ctx context.Context, sel ast.SelectionSet, v *models.JWTAuthResponse) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._JWTAuthResponse(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
