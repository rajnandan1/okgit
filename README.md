# okgit

![okgit1](https://github.com/rajnandan1/okgit/assets/16224367/53af6415-11cb-4bb1-aaaa-28b9dce0102a)

## Inspiration
I wanted to write conventional commit messages but I was too lazy to write them manually. So I created this tool to help me write conventional commits.

Use this tool to 
- Create and Commit conventional commits
- Other git chores

## Install

```bash
go install -v github.com/rajnandan1/okgit@latest
```

## Initialize

Run this inside a git repository 

```bash
okgit init
```
## Update

```bash
okgit update
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

## Setup
Update you `.gitignore` file to ignore the `.okgit` directory.

```
echo ".okgit/" >> .gitignore
```

## Run locally without install

```bash
go build -o bin ./... &&  ./bin/okgit cm
```