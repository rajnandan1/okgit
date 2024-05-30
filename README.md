# okgit

![okgit1](https://github.com/rajnandan1/okgit/assets/16224367/53af6415-11cb-4bb1-aaaa-28b9dce0102a)

## Inspiration
I wanted to write conventional commit messages but I was too lazy to write them manually. So I created this tool to help me write conventional commits.

Use this tool to 
- Create and Commit conventional commits
- Other git chores

## Install

```bash
brew tap rajnandan1/homebrew-rajnandan
brew install okgit 
```


## Update

```bash
brew update
brew upgrade okgit
```

## Conventional Commits

Read about conventional commits [here](https://www.conventionalcommits.org)

```bash
okgit cm
```
Output 

```bash
[Required] Type (feat, fix, docs, build, chore, ci, docs, style, refactor, perf, test, others): # select what kind commit this is
[Optional] Scope: # enter the scope of the commit
[Required] Summary: 
# enter the summary of the commit
[Optional] Details: 
# enter other optional details of the commit, like author, issue number, JIRA ID etc
[Required] Breaking change? (y/n): # if this commit introduces a breaking change
[Optional] What is breaking? 
# enter what is breaking in the commit
[Optional] Footer: 
# enter the footer of the commit
```

## Documentation
 
| Command     | Description                                                             |
|-------------|-------------------------------------------------------------------------|
| ad          | Stage files for commit. Similar to `git add`                             |
| bn          | Get current branch name. Similar to `git branch`                         |
| ch          | Switch branches or restore working tree files. Similar to `git checkout` |
| cm          | Create a conventional commit. Similar to `git commit`                    |
| completion  | Generate the autocompletion script for the specified shell               |
| done        | Do add commit and push at one go                                         |
| help        | Help about any command                                                   |
| init        | Add .okgit/ to the .gitignore file                                        |
| pl          | Pull remote branch changes. Similar to `git pull`                        |
| ps          | Push local branch changes to remote. Similar to `git push`                |
| rs          | Reset changes in the working directory. Similar to `git reset`            |
| sn          | Sync local branch with remote from -> to                                 |
| st          | Check the status of the repository. Similar to `git status`               |
| start       | Start working on new or existing branch                                   |



```bash
okgit --help
```


## Development

```bash
mkdir bin
go build -o bin ./... &&  ./bin/okgit cm
```