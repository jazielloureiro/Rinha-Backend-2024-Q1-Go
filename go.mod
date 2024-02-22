module github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go

go 1.22

require github.com/lib/pq v1.10.9

replace github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/persistence => ./internal/persistence

replace github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/helper => ./internal/helper

replace github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity => ./internal/entity
