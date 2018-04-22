package main

import (
    git "github.com/libgit2/git2go"
//    "errors"
    "fmt"
)

var PATH string = "/Users/lin_r/Desktop/project/go/fast_manager"
var name2repo map[string]*git.Repository = make(map[string]*git.Repository)

func ShowCurrentBranch(repo *git.Repository)(string, error){
    repoReference, err := repo.Head()
    if err != nil{
        fmt.Println("ShowCurrentBranch: ", err.Error())
        return "ShowCurrentBranch", err
    }
    branch := repoReference.Branch()
    return branch.Name()
}

// func GetHeadBranch(repo *git.Repository)(string, error){
//     branchIterator, err := repo.NewBranchIterator(git.BranchLocal)
//     defer branchIterator.Free()
//     if err != nil{
//         fmt.Println(err.Error())
//     }
//     for {
//         branch, _, err := branchIterator.Next()
//         if err != nil{
//             return "No", errors.New("branch list end")
//         }
//         isHead, err := branch.IsHead()
//         if err != nil{
//             fmt.Println(err.Error())
//         }
//         if isHead{
//             branchName, err := branch.Name()
//             return branchName, err
//         }
//     }
// }
func AddFiles(idx *git.Index, pathes []string){
    for _, path := range pathes{
        idx.AddByPath(path)
    }
    idx.Write()
    idx.WriteTree()
}

func Commit(repo *git.Repository){
    pathList, err := ShowRepoStatus(repo)
    fmt.Println(pathList)
    idx, err := repo.Index()
    if err != nil{
        fmt.Println(err.Error)
    }
    AddFiles(idx, pathList)
    fmt.Println(err)
    fmt.Println("123123")
}

func ShowRepoStatus(repo *git.Repository)([]string, error){
    /*
    显示指定仓库的文件状态
    */
    unCommitFileList := make([]string, 1)
    statusOption := git.StatusOptions{}
    statusOption.Show = git.StatusShowIndexAndWorkdir
    statusOption.Flags = git.StatusOptIncludeUntracked   |
                         git.StatusOptRenamesHeadToIndex |
                         git.StatusOptSortCaseSensitively
    statusList, _:= repo.StatusList(&statusOption)
    count, _ := statusList.EntryCount()
    for i :=0; i< count; i++{
        entry, err := statusList.ByIndex(i)
        if err != nil{
            fmt.Println("fuck", err.Error())
        }
        unCommitFileList = append(unCommitFileList, entry.IndexToWorkdir.NewFile.Path)
    }
    return unCommitFileList, nil
}

func AddRepositoryRecord(repoName, Path string) error{
    reposity, err := git.OpenRepository(Path)
    if err != nil{
        fmt.Println("add repo record error", err.Error())
        return err
    }
    name2repo[repoName] = reposity
    return nil
}

func main(){
    AddRepositoryRecord("test", PATH)
    status, err := ShowRepoStatus(name2repo["test"])
    if err != nil{
        fmt.Println(err.Error())
    }
    fmt.Println(status)
    name, err := ShowCurrentBranch(name2repo["test"])
    Commit(name2repo["test"])
    fmt.Println(name, err)
}
