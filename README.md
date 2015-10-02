# Toy Spelling Corrector [![Build Status][travis-svg]][travis-url] [![Coverage Status][coveralls-svg]][coveralls-url]

A rudimentary spelling corrector based on [How to Write a Spelling Corrector](http://norvig.com/spell-correct.html) by Peter Norvig. This is a direct port from the Python version with no optimizations or enhancements. Pull requests are welcome but may not be merged into the master branch. It would be interesting to see this project go both ways, meaning as concise as possible (fewest lines) and as performent as possible (fastest execution).

The original Python version is 22 lines of code and the Golang version is 88 lines of code. 

All credit goes to [Peter Norvig](http://norvig.com/) for his very educational [post](http://norvig.com/spell-correct.html) that inspired me to try my hand at porting his Python code to Golang. At the bottom of his post you can find links to other ports in many more programming languages.

### Usage

Since it's a package and not meant for real world use I recommend just cloning and testing. You can do that with the following commands:

```
git clone git@github.com:montanaflynn/toy-spelling-corrector.git
cd toy-spelling-corrector
go test
```

Additionally if you have python2 installed you can run the original python code by Peter Norvig for comparison:

```
git clone git@github.com:montanaflynn/toy-spelling-corrector.git
cd toy-spelling-corrector
python spell_test.py
```

### Benchmarks

On an old single core laptop running Ubuntu 15.04, Golang 1.4.2 and Python 2.7.9:

##### Golang

```
$ go test
n: 270, bad: 68, unknown: 15, pct: 74, secs: 13.771934394s
n: 400, bad: 130, unknown: 43, pct: 67, secs: 24.168009463s
```

##### Python

```
$ python spell_test.py 
{'bad': 68, 'bias': None, 'unknown': 15, 'secs': 22, 'pct': 74, 'n': 270}
{'bad': 130, 'bias': None, 'unknown': 43, 'secs': 40, 'pct': 67, 'n': 400}
```

[travis-url]: https://travis-ci.org/montanaflynn/toy-spelling-corrector
[travis-svg]: https://img.shields.io/travis/montanaflynn/toy-spelling-corrector.svg

[coveralls-url]: https://coveralls.io/r/montanaflynn/toy-spelling-corrector?branch=master
[coveralls-svg]: https://img.shields.io/coveralls/montanaflynn/toy-spelling-corrector.svg
