# ADR 005: Enabled snapshot by default

## Changelog

- September 22th, 2021: Initial draft;
- September 23th, 2021: Moved from DRAFT to PROPOSED

## Status

PROPOSED

## Abstract

Currently, most of the validators or full nodes does not have snapshot enabled because they use the default 
setting whose snapshot is not enabled. This makes state sync a new node very difficult. 
In oder to increase the number of snapshots in the network so that making creating a new node easily through 
`StateSync`, we SHOULD enable snapshot by default.  

## Context

Tendermint provides an alternative mechanism called `StateSync` for bootstrapping a new node without downloading 
blocks or going through consensus. It fetches a snapshot of the state machine at a given height and restores it.
`StateSync` will greatly improve the experience of joining a network, reducing the time required to sync a node
by several orders of magnitude.
However, state sync a node is almost impossible now since there is too few snapshot in the network now. In order to ensure 
the number of snapshots is enough, we SHOULD encourage validators and full nodes to set snapshot enabled. 
As the result, making it enabled by default can be more convenient to who wants to provide the snapshot to the network. 
In addition, it would increase the willing of validators or full nodes to enable snapshot as well.

## Decision

The implementation idea is to build a custom `initCmd` including `StatSync` setting to init the `config/app.toml` 
which is based on the cosmos-sdk one. In addition, the pruning related fields SHOULD be fixed into custom values 
in order to make sure they are compatible with `StateSync` since `snapshot-interval` MUST be a multiple of `pruning-keep-every`.

```go
// initCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func initCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {
	cmd := cosmosgenutilcli.InitCmd(mbm, defaultNodeHome)
	defaultRun := cmd.RunE
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx := client.GetClientContextFromCmd(cmd)
		err := defaultRun(cmd, args)
		if err != nil {
			return err
		}
		// Set app.toml file
		appConfig := srvconfig.DefaultConfig()
		appConfig.Pruning = "custom"
		appConfig.PruningKeepRecent = "100"
		appConfig.PruningKeepEvery = "500"
		appConfig.PruningInterval = "10"

		appConfig.StateSync.SnapshotInterval = 500
		srvconfig.WriteConfigFile(filepath.Join(clientCtx.HomeDir, "config", "app.toml"), appConfig)
		return nil
	}
	return cmd
}
```

Finally, replacing `cosmosgenutilcli.InitCmd` into the custom `initCmd` in `rootCmd`.
```go
rootCmd.AddCommand(
    initCmd(app.ModuleBasics, app.DefaultNodeHome),
    ...
)
```


## Consequences

### Backwards Compatibility

The change only includes replacing the `initCmd` to custom one to enable snapshot so the backwards compatibility is 
not relevant as there won't be any issue related to it.

### Positive

* Increase the number of snapshots in the network
* Make creating a new node easily by StateSync

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#469](https://github.com/desmos-labs/desmos/issues/469)