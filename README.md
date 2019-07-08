# json2yaml  
Convert between json and yaml formats.  Input on stdin or by filename arguments.  

## Install
```sh
go get github.com/travisstaloch/json2yaml -tags='gen'
```

## Usage  
Usage: json2yaml [OPTION?] [FILENAME*]  
Convert between json and yaml formats.  
Input(s) supplied by stdin or as FILENAME(s).  
If format is given, all files will be converted to format.  
Otherwise .json files are converted to yaml and .yaml to json.  
  
OPTION:  
	-f --format {yaml (default), json} : examples: --f=yaml -format json  
  
examples:  
  $ cat 1.json | json2yaml  
  $ cat 1.yaml | json2yaml -format json | tee file.json  
  $ echo {"a": 1, "b": 2} | json2yaml -f yaml  
  $ printf "a: 1\\nb: 2" | json2yaml --f json  
  $ json2yaml file.json 2.yaml | tee file.yaml  
  $ json2yaml -f=json file.yaml 2.json  
  
