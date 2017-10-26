# Rent
Compute your bedspace monthly rent cost.

### Install
`$ go get github.com/nmeji/rent`

### Usage
`$ rent monthly_rent total_utilities head_count`

### Example
```
$ rent 3166.667 3522.41 6
[Tenant #1] Name: Mark Ian
# of days stayed: 19
[Tenant #2] Name: Francis Mark
# of days stayed: 21
[Tenant #3] Name: Jason
# of days stayed: 21
[Tenant #4] Name: Jeric
# of days stayed: 21
[Tenant #5] Name: Niko
# of days stayed: 28
[Tenant #6] Name: Alden
# of days stayed: 18

[
  {
    "name": "Mark Ian",
    "days": 19,
    "rent": 3689.5247
  },
  {
    "name": "Francis Mark",
    "days": 21,
    "rent": 3744.5624
  },
  {
    "name": "Jason",
    "days": 21,
    "rent": 3744.5624
  },
  {
    "name": "Jeric",
    "days": 21,
    "rent": 3744.5624
  },
  {
    "name": "Niko",
    "days": 28,
    "rent": 3937.1942
  },
  {
    "name": "Alden",
    "days": 18,
    "rent": 3662.0059
  }
]
```

### How it works
1. Sum all the days you stayed in the last month
2. Compute your rate(%) of stay by averaging to totalStays
3. Calculate the monthly rent per each tenant using the ff. formula:
`utilities_total * average + monthly_rent`
