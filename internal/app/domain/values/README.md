# Values

Here we save our domain value objects. Compared to entities, value objects are immutable and typically do not have unique identifier.

For example a workday of a particular shop:
```golang
type Workday struct {
  Date string
  isWorking bool
}
```