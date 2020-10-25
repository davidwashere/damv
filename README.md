Simple utility that will move all files in subdirectories to the parent directory prefixed with the former subdirectory name

## Usage

`damv`

## Example
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
$ damv
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