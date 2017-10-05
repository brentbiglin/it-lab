package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Device struct {
	item.Item

	Model             string `json:"model"`
	Descriptor        string `json:"descriptor"`
	Photo             string `json:"photo"`
	IncludedMaterials string `json:"included_materials"`
	Availability      string `json:"availability"`
	BorrowerName      string `json:"borrower_name"`
	BorrowerEmail     string `json:"borrower_email"`
	Location          string `json:"location"`
	DeviceCondition   string `json:"device_condition"`
	UTID              string `json:"u_t_i_d"`
	SerialNumber      string `json:"serial_number"`
	SoftwareVersion   string `json:"software_version"`
	Other             string `json:"other"`
}

// MarshalEditor writes a buffer of html to edit a Device within the CMS
// and implements editor.Editable
func (d *Device) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(d,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Device field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Model", d, map[string]string{
				"label":       "Model",
				"type":        "text",
				"placeholder": "Enter the model here",
			}),
		},
		editor.Field{
			View: editor.Input("Descriptor", d, map[string]string{
				"label":       "Descriptor",
				"type":        "text",
				"placeholder": "Enter the descriptor here",
			}),
		},
		editor.Field{
			View: editor.File("Photo", d, map[string]string{
				"label":       "Photo",
				"placeholder": "Upload a photo here",
			}),
		},
		editor.Field{
			View: editor.Textarea("IncludedMaterials", d, map[string]string{
				"label":       "Included Materials",
				"placeholder": "Enter any included materials here",
			}),
		},
		editor.Field{
			View: editor.Select("Availability", d, map[string]string{
				"label": "Availability",
			}, map[string]string{
			// "value": "Display Name",
				"Available": "Available",
				"Reserved":  "Reserved",
				"In Use":    "In Use",
				"Due":	     "Due",
			}),
		},
		editor.Field{
			View: editor.Input("BorrowerName", d, map[string]string{
				"label":       "Borrower Name",
				"type":        "text",
				"placeholder": "Enter the borrower name here",
			}),
		},
		editor.Field{
			View: editor.Input("BorrowerEmail", d, map[string]string{
				"label":       "Borrower Email",
				"type":        "text",
				"placeholder": "Enter the borrower email here",
			}),
		},
		editor.Field{
			View: editor.Select("Location", d, map[string]string{
				"label": "Location",
			}, map[string]string{
			// "value": "Display Name",
				"Laptop Checkout Drawer": "Laptop Checkout Drawer",
				"Digital Camera Shelf":   "Digital Camera Shelf",
			}),
		},
		editor.Field{
			View: editor.Select("DeviceCondition", d, map[string]string{
				"label": "Device Condition",
			}, map[string]string{
			// "value": "Display Name",
				"Excellent": "Excellent",
				"Good":      "Good",
				"Fair":      "Fair",
				"Poor":      "Poor",
			}),
		},
		editor.Field{
			View: editor.Input("UTID", d, map[string]string{
				"label":       "UTID",
				"type":        "text",
				"placeholder": "Enter the UTID here",
			}),
		},
		editor.Field{
			View: editor.Input("SerialNumber", d, map[string]string{
				"label":       "Serial Number",
				"type":        "text",
				"placeholder": "Enter the serial number here",
			}),
		},
		editor.Field{
			View: editor.Input("SoftwareVersion", d, map[string]string{
				"label":       "Software Version",
				"type":        "text",
				"placeholder": "Enter the software version here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Other", d, map[string]string{
				"label":       "Other",
				"placeholder": "Enter any other relevant information here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Device editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Device"] = func() interface{} { return new(Device) }
}

// String defines how a Device is printed. Update it using more descriptive
// fields from the Device struct type
func (d *Device) String() string {
	return d.Model
}
