# sqlinflux

sqlinflux is a util to make a sql query against an mssql database, where the sql query returns a set of keys and values (normally a group by rollup) to insert into influxdb for charting.

## Installation

1. go build sqlinflux.go
2. copy sqlinflux and a config file to a disk location where it can be called from telegraf (e.g. /opt/sqlinflux)
3. configure telegraf to call sqlinflux with the config file as a parameter:
  [[inputs.exec]]
    commands = [
      "/opt/sqlinflux/sqlinflux /opt/sqlinflux/orionnpm.yaml",
      "/opt/sqlinflux/sqlinflux /opt/sqlinflux/yourquery.yaml" 
    ]
    data_format = "influx"

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

v0.1 - Initial release

## License

MIT License
