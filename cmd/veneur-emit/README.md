`veneur-emit` is a command line utility for emitting metrics to [Veneur](https://github.com/stripe/veneur).

Some common use cases:
* Instrument shell scripts
* Instrumenting shell-based tools like init scripts, startup scripts and more
* Testing

# Usage

Emitting a metric with veneur-emit:

```
$ veneur-emit -hostport udp://example.com:8125 -count 3 -name "my.test.metric" -tag "host:my.machine.local"
```

Full usage:

```
Usage of veneur-emit:
  -command string
    	Command to time. This will exec 'command', time it, and emit a timer metric.
  -count int
    	Report a 'count' metric. Value must be an integer.
  -debug
    	Turns on debug messages.
  -e_aggr_key string
    	Add an aggregation key to group event with others with same key.
  -e_alert_type string
    	Alert type must be 'error', 'warning', 'info', or 'success'. (default "info")
  -e_event_tags string
    	Tag(s) for event, comma separated. Ex: 'service:airflow,host_type:qa'
  -e_hostname string
    	Hostname for the event.
  -e_priority string
    	Priority of event. Must be 'low' or 'normal'. (default "normal")
  -e_source_type string
    	Add source type to the event.
  -e_text string
    	Text of event. Insert line breaks with an esaped slash (\\n) *
  -e_time string
    	Add timestamp to the event. Default is the current Unix epoch timestamp.
  -e_title string
    	Title of event. Ex: 'An exception occurred' *
  -gauge float
    	Report a 'gauge' metric. Value must be float64.
  -hostport string
    	Address of destination (hostport or listening address URL).
  -mode string
    	Mode for veneur-emit. Must be one of: 'metric', 'event', 'sc'. (default "metric")
  -name string
    	Name of metric to report. Ex: 'daemontools.service.starts'
  -sc_hostname string
    	Add hostname to the event.
  -sc_msg string
    	Message describing state of current state of service check.
  -sc_name string
    	Service check name. *
  -sc_status string
    	Integer corresponding to check status. (OK = 0, WARNING = 1, CRITICAL = 2, UNKNOWN = 3)*
  -sc_tags string
    	Tag(s) for service check, comma separated. Ex: 'service:airflow,host_type:qa'
  -sc_time string
    	Add timestamp to check. Default is current Unix epoch timestamp.
  -shellCommand
    	Turns on timeCommand mode. veneur-emit will grab everything after the first non-known-flag argument, time its execution, and report it as a timing metric.
  -ssf
    	Sends packets via SSF instead of StatsD. (https://github.com/stripe/veneur/blob/master/ssf/)
  -tag string
    	Tag(s) for metric, comma separated. Ex: 'service:airflow'
  -timing duration
    	Report a 'timing' metric. Value must be parseable by time.ParseDuration (https://golang.org/pkg/time/#ParseDuration).
```
