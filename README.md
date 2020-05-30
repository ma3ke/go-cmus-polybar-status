# go-cmus-polybar-status
A verbosely named go program that outputs a string with cmus player status, title, and progress, which can be used to produce a polybar module which shows the cmus player status.

## Installation

### Prerequisites
In order to install this program, it is required that you have installed [`go`](golang.org) and that your `$GOPATH` is in your `$PATH` shell variable.

You can check whether you have your `$GOPATH` in your `path`, by trying 

```zsh
echo $PATH
```

You are ready to go if `/home/username/go/bin` is among the output. Else, check out [this Arch Wiki page](https://wiki.archlinux.org/index.php/Go#$GOPATH) (also applies to other distrobutions.)

Also, the executable prints a single-line string to std out. This makes this program suited for other bars and even other applications than polybar! 

### Testing
You can quickly test whether everything is working, by running `go run main.go` inside of the project directory. If no errors occur and the output is to your satisfaction, you can proceed to installing it :).

### Installing the executable
To install the executable, you run the command `go install` inside of the project directory. This will build the program, and place it inside of your `$GOPATH/bin/` folder (typically at `~/go/bin/go-cmus-polybar-status`). 

### Adding the module to polybar
The following snippet can be added to your `.config/polybar/config` file. 

```toml
[module/cmus-status]
type = custom/exec
exec = ~/go/bin/go-cmus-polybar-status || echo "error!"
interval = 1
```

(Updating the module every second should be fine and does look the best. But, if you seek to minimize the overhead of your polybay, you could always increase the interval value. However, if polybar's performance is of great concern to you, it might not be the right choice.)

And to add the module to your bar, you can add it to one of the `modules` declarations. I have mine in the `modules-left` block.

```toml
modules-left = bspwm alsa cmus-status
```

## Configuration
There is no configuration file, and all configuration must be done by actually modifying the source code and recompiling. 

***But don't worry!*** This is very easy to do and does not require much knowledge of Go. For example, if you want to rearrange or exclude elements in the status string, you can simple rearrange the variables in the `output` variable in the `main()` function.

### Element arrangement
```go
// Default
output := []string {ind, pos, prog, dur, disp}
// Returns: '>  04:20  ------|---  06:23  Artist - Title'

// Smaller, no progress bar, and rearranged
output := []string {ind, disp, pos, "/", dur}
// Returns: '>  Artist - Title  04:20  /  06:23'
```

### Divider and filler characters
At various places inside `main.go`, there are variables where a single character is stored. These are used by the program to create the line (`-`) and indicator (`|`) for the progress bar, for example. You can also change the spacing between the elements, by modifying the `sep` variable inside of `main.go`.

Just play around with them, if you like.

### Recompiling
After making changes to the program, you can test it using `go run main.go` in the project directory. If no errors are returned, and the output is satisfactory, you can recompile the module with `go install`. Just like the installation!

