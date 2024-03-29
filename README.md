# json2yaml  
Convert between json and yaml formats.  Input on stdin or by filename arguments.  

## Install
```sh
go get github.com/travisstaloch/json2yaml
```

## Run
```sh
json2yaml 1.json # prints equivalent yaml
```

## Test
There are a few simple json and yaml files in the tests directory.  
Run the examples in Usage (below):  
```sh
cd $GOPATH/github.com/travisstaloch/json2yaml/tests
sh tests.sh
```


## Generate
`json2yaml --help` prints usage which is generated from the Usage text  
(below).  usage.go can be generated with:  
```sh
cd $GOPATH/github.com/travisstaloch/json2yaml
go generate
go install # apply changes to json2yaml executable
```

## Usage  
Usage: json2yaml [OPTION?] [FILENAME*]  
Convert between json and yaml formats.  
Input(s) supplied by stdin or as FILENAME(s).  
If format is given, all files will be converted to format.  
Otherwise .json files are converted to yaml and .yaml to json.  
  
OPTION:  
	-f --format {yaml (default), json} : examples: --f=yaml -format json  
  --help : show this text  
  
examples:  
  $ cat 1.json | json2yaml  
  $ cat 1.yaml | json2yaml -format json | tee file.json  
  $ echo {"a": 1, "b": 2} | json2yaml -f yaml  
  $ printf "a: 1\\nb: 2" | json2yaml --f json  
  $ json2yaml file.json 2.json | tee file.yaml  
  $ json2yaml -f=json file.yaml 2.yaml  
  