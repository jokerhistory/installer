## BUILD IMAGE

```bash
docker build -t containerops/coala -f ./analysis/coala/Dockerfile ./analysis/coala
docker build -t containerops/flake8 -f ./analysis/flake8/Dockerfile ./analysis/flake8
docker build -t containerops/line_profiler -f ./analysis/line_profiler/Dockerfile ./analysis/line_profiler
docker build -t containerops/memory_profiler -f ./analysis/memory_profiler/Dockerfile ./analysis/memory_profiler
docker build -t containerops/pep8 -f ./analysis/pep8/Dockerfile ./analysis/pep8
docker build -t containerops/pycallgraph -f ./analysis/pycallgraph/Dockerfile ./analysis/pycallgraph
docker build -t containerops/pylama -f ./analysis/pylama/Dockerfile ./analysis/pylama
docker build -t containerops/pylint -f ./analysis/pylint/Dockerfile ./analysis/pylint
docker build -t containerops/dh-virtualenv -f ./build/dh-virtualenv/Dockerfile ./build/dh-virtualenv
docker build -t containerops/nuitka -f ./build/nuitka/Dockerfile ./build/nuitka
docker build -t containerops/pybuilder -f ./build/pybuilder/Dockerfile ./build/pybuilder
docker build -t containerops/pyinstaller -f ./build/pyinstaller/Dockerfile ./build/pyinstaller
docker build -t containerops/pynsist -f ./build/pynsist/Dockerfile ./build/pynsist
docker build -t containerops/mkdocs -f ./document/mkdocs/Dockerfile ./document/mkdocs
docker build -t containerops/pdoc -f ./document/pdoc/Dockerfile ./document/pdoc
docker build -t containerops/pycco -f ./document/pycco/Dockerfile ./document/pycco
docker build -t containerops/sphinx -f ./document/sphinx/Dockerfile ./document/sphinx
docker build -t containerops/coverage -f ./test/coverage/Dockerfile ./test/coverage
docker build -t containerops/green -f ./test/green/Dockerfile ./test/green
docker build -t containerops/mamba -f ./test/mamba/Dockerfile ./test/mamba
docker build -t containerops/nose -f ./test/nose/Dockerfile ./test/nose
docker build -t containerops/nose2 -f ./test/nose2/Dockerfile ./test/nose2
docker build -t containerops/pytest -f ./test/pytest/Dockerfile ./test/pytest
docker build -t containerops/tox -f ./test/tox/Dockerfile ./test/tox
docker build -t containerops/unittest -f ./test/unittest/Dockerfile ./test/unittest
```

## RUN

```bash
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git' containerops/coala
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git' containerops/flake8
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git version=python' containerops/flake8
docker run --rm -e CO_DATA='git-url=https://github.com/istrategylabs/python-profiling entry-file=debug.py' containerops/line_profiler
docker run --rm -e CO_DATA='git-url=https://github.com/istrategylabs/python-profiling entry-file=debug.py version=python' containerops/line_profiler
docker run --rm -e CO_DATA='git-url=https://github.com/fabianp/memory_profiler.git entry-file=test/test_func.py' containerops/memory_profiler
docker run --rm -e CO_DATA='git-url=https://github.com/fabianp/memory_profiler.git entry-file=test/test_func.py version=python' containerops/memory_profiler
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git' containerops/pep8
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git version=python' containerops/pep8
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/bpnn.git entry-file=bpnn.py upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/pycallgraph' containerops/pycallgraph
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/bpnn.git entry-file=bpnn.py upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/pycallgraph version=python' containerops/pycallgraph
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git' containerops/pylama
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git version=python' containerops/pylama
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git' containerops/pylint
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/python-aio-periodic.git version=python' containerops/pylint
docker run --rm -e CO_DATA='git-url=https://github.com/spotify/dh-virtualenv.git upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/dh-virtualenv' containerops/dh-virtualenv
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/bpnn.git entry-file=bpnn.py upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/nuitka' containerops/nuitka
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/bpnn.git entry-file=bpnn.py upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/nuitka version=python' containerops/nuitka
docker run --rm -e CO_DATA='git-url=https://github.com/blanzp/amazon_examples.git entry-path=. task=run_unit_tests' containerops/pybuilder
docker run --rm -e CO_DATA='git-url=https://github.com/blanzp/amazon_examples.git entry-path=. task=run_unit_tests version=python' containerops/pybuilder
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/bpnn.git entry-file=bpnn.py upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/pyinstaller' containerops/pyinstaller
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/bpnn.git entry-file=bpnn.py upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/pyinstaller version=python' containerops/pyinstaller
docker run --rm -e CO_DATA='git-url=https://github.com/takluyver/pynsist.git entry-file=examples/console/installer.cfg upload=https://hub.opshub.sh/binary/v1/containerops/component/binary/v0.1/pynsist' containerops/pynsist
docker run --rm -e CO_DATA='git-url=https://github.com/mkdocs/mkdocs.git entry-path=.' containerops/mkdocs
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/grapy.git entry-mod=grapy' containerops/pdoc
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/grapy.git entry-mod=grapy version=python' containerops/pdoc
docker run --rm -e CO_DATA='git-url=https://github.com/pycco-docs/pycco.git' containerops/pycco
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/grapy.git entry-path=docs' containerops/sphinx
docker run --rm -e CO_DATA='git-url=https://github.com/Lupino/grapy.git entry-path=docs version=python' containerops/sphinx
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=test/test_regex.py' containerops/coverage
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=test/test_regex.py version=python' containerops/coverage
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=.' containerops/green
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=. version=python' containerops/green
docker run --rm -e CO_DATA='git-url=https://github.com/juanAFernandez/testing-with-python.git entry-file=examples/mamba_example.py' containerops/mamba
docker run --rm -e CO_DATA='git-url=https://github.com/juanAFernandez/testing-with-python.git entry-file=examples/mamba_example.py version=python' containerops/mamba
docker run --rm -e CO_DATA='git-url=https://github.com/nose-devs/nose.git entry-path=unit_tests' containerops/nose
docker run --rm -e CO_DATA='git-url=https://github.com/nose-devs/nose.git entry-path=unit_tests version=python' containerops/nose
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=.' containerops/nose2
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=. version=python' containerops/nose2
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=.' containerops/pytest
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-path=. version=python' containerops/pytest
docker run --rm -e CO_DATA='git-url=https://github.com/CleanCut/green.git entry-path=.' containerops/tox
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-module=test.test_regex' containerops/unittest
docker run --rm -e CO_DATA='git-url=https://github.com/minhhh/regex.git entry-module=test.test_regex version=python' containerops/unittest
```
