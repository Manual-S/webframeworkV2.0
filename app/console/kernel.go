package console

import (
	"webframeworkV2.0/framework"
	"webframeworkV2.0/framework/cobra"
)

func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "hade",    // 定义命令的关键字
		Short: "hade 命令", // 命令的简单介绍
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	return rootCmd.Execute()
}
