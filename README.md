zeit
----

```
                          ███████╗███████╗██╗████████╗                             
                          ╚══███╔╝██╔════╝██║╚══██╔══╝
                            ███╔╝ █████╗  ██║   ██║   
                           ███╔╝  ██╔══╝  ██║   ██║   
                          ███████╗███████╗██║   ██║   
                          ╚══════╝╚══════╝╚═╝   ╚═╝   
```

Zeit erfassen. A command line tool for tracking time spent on tasks & projects.

![zeit](documentation/header.jpg)


## Build

```sh
make
```

**Info**: This will build using the version 0.0.0. You can prefix the `make` 
command with `VERSION=x.y.z` and set `x`, `y` and `z` accordingly if you want 
the version in `zeit --help` to be a different one.


## Usage

Please make sure to `export ZEIT_DB=~/.config/zeit.db` (or whatever location 
you would like to have the zeit database at).

*zeit*'s data structure contains of the following key entities: `project`, 
`task` and `entry`. An `entry` consists of a `project` and a `task`. These
don't have to pre-exist and can be created on-the-fly inside a new `entry` using
e.g. `zeit track --project "New Project" --task "New Task"`. In order to
configure them, the `zeit project` and the `zeit task` commands can be utilised.


### Projects

A project can be configured using `zeit project`:

```sh
zeit project --help
```

#### Examples:

Set the project color to a hex color code, allowing `zeit stats` to display
information in that color (if your terminal supports colours):

```sh
zeit project --color '#d3d3d3' "cool project"
```


### Task

A task can be configured using `zeit task`:

```sh
zeit task --help
```

#### Examples:

Setting up a Git repository to have commit messages automatically imported
into the activity notes when an activity is finished:

```sh
zeit task --git ~/my/git/repository "development"
```

**Info:** You will have to have the `git` binary available in your `PATH` for 
this to work. *zeit* automatically limits the commit log to the exact time of 
the activity's beginning- and finish-time. Commit messages before or after these 
times won't be imported.


### Track activity

```sh
zeit track --help
```

#### Examples:

Begin tracking a new activity and reset the start time to 15 minutes ago:

```sh
zeit track --project project --task task --begin -0:15
```


### Show current activity

```sh
zeit tracking
```


### Finish tracking activity

```sh
zeit finish --help
```

#### Examples:

Finish tracking the currently tracked activity without adding any further info:

```sh
zeit finish
```

Finish tracking the currently tracked activity and change its task:

```sh
zeit finish --task other-task
```

Finish tracking the currently tracked activity and adjust its start time to 
4 PM:

```sh
zeit finish --begin 16:00
```


### List tracked activity

```sh
zeit list
```


### Erase tracked activity

```sh
zeit erase --help
```

#### Examples:

Erase a tracked activity by its internal ID:

```sh
zeit erase 14037730-5c2d-44ff-b70e-81f1dcd4eb5f
```


### Display statistics

![zeit stats](documentation/zeit_stats.png)

```sh
zeit stats
```

### Import tracked activities

```sh
zeit import --help
```

The following formats are supported as of right now:

#### Tyme 3 JSON

It is possible to import JSON exports from [Tyme 3](https://www.tyme-app.com). 
It is important that the JSON is exported with the following options set/unset:

![Tyme 3 JSON export](documentation/tyme3json.png)

- `Start`/`End` can be set as required
- `Format` has to be `JSON`
- `Export only unbilled entries` can be set as required
- `Mark exported entries as billed` can be set as required
- `Include non-billable tasks` can be set as required
- `Filter Projects & Tasks` can be set as required
- `Combine times by day & task` **must** be unchecked

During import, *zeit* will create SHA1 sums for every Tyme 3 entry, which 
allows it to identify every imported activity. This way *zeit* won't import the 
exact same entry twice. Keep this in mind if you change entries in Tyme and 
then import them again into *zeit*.

#### Examples:

Import a Tyme 3 JSON export:

```sh
zeit import --tyme ./tyme.export.json
```

### Export tracked activities

```sh
zeit export --help
```

The following formats are supported as of right now:

#### Tyme 3 JSON

It is possible to export JSON compatible to the Tyme 3 JSON format. Fields that
are not available in *zeit* will be filled with dummy values, e.g.
`Billing: "UNBILLED"`.

#### Examples:

Export a Tyme 3 JSON:

```sh
zeit export --tyme --project "my project" --since "2020-04-01T15:04:05+07:00" --until "2020-04-04T15:04:05+07:00"
```
