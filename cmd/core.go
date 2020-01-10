package cmd

import(
	"fmt"
	"os"
	"io/ioutil"
	"path"
	"strconv"
	"github.com/spf13/cobra"
)

var(
	src string
	name string
	index int
)

func init() {
	RootCmd.AddCommand(renCmd)
	renCmd.Flags().StringVarP(&src, "path", "p", "", "Source directory to read from")
	renCmd.MarkFlagRequired("path")
	renCmd.Flags().StringVarP(&name, "name", "n", "", "prefix name")
	renCmd.Flags().IntVarP(&index, "index", "i", 1, "rename by this index number")
}

var RootCmd = &cobra.Command{
	Use: "itool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to use itool.")
	},
}

var renCmd = &cobra.Command{
	Use: "rename",
	Short: "rename pictures.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(src)==0 {
			RootCmd.Help()
			return
		}
		if len(name)==0 {
			name = "img_"
		}
		fmt.Println("目标文件夹："+src)
		fmt.Println("目标前缀："+name)
		fmt.Printf("起始序号：%d\n", index)
		readFile(src,name,index)
	},
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func readFile(src string, name string, i int) {
	if exists,_ := pathExists(src); exists == false {
		fmt.Println("未找到目标文件夹")
		return
	}
	files,_ := ioutil.ReadDir(src)
	for _,f := range files {
		 if f.IsDir() {
            continue
        } else {
			fileName := f.Name()
			fileSuffix := path.Ext(fileName)
			if fileSuffix==".png" || fileSuffix==".jpg" {
				newName := name+strconv.Itoa(i)+fileSuffix
				// fmt.Println(fileName)
				// fmt.Println(fileSuffix)
				// fmt.Println(newName)
				err := os.Rename(src+"/"+fileName, src+"/"+newName)
				if err != nil {
					fmt.Println("reName "+fileName+" Error", err)
					continue
				}
				i++
			}
		}
	}
	fmt.Println("执行完成！")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}