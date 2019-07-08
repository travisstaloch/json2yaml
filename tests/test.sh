# filter the output of --help and run the example commands,
# printing command and newline separators before output
json2yaml --help | grep "  $ " | awk -F'$' '{print $2}' | while read line; do printf "\n\n$line\n" ; sh -c "$line"; done
rm -f file.yaml file.json