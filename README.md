## SUMRZ ("summarize")

Summarize a CSV file. Get quick stats about each column.

#### Build

    go build -o sumrz src/*.go

#### Install
On Linux, Mac, or Cygwin:

    cp sumrz /usr/bin/

#### Binaries

You can download binaries for Linux, Mac, and Windows directly, they're in `bin`.

By the magic of Go, they're staticly linked and should work with no dependencies.


## Examples

#### Nasdaq symbols

```shell
$ wget ftp://ftp.nasdaqtrader.com/SymbolDirectory/nasdaqtraded.txt

$ head -3 nasdaqtraded.com

Nasdaq Traded|Symbol|Security Name|Listing Exchange|Market Category|ETF|Round Lot Size|Test Issue|Financial Status|CQS Symbol|NASDAQ Symbol
Y|A|Agilent Technologies, Inc. Common Stock|N| |N|100|N||A|A
Y|AA|Alcoa Inc. Common Stock|N| |N|100|N||AA|AA

$ sumrz '|' < nasdaqtraded.txt

Table stats - 11 columns, 8012 rows
=================================

'Nasdaq Traded' all text
All values:
           7983 'Y'
             29 'N'

'Symbol' all text
Random sample:
                'CORR'
                'GE'
                'CONE'
                'IJJ'
                'GPI'
                'CII'
                'FCS'
                'OKS'
                'CU'
                'SBSA'

'Security Name' all text
Random sample:
                'United-Guardian, Inc. - Common Stock'
                'Thermon Group Holdings, Inc. Common Stock'
                'U.S. Bancorp Depositary Shares Representing 1/1000th Interest in a Shares Series F'
                'iShares MSCI Agriculture Producers Fund'
                'HomeAway, Inc. - Common Stock'
                'PIMCO Foreign Currency Strategy ETF'
                'Wells Fargo & Company Dep Shs Repstg 1/1000th Int Perp Pfd Cl A (Ser R Fixed To Flltg)'
                'Quidel Corporation - Common Stock'
                'Synta Pharmaceuticals Corp. - Common Stock'
                'Harmony Gold Mining Company Limited'

'Listing Exchange' all text
All values:
           3289 'N'
           2827 'Q'
           1437 'P'
            434 'A'
             25 'Z'

'Market Category' all text
All values:
           5185 ' '
           1620 'Q'
            624 'S'
            583 'G'

'ETF' all text
All values:
           6632 'N'
           1375 'Y'
              5 ' '

'Round Lot Size' all integers, min 1 max 100 avg 99.705691 stdev 99.832899
All values:
           7986 '100'
             24 '10'
              2 '1'

'Test Issue' all text
All values:
           7972 'N'
             40 'Y'

'Financial Status' 5185 blanks, 2827 text
All values:
           2770 'N'
             49 'D'
              6 'E'
              2 'H'

'CQS Symbol' 2827 blanks, 5185 text
Random sample:
                'MCN'
                'UBPpD'
                'ATHM'
                'WES'
                'GHYG'
                'LND'
                'ITT'
                'MA'
                'FFG'
                'MOD'

'NASDAQ Symbol' all text
Random sample:
                'EFZ'
                'NU'
                'EWV'
                'MEIL'
                'NCT-D'
                'PANW'
                'MHD'
                'ROIQ'
                'LMCA'
                'BTG'
```


## Docs

`sumrz` takes one optional argument: the separator character. The default is tab. Standard input should be a CSV with a header row, and it should be UTF8 or ASCII encoded.

