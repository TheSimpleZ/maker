# Recursive make

This tiny program is a wrapper around GNU Make.
It allows you to run a make target that exists in any ancestor directory.

It will however stop traveling upwards at the root of any git repository you are inside of.

## Installation

Make sure you have GNU Make installed on your system already
and that it's available on your `$PATH`.

Then run:
```bash
curl 'https://i.jpillora.com/TheSimpleZ/maker!!' | bash
```

To uninstall, simply delete `/usr/local/bin/maker`.

Or download the appropriate binary for your system from the releases page and put in in your `$PATH`.

## Usage

This program should work the same way as `make`. Apart from recursively running `make` in the parent directories, all arguments are passed as-is to `make`.

Either call `maker <your_target>` or alias it to `make`:

```bash
alias make="maker"
```

_Put the above in your ~/.*rc file to make it permanent._

## Example

Clone this repo and run:

```
go build
cd example/subfolder
../../maker targetx
```

targetx exists in the makefile of the /example folder.
