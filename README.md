Below are the two exercises that Chaos/Performance Engineer candidates are requested to complete as a foundation for the technical evaluation with Fleet. Please respect the 2-2.5 hour time limits as specified below.

Your solutions to the following exercises will be used as the basis for discussion in interviews to follow.

## Performance Debugging

Please spend _30-60 minutes_ digging into the performance issue described below. Document everything you learned and understand about the situation in [perf.md](./perf.md). Your goal here is to show the Fleet team that you can debug performance issues and dig down to provide as much relevant detail as possible towards understanding the problem.

The following is based on a real-life debugging scenario we encountered with an open-source user of Fleet.

The user reports that Fleet server CPU utilization is extremely high (hitting 100%) with around 2,000 hosts connecting. The user reports seeing errors in the Fleet logs like the following:

``` 
2022/01/10 10:15:32 http: TLS handshake error from 10.3.44.19:15894: EOF
```

To the Fleet team, this seems like an unusually high CPU utilization for the number of connected hosts. We ask the user to collect a [debug archive](https://fleetdm.com/docs/using-fleet/monitoring-fleet#generate-debug-archive-fleet-3-4-0) so that we can perform further analysis.

Included here are the contents of two debug archives:

- [./pprof-normal](./pprof-normal) is profiles representative of a typical Fleet server (Fleet 4.8.0).
- [./pprof-bad-cpu](./pprof-bad-cpu) is profiles representative of the problem described by the user (Fleet 4.8.0).

The various profiles are collected with, and can be analyzed by Go's [pprof tools](https://github.com/google/pprof/blob/master/doc/README.md).

An example invocation of the tools:

``` sh
go tool pprof -http localhost:1336 pprof-normal/profile
```

This will open a web browser with interactive visualization of the CPU profile.

It may also be helpful to look at the [Fleet server source](https://github.com/fleetdm/fleet/) for context, though the profiles provide everything strictly necessary for analysis.


## Testing & Monitoring

Please spend _90 minutes_ on this exercise.

In the [./wordgame-server](./wordgame-server) directory, we have provided a sample server, written in Go, implementing a simple word game. Your task is to build tooling and instrumentation to test and monitor this server.

While working within the time constraints, please consider what kind of test and monitoring tooling you would want to utilize if this server were going to production. If you utilize any infrastructure dependencies, please document or provide a docker-compose file for our testing. If there is more to note that would need to be carried out within separate infrastructure or does not fit into the time range allotted, please feel free to note this in [./wordgame-server/README.md](./wordgame-server/README.md).

Note that this task is about testing and monitoring the server, not developing it. Please feel free to modify the server code as necessary for these goals, and to improve correctness and reliability for anything that you notice. Implementing a new datastore or similar is out of scope; the server should behave substantially the same.
