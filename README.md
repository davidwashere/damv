# DAMV
Simple utility that helps with renaming files

## Install

`go install github.com/davidwashere/damv`

## Usage
Use the contextual help for details

```
$ damv help
Moves and renames files

Usage:
  damv [command]

Available Commands:
  help        Help about any command
  prefix      Adds prefix to files in current directory
  seq         Renames files with a base name + sequence number
  subdir      Moves files in subdirs to current dir while prefixing subdir to filename

Flags:
  -h, --help   help for damv

Use "damv [command] --help" for more information about a command.
```

## Example: subdir
Assuming:
```
dir1/
  file1.txt
dir2/
  file2a.txt
  file2b.txt
dir3/
  file3.txt
```

Running:
```
$ damv subdir
Pending moves:

  dir1\file1.txt   => dir1.file1.txt
  dir2\file2.txt   => dir2.file2.txt
  dir2\file2x1.txt => dir2.file2x1.txt
  dir3\file3.txt   => dir3.file3.txt

[Enter] to execute rename, typing anything else aborts: q
```

Will yield:
```
dir1.file1.txt
dir2.file2a.txt
dir2.file2b.txt
dir3.file3.txt
```

## Example: prefix
Assuming:
```
file1.txt
file2.txt
```

Running:
```
$ damv prefix hello.
Pending moves:

  file1.txt => hello.file1.txt
  file2.txt => hello.file2.txt

[Enter] to continue, anything else aborts:
```

Will yield:
```
hello.file1.txt
hello.file2.txt
```

## Example: prefix -r
Assuming:
```
file1.txt
file2.txt
```

Running:
```
$ damv prefix hello -r file
Pending moves:

  file1.txt => hello1.txt
  file2.txt => hello2.txt

[Enter] to continue, anything else aborts:
```

Will yield:
```
hello1.txt
hello2.txt
```

## Example: seq
Assuming:
```
hello.txt
world.txt
hola.txt
adios.txt
```

Running:
```
$ damv seq goodbye -s 1
Pending moves:

  adios.txt => goodbye1.txt
  hello.txt => goodbye2.txt
  hola.txt  => goodbye3.txt
  world.txt => goodbye4.txt

[Enter] to continue, anything else aborts:
```

Will yield:
```
goodbye1.txt
goodbye2.txt
goodbye3.txt
goodbye4.txt
```