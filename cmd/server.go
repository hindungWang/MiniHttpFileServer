// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"MiniHttpFileServer/pkg/server"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a http sever",
	Long: `This command need two args, run like this: server -d . -p 8080.
           It can expose your path just like nginx http file index.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := &server.Server{
			Path: Path,
			Port: Port,
		}
		err := s.Start()
		if err != nil {
			log.Error(err)
		}
	},
}

var Path string
var Port string

func init() {
	serverCmd.Flags().StringVarP(&Path, "path", "d", "", "Path of file (required)")
	serverCmd.Flags().StringVarP(&Port, "port", "p", "", "server listen port (required)")
	serverCmd.MarkFlagRequired("port")
	serverCmd.MarkFlagRequired("path")
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
