// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ory/hydra/cmd/server"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Serves Administrative HTTP/2 APIs",
	Long: `This command opens one port and listens to HTTP/2 API requests. The exposed API handles administrative
requests like managing OAuth 2.0 Clients, JSON Web Keys, login and consent sessions, and others.

This command is configurable using the same options available to "serve public" and "serve all".

It is generally recommended to use this command only if you require granular control over the administrative and public APIs.
For example, you might want to run different TLS certificates or CORS settings on the public and administrative API.

This command does not work with the "memory" database. Both services (administrative, public) MUST use the same database
connection to be able to synchronize.

` + serveControls,
	Run: server.RunServeAdmin(Version, Commit, Date),
}

func init() {
	serveCmd.AddCommand(adminCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
