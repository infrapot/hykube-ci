/*
 * Copyright 2024 by infrapot
 *
 * This program is a free software product. You can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License (AGPL)
 * version 3 as published by the Free Software Foundation.
 *
 * For details, see the GNU AGPL at: http://www.gnu.org/licenses/agpl-3.0.html
 */

package plan

import (
	"fmt"
	"github.com/infrapot/hykube-cli/pkg/config"
	"github.com/spf13/cobra"
)

const description = "plan Hykube resources within a namespace"

func Command(options *config.HykubeOptions) *cobra.Command {

	cmd := &cobra.Command{
		Use:          "plan",
		Short:        description,
		Long:         description,
		Example:      fmt.Sprintf("%[1]s hykube plan -n test-cloud", "kubectl"),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := options.Complete(cmd, args); err != nil {
				return err
			}

			panic("To implement")
		},
	}

	options.AddFlags(cmd.Flags())

	return cmd
}
