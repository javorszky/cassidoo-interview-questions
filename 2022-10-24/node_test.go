package oct242022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_node_String(t *testing.T) {
	tests := []struct {
		name string
		node *node
		want string
	}{
		{
			name: "node with no branches",
			node: &node{
				value: 1,
				left:  nil,
				right: nil,
			},
			want: "",
		},
		{
			name: "node with one branch on the left",
			node: &node{
				value: 1,
				left: &node{
					value: 2,
				},
			},
			want: "" +
				" /  \n" +
				"  \n",
		},
		{
			name: "node with one branch on the right",
			node: &node{
				value: 1,
				right: &node{
					value: 2,
				},
			},
			want: "" +
				"  \\ \n" +
				"    \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.node.String(), "String()")
		})
	}
}
