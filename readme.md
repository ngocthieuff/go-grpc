Guide for arch linux:

error:
```console
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
```

solution:
```console
git clone https://aur.archlinux.org/protoc-gen-go.git
ls
makepkg -si
```