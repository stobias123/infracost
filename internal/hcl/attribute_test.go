package hcl

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestAttributeValueWithIncompleteContextAndConditionalShouldNotPanic(t *testing.T) {
	p := hclparse.NewParser()
	f, diags := p.ParseHCL([]byte(`
locals {
  original_tags    = "test"
  transformed_tags = local.original_tags
  id = var.enabled ? local.transformed_tags : "test3"
}
`), "test")

	require.False(t, diags.HasErrors(), fmt.Sprintf("diags has unexpected error %s from parsing input string", diags.Error()))

	c, _, diags := f.Body.PartialContent(terraformSchemaV012)
	require.False(t, diags.HasErrors(), "diags has unexpected error %s from parsing body content", diags.Error())

	var block *hcl.Block
	for _, b := range c.Blocks {
		if b.Type == "locals" {
			block = b
		}
	}

	require.NotNil(t, block, "could not find required test block")

	attrs, diags := block.Body.JustAttributes()
	require.False(t, diags.HasErrors(), "diags has unexpected error %s fetching attributes", diags.Error())

	buf := bytes.NewBuffer([]byte{})
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	l.SetOutput(buf)
	l.SetLevel(logrus.DebugLevel)
	logger := logrus.NewEntry(l)

	l2 := logrus.New()
	l2.SetOutput(io.Discard)
	discard := logrus.NewEntry(l2)

	tag := Attribute{
		HCLAttr: attrs["transformed_tags"],
		Ctx: &Context{ctx: &hcl.EvalContext{
			Variables: map[string]cty.Value{},
		}},
		Logger: discard,
	}

	attr := Attribute{
		HCLAttr: attrs["id"],
		Ctx: &Context{
			ctx: &hcl.EvalContext{
				Variables: map[string]cty.Value{
					"local": cty.ObjectVal(map[string]cty.Value{
						"original_tags":    cty.StringVal("test"),
						"transformed_tags": tag.Value(),
					}),
					"var": cty.ObjectVal(map[string]cty.Value{
						"enabled": cty.BoolVal(true),
					}),
				},
			},
			logger: logger,
		},
		Verbose: false,
		Logger:  logger,
	}

	v := attr.Value()
	assert.Equal(t, cty.DynamicVal, v)

	b, err := io.ReadAll(buf)
	require.NoError(t, err)

	assert.NotContains(t, string(b), "invalid memory address")
}
