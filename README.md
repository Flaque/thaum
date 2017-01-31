# Thaum

Thaum is a tiny boilerplate/templating utility. It purposely requires little to no configuration and does not require you to have anything more than a [mustache](http://mustache.github.io/) template to make a boilerplate. (Read: No script.)

# Install

If you have [go], you can run:

```
$ go get github.com/flaque/thaum
```

or, you can run:
```
$ wget "https://github.com/Flaque/thaum/releases/download/v0.1.0-beta/thaum" -P /usr/local/bin/
```

to install the binary into your path.


### If you have problems
If you try running `$ thaum` and get an error that says something like:

```
-bash: /usr/local/bin/thaum: Permission denied
```

You might [need to add permissions](http://superuser.com/questions/717663/permission-denied-when-trying-to-cd-usr-local-bin-from-terminal/717683) to your bin via:

```
$ sudo chmod -R 755 /usr/local/bin
```

And all should work.

# Features

- No configuration.
- Global templates and git included project templates.
- Compiles mustache in paths for files like `{{name}}.spec.js`.

# Usage

Your templates go in a folder called `thaum_files`. Thaum will look up the file
tree for the nearest `thaum_files` folder and use that one for your command.

## Creating a Template
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
