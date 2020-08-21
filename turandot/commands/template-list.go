package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tliron/puccini/ard"
	puccinicommon "github.com/tliron/puccini/common"
	formatpkg "github.com/tliron/puccini/common/format"
	"github.com/tliron/puccini/common/terminal"
	urlpkg "github.com/tliron/puccini/url"
	clientpkg "github.com/tliron/turandot/client"
	"github.com/tliron/turandot/common"
)

func init() {
	templateCommand.AddCommand(templateListCommand)
}

var templateListCommand = &cobra.Command{
	Use:   "list",
	Short: "List service templates registered in the inventory",
	Run: func(cmd *cobra.Command, args []string) {
		ListServiceTemplates()
	},
}

func ListServiceTemplates() {
	images, err := NewClient().Spooler().List()
	puccinicommon.FailOnError(err)
	if len(images) == 0 {
		return
	}
	sort.Strings(images)

	switch format {
	case "":
		client := NewClient().Turandot()
		urlContext := urlpkg.NewContext()
		defer urlContext.Release()

		table := common.NewTable(maxWidth, "Name", "Services")
		for _, image := range images {
			if serviceTemplateName, ok := clientpkg.ServiceTemplateNameFromInventoryImageName(image); ok {
				services, err := client.ListServicesForImage(image, urlContext)
				puccinicommon.FailOnError(err)
				sort.Strings(services)
				table.Add(serviceTemplateName, strings.Join(services, "\n"))
			}
		}
		table.Print()

	case "bare":
		for _, image := range images {
			if serviceTemplateName, ok := clientpkg.ServiceTemplateNameFromInventoryImageName(image); ok {
				fmt.Fprintln(terminal.Stdout, serviceTemplateName)
			}
		}

	default:
		client := NewClient().Turandot()
		urlContext := urlpkg.NewContext()
		defer urlContext.Release()

		list := make(ard.List, 0, len(images))
		for _, image := range images {
			if serviceTemplateName, ok := clientpkg.ServiceTemplateNameFromInventoryImageName(image); ok {
				map_ := make(ard.StringMap)
				map_["Name"] = serviceTemplateName
				map_["Services"], err = client.ListServicesForImage(image, urlContext)
				puccinicommon.FailOnError(err)
				list = append(list, map_)
			}
		}
		formatpkg.Print(list, format, terminal.Stdout, strict, pretty)
	}
}
