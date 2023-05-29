package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

var str string
var mode int8
var password bool

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
	rootCmd.AddCommand(wordCmd)
	passwordCmd.Flags().BoolVarP(&password, "password", "p", false, "直接获取一个密码")
	rootCmd.AddCommand(passwordCmd)
}
