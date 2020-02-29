# Design-lab-1
Continuous integration

## Installation
```bash
git clone https://github.com/AnastasiaYarema/design-lab-1
cd disign-lab-1/
make
```

## Usage
Pass your postfix expression as arguments to ./out/example
```bash
./out/example 1 2 + 43 - # pass your expression after whitespace to ./out/example 
```

## Testing
```bash
make test
```

## Documentation
```bash
godoc -goroot=$(pwd)
# then open http://localhost:6060 in your browser
```

## Continuous integration build examples
Successfull builds:
- [successfull build №1](https://travis-ci.org/AnastasiaYarema/design-lab-1/builds/656036086)
- [successfull build №2](https://travis-ci.org/AnastasiaYarema/design-lab-1/builds/656638712)
- [failed build](https://travis-ci.org/AnastasiaYarema/design-lab-1/builds/656033247)
- [Pull Request build](https://travis-ci.org/AnastasiaYarema/design-lab-1/builds/656637933)
