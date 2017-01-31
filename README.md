# Thaum

Thaum is a tiny boilerplate/templating utility. It purposely requires little to no configuration and does not require you to have anything more than a [mustache](http://mustache.github.io/) template to make a boilerplate. (Read: No script.)

# Install

At the moment, you can install either by downloading the binary from the
github releases, or using `go`, you can run:

```
go get github.com/flaque/thaum
```

# Features

- No configuration.
- Global templates and git included project templates.
- Compiles mustache in paths for files like `{{name}}.spec.js`.

# Usage

Your templates go in a folder called `thaum_files`. Thaum will look up the file
tree for the nearest `thaum_files` folder and use that one for your command.

## Setup
Create your thaum_files in your root project.

```
$ mkdir thaum_files
```

Then create your first template! This is just a folder in your `thaum_files`.

```
$ mkdir thaum_files/myTemplate
```

Then, we can create a file or a whole folder system if you want inside.

```
$ touch thaum_files/myTemplate/myWidget.js
```

In that file, you can put something like this:

```
import "allMyThings";

export class {{name}} {
  constructor(foo, bar) {
    //do something
  }
}
```

## Running Thaum

Once you have a template, you can run `thaum` like so:

```
thaum myTemplate theName
```
