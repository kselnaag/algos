module github.com/kselnaag/algos

go 1.19

replace (
	github.com/kselnaag/algos/array => ./array
	github.com/kselnaag/algos/graph => ./graph
	github.com/kselnaag/algos/list => ./list
	github.com/kselnaag/algos/amath => ./amath
	github.com/kselnaag/algos/radix => ./radix
	github.com/kselnaag/algos/tree => ./tree
	github.com/kselnaag/algos/types => ./types
)

require (
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
