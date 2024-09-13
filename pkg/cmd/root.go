/*
 * Copyright 2024 by infrapot
 *
 * This program is a free software product. You can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License (AGPL)
 * version 3 as published by the Free Software Foundation.
 *
 * For details, see the GNU AGPL at: http://www.gnu.org/licenses/agpl-3.0.html
 */

package cmd

import (
	"fmt"
	"github.com/infrapot/hykube-cli/pkg/cmd/plan"
	"github.com/infrapot/hykube-cli/pkg/config"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

const hykubeExample = `
	# plan resources in a given namespace
	%[1]s hykube plan -n test-cloud

	# apply the changes for resources in a given namespace
	%[1]s hykube deploy -n test-cloud

	# switch your current-context to one that contains the desired namespace
	%[1]s hykube destroy -n test-cloud
`

func NewRootCmd(streams genericiooptions.IOStreams) *cobra.Command {
	o := config.NewHykubeOptions(streams)

	cmd := &cobra.Command{
		Use:          "hykube",
		Short:        "Invoke operations on Hykube resources",
		Example:      fmt.Sprintf(hykubeExample, "kubectl"),
		SilenceUsage: true,
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "kubectl hykube",
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(
		plan.Command(o))

	return cmd
}
