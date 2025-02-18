# stat-go-release

Statistic for [Go Release](https://go.dev/doc/devel/release).

## run
local
```
go run main.go > release.csv
uv sync
uv run jupyter nbconvert --to notebook --execute --inplace release.ipynb
```

## output

all releases

<img src="./releaseAll.png" alt="" width="500">

latest 10 releases

<img src="./releaseLatest10.png" alt="" width="500">
