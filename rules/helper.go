package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

var headMetaArgPriority = map[string]int{"for_each": 0, "count": 0, "provider": 1}
var tailMetaArgPriority = map[string]int{"lifecycle": 0, "depends_on": 1}

// IsHeadMeta checks whether a name represents a type of head Meta arg
func IsHeadMeta(argName string) bool {
	_, isHeadMeta := headMetaArgPriority[argName]
	return isHeadMeta
}

// IsTailMeta checks whether a name represents a type of tail Meta arg
func IsTailMeta(argName string) bool {
	_, isTailMeta := tailMetaArgPriority[argName]
	return isTailMeta
}

// ComparePos compares the value of hcl.Pos pos1 and pos2,
//negative result means pos1 is prior to pos2,
//zero result means the 2 positions are identical,
//positive result means pos2 is prior to pos1
func ComparePos(pos1 hcl.Pos, pos2 hcl.Pos) int {
	if pos1.Line == pos2.Line {
		return pos1.Column - pos2.Column
	}
	return pos1.Line - pos2.Line
}

func getExistedRules() map[string]tflint.Rule {
	rules := make(map[string]tflint.Rule)
	for _, rule := range Rules {
		rules[rule.Name()] = rule
	}
	return rules
}
