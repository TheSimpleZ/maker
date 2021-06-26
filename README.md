# Recursive make

This tiny program is a wrapper around GNU Make.
It allows you to run a make target that exists in a parent directory.

## Installation

Make sure you have GNU Make installed on your system already
and that it's available on your `$PATH`.

Then run:
```bash
curl https://i.jpillora.com/TheSimpleZ/maker! | bash
```

Or download the appropriate binary for your system from the releases page and put in in your `$PATH`.

## Usage

This program should work the same way as `make`. Apart from recursively running `make` in the parent directories, all arguments are passed as-is to `make`.

Either call `maker <your_target>` or alias it to `make`:

```bash
alias make="maker"
```

_Put the above in your ~/.*rc file to make it permanent._