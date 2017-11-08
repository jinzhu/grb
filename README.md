# GRB

A tool to simplify working with remote branches.

## Install

```sh
go install github.com/jinzhu/grb
```

## Usage

* rename `branch1` to `branch2`
  $ grb mv   [branch1] [branch2] [--explain]
* rename current branch to `branch`
  $ grb mv branch [--explain]
* add a remote repo
  $ grb remote_add `name` `repo path` [--explain]
* remove a remote repo
  $ grb remote_rm `name` [--explain]
* delete branch `branch`,default current_branch
  $ grb rm [branch] [--explain]
* pull branch `branch`,default current_branch
  $ grb pull [branch] [--explain]
* push branch `branch`, default current_branch
  $ grb push [branch] [--explain]
* create new branch `branch`
  $ grb new [branch] [--explain]
* prune dead remote branches
  $ grb prune


## License

Copyright © 2009-present MIT

## Author

* Jinzhu Zhang
* http://github.com/jinzhu
* http://twitter.com/zhangjinzhu
