// Author: yangzq80@gmail.com
// Date: 2020-08-26
//
package c4model

import (
	"fmt"
	. "goa.design/model/dsl"
	"goa.design/model/eval"
	"goa.design/model/service"
	"os"
	"testing"
)

// DSL that describes software architecture model.
var _ = Workspace("Getting Started", "This is a model of my software system.", func() {
	var System = SoftwareSystem("Software System", "My software system.", func() {
		Tag("system")
	})

	Person("User", "A user of my software system.", func() {
		Uses(System, "Uses")
		Tag("person")
	})

	Views(func() {
		SystemContextView(System, "SystemContext", "An example of a System Context diagram.", func() {
			AddAll()
			AutoLayout(RankLeftRight)
		})
		Styles(func() {
			ElementStyle("system", func() {
				Background("#1168bd")
				Color("#ffffff")
			})
			ElementStyle("person", func() {
				Shape(ShapePerson)
				Background("#08427b")
				Color("#ffffff")
			})
		})
	})
})

// Executes the DSL and uploads the corresponding workspace to Structurizr.
func TestC4(t *testing.T) {
	// Run the model DSL
	w, err := eval.RunDSL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid model: %s", err.Error())
		os.Exit(1)
	}

	// Upload the model to the Structurizr service.
	// The API key and secret must be set in the STRUCTURIZR_KEY and
	// STRUCTURIZR_SECRET environment variables respectively. The
	// workspace ID must be set in STRUCTURIZR_WORKSPACE_ID.
	var (
		//key    = os.Getenv("STRUCTURIZR_KEY")
		//secret = os.Getenv("STRUCTURIZR_SECRET")
		//wid    = os.Getenv("STRUCTURIZR_WORKSPACE_ID")
		key    = "6889ceeb-3fcb-499f-b066-d8bf2fed2d9b"
		secret = "d06eb874-a590-4a20-9f2c-c231377df966"
		wid    = "57780"
	)

	c := service.NewClient(key, secret)
	if err := c.Put(wid, w); err != nil {
		fmt.Fprintf(os.Stderr, "failed to store workspace: %s", err.Error())
		os.Exit(1)
	}
}
