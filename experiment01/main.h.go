package main

import (
	"plugin"
)

const BOOTSTRAP_LEVEL_0 = 0
const BOOTSTRAP_LEVEL_1 = 1
const BOOTSTRAP_LEVEL_2 = 2
const BOOTSTRAP_LEVEL_3 = 3

type Boot func(int)
type BootAll func()
type RunLevel2 func()
type RunLevel3 func()
type SetRoute func()
