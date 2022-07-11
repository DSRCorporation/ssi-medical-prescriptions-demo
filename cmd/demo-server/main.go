/*
  Copyright 2022 DSR Corporation, Denver, Colorado.
  https://www.dsr-corporation.com

  This file is part of ssi-medical-prescriptions-demo.

  ssi-medical-prescriptions-demo is free software: you can redistribute it
  and/or modify it under the terms of the GNU Affero General Public License
  as published by the Free Software Foundation, either version 3 of the License,
  or (at your option) any later version.

  ssi-medical-prescriptions-demo is distributed in the hope that it will be
  useful, but WITHOUT ANY WARRANTY; without even the implied warranty
  of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
  See the GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License along
  with ssi-medical-prescriptions-demo. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/DSRCorporation/ssi-medical-prescriptions-demo/cmd/demo-server/cmd"
)

func main() {

	rootCmd := &cobra.Command{
		Use: "demo-server",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	startCmd, err := cmd.Cmd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing cobra cmd: %v\n", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run rest server: %v\n", err)
		os.Exit(1)
	}
}
