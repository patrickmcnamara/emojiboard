# Emojiboard

![latest version](https://img.shields.io/github/v/tag/patrickmcnamara/emojiboard?label=latest%20version)
![last commit](https://img.shields.io/github/last-commit/patrickmcnamara/emojiboard)
![top language](https://img.shields.io/github/languages/top/patrickmcnamara/emojiboard)
![licence](https://img.shields.io/github/license/patrickmcnamara/emojiboard?label=licence)

Emojiboard is an emoji keyboard for GNU/Linux machines running X11. 🖥

## Dependencies

- [dmenu](https://tools.suckless.org/dmenu/)
- [xdotool](https://www.semicomplete.com/projects/xdotool/)

## Installation

Run `go get -u github.com/patrickmcnamara/emojiboard`.

## Usage

Run `emojiboard` after installation, assuming that your `$GOPATH` is in your `$PATH`.
You should probably add it to a hotkey for convenience.

After running it, type in the name of an emoji you want.
For example, if you want this "🤪" emoji, you can type in "Zany Face" and press `Enter`.

You can join emoji together to get other emoji.

| Joined emoji                                                  | Result emoji                      |
|---------------------------------------------------------------|-----------------------------------|
| Mage (🧙) + Dark Skin Tone (🏿)                               | Dark Skin Mage (🧙🏿)               |
| Woman (👩) + Light Skin Tone (🏻) + ZWJ + School (🏫)         | Light Skin Woman Teacher (👩🏻‍🏫)     |
| Man (👨) + ZWJ + Man (👨) + ZWJ + Girl (👧) + ZWJ + Girl (👧) | Family: Man, Man, Girl, Girl (👨‍👨‍👧‍👧) |

Some of these are already available directly, but not all of them.

# Licence

This project is licenced under the European Union Public Licence v1.2.
