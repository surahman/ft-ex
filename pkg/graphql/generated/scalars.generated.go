// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_generated

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNAny2interface(ctx context.Context, v interface{}) (any, error) {
	res, err := graphql.UnmarshalAny(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNAny2interface(ctx context.Context, sel ast.SelectionSet, v any) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	res := graphql.MarshalAny(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNAny2ᚕinterfaceᚄ(ctx context.Context, v interface{}) ([]any, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]any, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNAny2interface(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNAny2ᚕinterfaceᚄ(ctx context.Context, sel ast.SelectionSet, v []any) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNAny2interface(ctx, sel, v[i])
	}

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) unmarshalNInt642int64(ctx context.Context, v interface{}) (int64, error) {
	res, err := graphql.UnmarshalInt64(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNInt642int64(ctx context.Context, sel ast.SelectionSet, v int64) graphql.Marshaler {
	res := graphql.MarshalInt64(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNUUID2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNUUID2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalOInt322ᚖint32(ctx context.Context, v interface{}) (*int32, error) {
	if v == nil {
		return nil, nil
	}
	res, err := graphql.UnmarshalInt32(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOInt322ᚖint32(ctx context.Context, sel ast.SelectionSet, v *int32) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	res := graphql.MarshalInt32(*v)
	return res
}

// endregion ***************************** type.gotpl *****************************
