# appdirs

[![Go Reference](https://pkg.go.dev/badge/github.com/chasinglogic/appdirs.svg)](https://pkg.go.dev/github.com/chasinglogic/appdirs)
![CI](https://github.com/chasinglogic/appdirs/actions/workflows/ci.yml/badge.svg)

This is a port of the excellent python module named the same, which can be found
here [appdirs](https://github.com/ActiveState/appdirs) (now deprecated and
replaced by [platformdirs](https://github.com/platformdirs/platformdirs)).

The original port was written by [wessie](https://github.com/wessie) this is
a maintained fork of that project.

## The Problem

When writing desktop application, finding the right location to store user data
and configuration varies per platform. Even for single-platform apps, there
may by plenty of nuances in figuring out the right location.

For example, if running on macOS, you should use

```
~/Library/Application Support/<AppName>
```

If on Windows (at least English Win) that should be::

```
C:\Documents and Settings\<User>\Application Data\Local Settings\<AppAuthor>\<AppName>
```

or possibly::

```
C:\Documents and Settings\<User>\Application Data\<AppAuthor>\<AppName>
```

for [roaming profiles](https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-vista/cc766489(v=ws.10)) 
but that is another story.

On Linux (and other Unices), according to the 
[XDG Basedir Spec](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html), 
it should be

```
~/.local/share/<AppName>
```

## appdirs to the rescue

This library provides a simple to use API to get the appropriate storage
directories no matter the platform and what you want to store.
