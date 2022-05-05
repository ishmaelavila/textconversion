# Text Converter

A WIP toy project to convert text from various sources into other formats for my uses.


## AMEX To CSV

This is currently the only converter in this repo. The American Express site provides several ways to export transactions (CSV, QBO, etc) but none of them include
pending transactions. This is painful to me because I update my ledger almost daily. This utility allows me to directly copy paste the text from the site into a file,
run the executable via CLI and receive a nicely formatted CSV including pending transactions that I can upload to my ledger.
