// Copyright 2019 Red Hat
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

package aliyun

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cmdDelete = &cobra.Command{
		Use:   "delete-image",
		Short: "Delete image on aliyun",
		Long:  `Delete an image from aliyun.`,
		RunE:  runDelete,

		SilenceUsage: true,
	}

	deleteId    string
	deleteForce bool
)

func init() {
	Aliyun.AddCommand(cmdDelete)
	cmdDelete.Flags().StringVar(&deleteId, "id", "", "image ID")
	cmdDelete.Flags().BoolVar(&deleteForce, "force", false, "force deletion")
}

func runDelete(cmd *cobra.Command, args []string) error {
	err := API.DeleteImage(deleteId, deleteForce)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't delete image: %v\n", err)
		os.Exit(1)
	}
	return nil
}
