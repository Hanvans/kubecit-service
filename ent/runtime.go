// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kubecit-service/ent/category"
	"kubecit-service/ent/schema"
	"kubecit-service/ent/slider"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.DefaultName holds the default value on creation for the name field.
	category.DefaultName = categoryDescName.Default.(string)
	sliderFields := schema.Slider{}.Fields()
	_ = sliderFields
	// sliderDescTitle is the schema descriptor for title field.
	sliderDescTitle := sliderFields[0].Descriptor()
	// slider.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	slider.TitleValidator = sliderDescTitle.Validators[0].(func(string) error)
	// sliderDescContent is the schema descriptor for content field.
	sliderDescContent := sliderFields[1].Descriptor()
	// slider.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	slider.ContentValidator = sliderDescContent.Validators[0].(func(string) error)
	// sliderDescImageLink is the schema descriptor for image_link field.
	sliderDescImageLink := sliderFields[2].Descriptor()
	// slider.ImageLinkValidator is a validator for the "image_link" field. It is called by the builders before save.
	slider.ImageLinkValidator = sliderDescImageLink.Validators[0].(func(string) error)
	// sliderDescCreateAt is the schema descriptor for create_at field.
	sliderDescCreateAt := sliderFields[3].Descriptor()
	// slider.DefaultCreateAt holds the default value on creation for the create_at field.
	slider.DefaultCreateAt = sliderDescCreateAt.Default.(time.Time)
	// sliderDescUpdateAt is the schema descriptor for update_at field.
	sliderDescUpdateAt := sliderFields[4].Descriptor()
	// slider.DefaultUpdateAt holds the default value on creation for the update_at field.
	slider.DefaultUpdateAt = sliderDescUpdateAt.Default.(time.Time)
	// slider.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	slider.UpdateDefaultUpdateAt = sliderDescUpdateAt.UpdateDefault.(func() time.Time)
	// sliderDescIsValid is the schema descriptor for is_valid field.
	sliderDescIsValid := sliderFields[5].Descriptor()
	// slider.DefaultIsValid holds the default value on creation for the is_valid field.
	slider.DefaultIsValid = sliderDescIsValid.Default.(bool)
}
