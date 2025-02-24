package resolve

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubscriptionFilter(t *testing.T) {
	t.Run("in: predicate is true (boolean)", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"true"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":true}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("in: predicate is true (float)", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"1.13"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":1.13}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("in: predicate is true (int)", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"49"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":49}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("in: predicate is false (boolean)", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"false"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":true}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("in: predicate is false (float)", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"9.77"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":8.01}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("in: predicate is false (int)", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":123}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":321}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("in: predicate is false (string)", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"b"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("in: array predicate is false", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":["a","b"]}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("in: array predicate is true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			In: &SubscriptionFieldFilter{
				FieldPath: []string{"event"},
				Values: []InputTemplate{
					{
						Segments: []TemplateSegment{
							{
								SegmentType:        VariableSegmentType,
								VariableKind:       ContextVariableKind,
								VariableSourcePath: []string{"var"},
								Renderer:           NewPlainVariableRenderer(),
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":["a","b","c"]}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("not in: predicate is true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			Not: &SubscriptionFilter{
				In: &SubscriptionFieldFilter{
					FieldPath: []string{"event"},
					Values: []InputTemplate{
						{
							Segments: []TemplateSegment{
								{
									SegmentType:        VariableSegmentType,
									VariableKind:       ContextVariableKind,
									VariableSourcePath: []string{"var"},
									Renderer:           NewPlainVariableRenderer(),
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"b"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":"b"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("not in: predicate is false", func(t *testing.T) {
		filter := &SubscriptionFilter{
			Not: &SubscriptionFilter{
				In: &SubscriptionFieldFilter{
					FieldPath: []string{"event"},
					Values: []InputTemplate{
						{
							Segments: []TemplateSegment{
								{
									SegmentType:        VariableSegmentType,
									VariableKind:       ContextVariableKind,
									VariableSourcePath: []string{"var"},
									Renderer:           NewPlainVariableRenderer(),
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"var":"b"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"event":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("and: both in predicates are true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			And: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"first"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"second"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"first":"b","second":"c"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("and: static predicates are true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			And: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType: StaticSegmentType,
										Data:        []byte("b"),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType: StaticSegmentType,
										Data:        []byte("c"),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("and: first in predicate is false", func(t *testing.T) {
		filter := &SubscriptionFilter{
			And: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"first"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"second"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"first":"d","second":"c"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("and: second in predicate is false", func(t *testing.T) {
		filter := &SubscriptionFilter{
			And: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"first"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"second"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"first":"b","unused":"c"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
	t.Run("or: both in predicates are true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			Or: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"first"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"second"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"first":"b","second":"c"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("or: first in predicate is true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			Or: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"first"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"second"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"first":"b","unused":"c"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("or: second in predicate is true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			Or: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"first"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"second"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"third":"b","second":"c"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, false, skip)
	})
	t.Run("or: neither in predicate is true", func(t *testing.T) {
		filter := &SubscriptionFilter{
			Or: []SubscriptionFilter{
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventX"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"first"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
				{
					In: &SubscriptionFieldFilter{
						FieldPath: []string{"eventY"},
						Values: []InputTemplate{
							{
								Segments: []TemplateSegment{
									{
										SegmentType:        VariableSegmentType,
										VariableKind:       ContextVariableKind,
										VariableSourcePath: []string{"second"},
										Renderer:           NewPlainVariableRenderer(),
									},
								},
							},
						},
					},
				},
			},
		}
		c := &Context{
			Variables: []byte(`{"fourth":"b","third":"c"}`),
		}
		buf := &bytes.Buffer{}
		data := []byte(`{"eventX":"b","eventY":"c"}`)
		skip, err := filter.SkipEvent(c, data, buf)
		assert.NoError(t, err)
		assert.Equal(t, true, skip)
	})
}
