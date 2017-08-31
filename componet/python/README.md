[memory_profiler](https://github.com/fabianp/memory_profiler) is a python module for monitoring memory consumption of a process as well as line-by-line analysis of memory consumption for python programs.

## Build image

```bash
docker build -t containerops/memory_profiler .
```

## Test

```bash
docker run --rm -e CO_DATA='git-url=https://github.com/fabianp/memory_profiler.git entry-file=test/test_func.py' containerops/memory_profiler
# test with python2
docker run --rm -e CO_DATA='git-url=https://github.com/fabianp/memory_profiler.git entry-file=test/test_func.py version=python' containerops/memory_profiler
```

## CO_DATA arguments

* `git-url` is the source git repo url
* `version` is one of `python`, `python2`, `python3`, `py3k`.  default is `py3k`
* `entry-file` is the entry file for memory profile
