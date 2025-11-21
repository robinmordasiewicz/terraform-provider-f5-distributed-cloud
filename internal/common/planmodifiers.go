// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// UseStateForUnknownString returns a plan modifier that copies the prior state
// value to the planned value when the planned value is unknown.
func UseStateForUnknownString() planmodifier.String {
	return useStateForUnknownString{}
}

type useStateForUnknownString struct{}

func (m useStateForUnknownString) Description(ctx context.Context) string {
	return "Use the prior state value if the planned value is unknown."
}

func (m useStateForUnknownString) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateForUnknownString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do not modify if there's no state value.
	if req.StateValue.IsNull() {
		return
	}

	// Do not modify if the value is known.
	if !req.PlanValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}

// UseStateForUnknownBool returns a plan modifier that copies the prior state
// value to the planned value when the planned value is unknown.
func UseStateForUnknownBool() planmodifier.Bool {
	return useStateForUnknownBool{}
}

type useStateForUnknownBool struct{}

func (m useStateForUnknownBool) Description(ctx context.Context) string {
	return "Use the prior state value if the planned value is unknown."
}

func (m useStateForUnknownBool) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateForUnknownBool) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.StateValue.IsNull() {
		return
	}

	if !req.PlanValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}

// UseStateForUnknownInt64 returns a plan modifier that copies the prior state
// value to the planned value when the planned value is unknown.
func UseStateForUnknownInt64() planmodifier.Int64 {
	return useStateForUnknownInt64{}
}

type useStateForUnknownInt64 struct{}

func (m useStateForUnknownInt64) Description(ctx context.Context) string {
	return "Use the prior state value if the planned value is unknown."
}

func (m useStateForUnknownInt64) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateForUnknownInt64) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.StateValue.IsNull() {
		return
	}

	if !req.PlanValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}

// UseStateForUnknownObject returns a plan modifier that copies the prior state
// value to the planned value when the planned value is unknown.
func UseStateForUnknownObject() planmodifier.Object {
	return useStateForUnknownObject{}
}

type useStateForUnknownObject struct{}

func (m useStateForUnknownObject) Description(ctx context.Context) string {
	return "Use the prior state value if the planned value is unknown."
}

func (m useStateForUnknownObject) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateForUnknownObject) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.StateValue.IsNull() {
		return
	}

	if !req.PlanValue.IsUnknown() {
		return
	}

	resp.PlanValue = req.StateValue
}

// DefaultValueString returns a plan modifier that sets a default value when
// the configuration value is null.
func DefaultValueString(defaultValue string) planmodifier.String {
	return defaultValueString{defaultValue: types.StringValue(defaultValue)}
}

type defaultValueString struct {
	defaultValue types.String
}

func (m defaultValueString) Description(ctx context.Context) string {
	return "Set a default value if the configuration value is null."
}

func (m defaultValueString) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m defaultValueString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}

	resp.PlanValue = m.defaultValue
}

// DefaultValueBool returns a plan modifier that sets a default value when
// the configuration value is null.
func DefaultValueBool(defaultValue bool) planmodifier.Bool {
	return defaultValueBool{defaultValue: types.BoolValue(defaultValue)}
}

type defaultValueBool struct {
	defaultValue types.Bool
}

func (m defaultValueBool) Description(ctx context.Context) string {
	return "Set a default value if the configuration value is null."
}

func (m defaultValueBool) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m defaultValueBool) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}

	resp.PlanValue = m.defaultValue
}

// DefaultValueInt64 returns a plan modifier that sets a default value when
// the configuration value is null.
func DefaultValueInt64(defaultValue int64) planmodifier.Int64 {
	return defaultValueInt64{defaultValue: types.Int64Value(defaultValue)}
}

type defaultValueInt64 struct {
	defaultValue types.Int64
}

func (m defaultValueInt64) Description(ctx context.Context) string {
	return "Set a default value if the configuration value is null."
}

func (m defaultValueInt64) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m defaultValueInt64) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if !req.ConfigValue.IsNull() {
		return
	}

	resp.PlanValue = m.defaultValue
}
