package main

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"path/filepath"
	"regexp"
)

const HELP = `Usage: photo-organization-tool target standard
使い方: photo-organization-tool 対象ディレクトリ 基準ディレクトリ`

type MvPath struct {
	old_path string
	new_path string
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("引数の数が間違っています")
		fmt.Println(HELP)
		os.Exit(1)
	}

	target_dir_path := args[1]
	standard_dir_path := args[2]

	target_files, err := GetFilesInDirectory(target_dir_path, "対象")
	if err != nil {
		fmt.Println(err)
		fmt.Println(HELP)
		os.Exit(1)
	}

	standard_files, err := GetFilesInDirectory(standard_dir_path, "基準")
	if err != nil {
		fmt.Println(err)
		fmt.Println(HELP)
		os.Exit(1)
	}

	var standard_files_list []string
	extension_regex := regexp.MustCompile(`(.+?)\..+`)
	for _, file := range standard_files {
		if !file.IsDir() {
			standard_files_list = append(standard_files_list, extension_regex.ReplaceAllString(file.Name(), "$1"))
		}
	}

	tmp_dir_path := filepath.Join(target_dir_path, "tmp")
	if err := os.MkdirAll(tmp_dir_path, 0777); err != nil {
		fmt.Println("ディレクトリの作成に失敗しました")
		os.Exit(1)
	}

	var target_paths []MvPath
	for _, file := range target_files {
		file_name := file.Name()
		name := extension_regex.ReplaceAllString(file_name, "$1")
		if !(file.IsDir() || slices.Contains(standard_files_list, name)) {
			target_path := MvPath{old_path: filepath.Join(target_dir_path, file_name), new_path: filepath.Join(tmp_dir_path, file_name)}

			target_paths = append(target_paths, target_path)
			fmt.Println(target_path.old_path, "->", target_path.new_path)
		}
	}

	if len(target_paths) == 0 {
		fmt.Println("対象となるファイルが存在しません")
		os.Exit(0)
	}

	fmt.Println("よろしいですか？[y/N]")
	var confirmation string
	fmt.Scan(&confirmation)
	if confirmation == "y" || confirmation == "Y" {
		for _,path := range target_paths{
		if err := os.Rename(path.old_path, path.new_path); err != nil {
			fmt.Println("処理に失敗しました")
			os.Exit(1)
		}
		}
	} else {
		fmt.Println("操作を中止します")
		os.Exit(0)
	}

	fmt.Println("一時ディレクトリにファイルを移動しました。削除も行いますか？[y/N]")
	var confirmation2 string
	fmt.Scan(&confirmation2)
	if confirmation2 == "y" || confirmation2 == "Y" {
		if err:=os.RemoveAll(tmp_dir_path);err !=nil{
			fmt.Println("削除に失敗しました")
			os.Exit(1)
		}
		fmt.Println("一時ディレクトリを削除しました")
		os.Exit(0)
	} else {
		fmt.Println("削除せず終了します")
		os.Exit(0)
	}
}

func GetFilesInDirectory(path string, name string) ([]os.DirEntry, error) {
	target, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintln(name, "ディレクトリの読み込みに失敗しました: ", err))
	}

	target_stat, err := target.Stat()
	if err != nil {
		return nil, errors.New(fmt.Sprintln(name, "ディレクトリの読み込みに失敗しました: ", err))
	}

	if !target_stat.IsDir() {
		return nil, errors.New(fmt.Sprintln(name, "がディレクトリではありません"))
	}

	target_files, err := target.ReadDir(0)
	if err != nil {
		return nil, errors.New(fmt.Sprintln(name, "ディレクトリの中のファイル一覧の取得に失敗しました: ", err))
	}

	defer target.Close()

	return target_files, nil
}
