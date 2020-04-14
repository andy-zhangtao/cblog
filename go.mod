module cblog

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/qiniu/api.v7/v7 v7.4.1
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
