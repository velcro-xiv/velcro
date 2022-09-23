# velcro
Archive data from [xivsniff](https://github.com/velcro-xiv/xivsniff).

## Viewing your data
[DBeaver](https://dbeaver.io/) is a useful tool for viewing archived packet data.
It has a built-in hex viewer for SQLite `BLOB` columns, which can be used to save a blob to a file for analysis
in a tool such as [ImHex](https://imhex.werwolv.net/) or [010](https://www.sweetscape.com/010editor/).

With DBeaver (or any other SQLite viewer), packets can be queried using SQL statements, and grouped by data such as their opcodes
or source IP addresses. DBeaver also supports directly opening a hexdump in an external editor, making the workflow very easy to work with.