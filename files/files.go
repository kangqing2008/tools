/*
文件操作相关的工具类
@康庆 2016-11-24
*/
package files

import (
	"path/filepath"
	"os"
	"os/exec"
	"strings"
	"sort"
	"io"
	"fmt"
	"bufio"
	"kangqing2008/tools"
)

const StrPathSeparator string = string(os.PathSeparator)

//获取当前运行的文件的路径，并将路径中的\替换成为/
func ExecFilePath()string{
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return strings.Replace(path, "\\", "/", -1)
}

//获取当前运行的文件所在的目录
func CurrentPath()string{
	exe := ExecFilePath()
	return tools.Substr(exe, 0, strings.LastIndex(exe, "/"))

}


// 判断文件是否存在
func Exists(filename string)(bool,os.FileInfo){
	info,err := os.Stat(filename)
	return err == nil || os.IsExist(err),info
}

// 罗列目录下的文件清单
func ListFiles(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

//获取制定的文件所在的目录
//只针对文件
func GetFileDir(filename string)string{
	filename = strings.Replace(filename,"\\","/",-1)
	if strings.LastIndex(filename,"/") > -1{
		return tools.Substr(filename, 0, strings.LastIndex(filename, "/"))
	}else{
		return filename
	}
}

//根据文件全路径，获取制定的文件的名称和后缀
//返回 filename和suffix
//只针对文件
func GetFileName(fullpath string)(string,string){
	fullpath = strings.Replace(fullpath,"\\","/",-1)
	var fullname,filename,suffix string
	// 找到最后一个/
	if strings.LastIndex(fullpath,"/") > -1{
		fullname = tools.SubstrToEnd(fullpath, strings.LastIndex(fullpath, "/") + 1)
	}else{
		fullname = fullpath
	}
	// 找到最后一个.
	if strings.LastIndex(fullname,".") > -1{
		filename = tools.Substr(fullname,0,strings.LastIndex(fullname,"."))
		suffix = tools.SubstrToEnd(fullname,strings.LastIndex(fullname,".") + 1)
	}else{
		filename = fullname
		suffix = ""
	}
	return filename,suffix
}

//拷贝文件
func CopyFile(src,dest string)(int64,error){
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()
	desFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()
	return io.Copy(desFile, srcFile)
}

// 读取文件中的所有文本，去掉换行符后组装成数组
func ReadAllLines(filename string)([]string,error){
	f, err := os.Open(filename)
	if err != nil {
		return nil,err
	}
	var results []string
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				return results,nil
			}else{
				fmt.Println("读取文件出错：",filename)
			}
			return results,err
		}else{
			results = append(results,line)
		}
	}
	return results,nil
}