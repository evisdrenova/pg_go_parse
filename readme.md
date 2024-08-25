# WIP

This repo is a WIP

# Introduction

A native go library can parse a postgres sql statement and return the postgres syntax tree of that sql statement

# Motivation

The current tools, [go-pgquery](https://github.com/wasilibs/go-pgquery?tab=readme-ov-file) and [pg_query_go](https://github.com/pganalyze/pg_query_go) rely on wasm and underlying c files to build the syntax tree. That might work for some folks but it's also a heavy handed approach that might take up too much memory.

So we want to write a native go library that can parse a sql statement represented as a string and return the postgres syntax tree without having to worry about building a wasm module with c files.

# todo

- make all statements case insensitive in parser
- update targetlist in select stmt to be more flexible
