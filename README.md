# GRB

A tool to simplify working with git remote branches.

## Install

```sh
go install github.com/jinzhu/grb
```

## Usage

```sh
// rename `branch1` to `branch2` and setup git tracking
grb mv `branch1` `branch2`

// rename current branch to `branch` and setup git tracking
grb mv `branch`

// add a remote repo
grb remote_add `name` `repo path`

// remove a remote repo
grb remote_rm `name`

// delete branch `branch`, default current branch
grb rm `branch`

// pull branch `branch` and setup git tracking, default current branch
grb pull `branch`

// push branch `branch` and setup git tracking, default current branch
grb push `branch`

// create new branch `branch`
grb new `branch`

// prune dead remote branches
grb prune

// display help
grb --help
```

## License

Copyright © 2009-present MIT

## Author

* Jinzhu Zhang
* http://github.com/jinzhu
* http://twitter.com/zhangjinzhu
