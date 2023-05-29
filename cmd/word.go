package cmd

import (
	"github.com/spf13/cobra"
	"go-learn/internal/word"
	"log"
	"strings"
	"time"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂时不支持该格式转换，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果为：%s", content)
	},
}

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "获取密码",
	Long:  "获取一个秘密",
	Run: func(cmd *cobra.Command, args []string) {
		if password {
			password := time.Now().Unix()
			log.Printf("密码为：%v", password)
		} else {
			log.Println("获取密码失败！")
		}
	},
}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下划线单词转为大驼峰单词",
	"4：下划线单词转为小驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")
