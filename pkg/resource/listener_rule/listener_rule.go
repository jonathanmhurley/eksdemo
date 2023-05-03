package listener_rule

import (
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/resource"
)

func NewResource() *resource.Resource {
	res := &resource.Resource{
		Command: cmd.Command{
			Name:        "listener-rule",
			Description: "Load Balancer Listener Rule",
			Aliases:     []string{"listener-rules", "listen-rules", "listen-rule", "lr"},
			Args:        []string{"ID"},
		},

		Getter: &Getter{},
	}

	res.Options, res.GetFlags = newOptions()

	return res
}
