// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"github.com/pkg/errors"
)

// cwdCmd represents the cwd command
var cwdCmd = &cobra.Command{
	Use:   "cwd",
	Short: "Change current directory",
	Long: `Changes program current directory
and prints it out.`,
	Example: "golang-test-cli cwd dir",
}

func init() {
	RootCmd.AddCommand(cwdCmd)

	dir := cwdCmd.Flags().StringP("dir", "d", "", "Directory to change current working directory to")

	cwdCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if (dir == nil || *dir == "") && len(args) > 0 {
			dir = &args[0]
		}

		if dir == nil || *dir == "" {
			return errors.New("dir is required argument or an --directory option")
		}

		if err := os.Chdir(*dir); err != nil {
			return err
		}

		dir, err := filepath.Abs(filepath.Dir("."))
		if err != nil {
			return err
		}

		fmt.Println("pwd:", dir)

		return nil
	}
}
