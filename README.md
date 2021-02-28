# dise
[![Build](https://github.com/isso0424/dise/actions/workflows/build.yml/badge.svg)](https://github.com/isso0424/dise/actions/workflows/build.yml)
[![Lint](https://github.com/isso0424/dise/actions/workflows/lint.yml/badge.svg)](https://github.com/isso0424/dise/actions/workflows/lint.yml)
[![Test](https://github.com/isso0424/dise/actions/workflows/test.yml/badge.svg)](https://github.com/isso0424/dise/actions/workflows/test.yml)
[![Issue](https://img.shields.io/github/issues/isso0424/dise)](https://github.com/isso0424/dise/issues)
![Language](https://img.shields.io/github/languages/top/isso0424/dise)
![License](https://img.shields.io/github/license/isso0424/dise)
## About
Dice bot for discord.

## Usage
### !roll
Roll dices.  
You can roll multiple dices and get sum for the result.
```
!roll [dice template] [dice template] ...
```

#### dice template
When dice template is nDm, you roll n m-sided dices.

### !judge
Judge with 100-sided dice.
```
!judge <target>
```

#### target
Target is successful line for the judge.  
If result is equal or under to target, it is successful.
