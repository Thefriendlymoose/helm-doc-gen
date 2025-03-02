helm-doc-gen
============

## Usage
Only works for the values.yaml file.
```yaml
test:
    test1:
    # @params @descr blablabla
    test2: true
```

above will generate
| path | type | default | description |
| ---- | ---- | ------- | ----------- |
| test.test2 | boolean | true | blablabla |

## Available commands
```yaml
# @params
# @params @type
# @params @descr
# @params @type @descr
```

## Limitations
```yaml
# @params
test:
    test1:
    # @params @descr blablabla
    # @params @type string
    test2: true
    # @params
    test3: {}
    # @params @type string:int @descr blabla this is a map
    test4: {}
```

above will generate
| path | type | default | description |
| ---- | ---- | ------- | ----------- |
| test.test2 | boolean | true | blablabla |
| test.test2 | string | true |  |
| test.test3 | string | {} |  |
| test.test4 | string:int | {} | blabla this is a map |

Will generate multiple tablerows if more comments with @params are added.
Will only generate comments for key that are pointing to values, lists or empty objects.
For empty objects {} deriving type is not working, recommend to add manually

## To do
- Derive 
- | multi-line strings does not work
- add sections? could add now but then would need section comment for all top level objects otherwise formatting would be wierd
- if sections added, add possiblity to add section descr?
- add @header comment type? to generate header description?
- add run flags
  - -f [ pathToFile ] to run on a single file absolute path
  - --from-wd to run program from working dir
  - --git to find git root dir and find all helm charts
  - -t [ md | html | rtf ] output formatting
  - -o [ pathToDir ] output directory
