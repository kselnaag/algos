module algos

go 1.19

replace (
	algos/array => ./array
	algos/graph => ./graph
	algos/list => ./list
	algos/math => ./math
	algos/radix => ./radix
	algos/tree => ./tree
	algos/types => ./types
)

require (
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
