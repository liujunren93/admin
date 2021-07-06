/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"github.com/liujunren93/admin/server"
	"github.com/liujunren93/admin/view/page/api"
	"github.com/liujunren93/admin/view/page/info"
	"github.com/liujunren93/admin/view/page/table"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A brief description of your application",
	//	Long: `A longer description that spans multiple lines and likely contains
	//examples and usage of using your application. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		source := cmd.Flag("source").Value.String()
		if source == "" {
			return errors.New("请指定源文件")
		}
		var data []*core.Group
		if strings.Index(source, ".go") >= 0 {
			file := core.ParseFile(source)
			data = append(data, file)
		} else {
			data = core.ParsePath(source)

		}
		root := cmd.Flag("api")
		if root.Value.String() != "" {

			if cmd.Flag("mod").Value.String() != "" {
				abs, err := filepath.Abs(root.Value.String() )

				if err != nil {
					return err
				}
				global.ApiRoot = abs

				global.Mod = cmd.Flag("mod").Value.String()
				server.NewCtrl(data...)
				server.NewDao(data...)
				server.NewEntity(data...)
				server.NewConfigFile()
				server.NewRoute(data...)
				server.NewDBFile()
				server.NewUtils()
			} else {

				return errors.New("请设置 go mod")
			}
		}
		if cmd.Flag("web")!=nil&&cmd.Flag("web").Value.String() != "" {
			abs, err := filepath.Abs(cmd.Flag("web").Value.String() + "/")
			if err != nil {
				return err
			}
			global.WebRoot = abs
			api.BuildPage(data)
			table.BuildPage(data)
			info.BuildPage(data)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String("source", "", "源文件或源文件目录")
	rootCmd.PersistentFlags().String("api", "", "server api project path")
	rootCmd.PersistentFlags().String("mod", "", "server api go mod")
	rootCmd.PersistentFlags().String("web", "", "web path")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Find home directory.
		home, err := os.UserHomeDir()

		cobra.CheckErr(err)

		// Search config in home directory with name ".cmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
