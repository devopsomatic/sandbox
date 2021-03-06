/*
Based on Minikube (https://github.com/kubernetes/minikube), which is licensed under The Apache License, Version 2.0:

Copyright 2016 The Kubernetes Authors All rights reserved.
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
	//goflag "flag"
	//"io/ioutil"

	"os"
	"strings"

	//"github.com/docker/machine/libmachine/log"
	"github.com/spf13/cobra"
	//"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stackfoundation/sandbox/core/pkg/minikube/config"
	"github.com/stackfoundation/sandbox/core/pkg/minikube/constants"
	"github.com/stackfoundation/sandbox/log"
	//"github.com/golang/glog"
)

var debug bool
var dirs = [...]string{
	constants.GetMinipath(),
	constants.MakeMiniPath("certs"),
	constants.MakeMiniPath("machines"),
	constants.MakeMiniPath("cache"),
	constants.MakeMiniPath("cache", "iso"),
	constants.MakeMiniPath("cache", "localkube"),
	constants.MakeMiniPath("config"),
	constants.MakeMiniPath("addons"),
	constants.MakeMiniPath("logs"),
}
var originalCommand string
var proxyArg string

var viperWhiteList = []string{
	"v",
	"alsologtostderr",
	"log_dir",
}

var RootCmd = &cobra.Command{
	Use:   "sbox",
	Short: "Sandbox is a tool that runs Docker-based workflows.",
	Long: `Sandbox is a tool that runs Docker-based workflows.
For help, visit https://stack.foundation/docs`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetDebug(true)
		}

		if len(originalCommand) > 0 {
			cmd.Use = originalCommand
		}
		for _, path := range dirs {
			if err := os.MkdirAll(path, 0777); err != nil {
				//glog.Exitf("Error creating sbox directory: %s", err)
			}
		}

		// Log level 3 or greater enables libmachine logs
		//if !glog.V(3) {
		//        log.SetOutWriter(ioutil.Discard)
		//        log.SetErrWriter(ioutil.Discard)
		//}
		//
		//// Log level 7 or greater enables debug level logs
		//if glog.V(7) {
		//        log.SetDebug(true)
		//}

		//logDir := pflag.Lookup("log_dir")
		//if !logDir.Changed {
		//        logDir.Value.Set(constants.MakeMiniPath("logs"))
		//}

	},
}

func Execute() {
	_ = RootCmd.Execute()
}

// Handle config values for flags used in external packages (e.g. glog)
// by setting them directly, using values from viper when not passed in as args
func setFlagsUsingViper() {
	//for _, config := range viperWhiteList {
	//        var a = pflag.Lookup(config)
	//        viper.SetDefault(a.Name, a.DefValue)
	//        // If the flag is set, override viper value
	//        if a.Changed {
	//                viper.Set(a.Name, a.Value.String())
	//        }
	//        // Viper will give precedence first to calls to the Set command,
	//        // then to values from the config.yml
	//        a.Value.Set(viper.GetString(a.Name))
	//        a.Changed = true
	//}
}

func init() {
	RootCmd.PersistentFlags().StringP(config.MachineProfile, "p", constants.DefaultMachineName, `The name of the VM being used.
	This can be modified to allow for multiple VMs to be run independently`)

	//pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Turn on debug logging")
	RootCmd.PersistentFlags().StringVar(&proxyArg, "proxy", "", "Specify a proxy to use")

	RootCmd.PersistentFlags().StringVar(&originalCommand, "original-command", "", "Original path of sbox command")
	RootCmd.PersistentFlags().MarkHidden("original-command")

	cobra.OnInitialize(initConfig)
	viper.BindPFlags(RootCmd.PersistentFlags())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configPath := constants.ConfigFile
	viper.SetConfigFile(configPath)
	viper.SetConfigType("properties")
	err := viper.ReadInConfig()
	if err != nil {
		//glog.Warningf("Error reading config file at %s: %s", configPath, err)
	}
	setupViper()
}

func setupViper() {
	viper.SetEnvPrefix(constants.MinikubeEnvPrefix)
	// Replaces '-' in flags with '_' in env variables
	// e.g. iso-url => $ENVPREFIX_ISO_URL
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	setFlagsUsingViper()
}
