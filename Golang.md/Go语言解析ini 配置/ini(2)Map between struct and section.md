### Map To Struct

Want more objective way to play with INI? Cool.

```
Name = Unknwon
age = 21
Male = true
Born = 1993-01-01T20:17:05Z

[Note]
Content = Hi is a good man!
Cities = HangZhou, Boston

```

```
type Note struct {
    Content string
    Cities  []string
}

type Person struct {
    Name string
    Age  int `ini:"age"`
    Male bool
    Born time.Time
    Note
    Created time.Time `ini:"-"`
}

func main() {
    cfg, err := ini.Load("path/to/ini")
    // ...
    p := new(Person)
    err = cfg.MapTo(p)
    // ...

    // Things can be simpler.
    err = ini.MapTo(p, "path/to/ini")
    // ...

    // Just map a section? Fine.
    n := new(Note)
    err = cfg.Section("Note").MapTo(n)
    // ...
}

```

Can I have default value for field? Absolutely.

Assign it before you map to struct. It will keep the value as it is if the key is not presented or got wrong type.

```
// ...
p := &Person{
    Name: "Joe",
}
// ...

```

It’s really cool, but what’s the point if you can’t give me my file back from struct?

### Reflect From Struct[](https://ini.unknwon.io/docs/advanced/map_and_reflect#reflect-from-struct)

Why not?

```
type Embeded struct {
    Dates  []time.Time `delim:"|" comment:"Time data"`
    Places []string    `ini:"places,omitempty"`
    None   []int       `ini:",omitempty"`
}

type Author struct {
    Name      string `ini:"NAME"`
    Male      bool
    Age       int `comment:"Author's age"`
    GPA       float64
    NeverMind string `ini:"-"`
    *Embeded `comment:"Embeded section"`
}

func main() {
    a := &Author{"Unknwon", true, 21, 2.8, "",
        &Embeded{
            []time.Time{time.Now(), time.Now()},
            []string{"HangZhou", "Boston"},
            []int{},
        }}
    cfg := ini.Empty()
    err = ini.ReflectFrom(cfg, a)
    // ...
}

```

So, what do I get?

```
NAME = Unknwon
Male = true
; Author's age
Age = 21
GPA = 2.8

; Embeded section
[Embeded]
; Time data
Dates = 2015-08-07T22:14:22+08:00|2015-08-07T22:14:22+08:00
places = HangZhou,Boston

```

### Map with ShadowLoad[](https://ini.unknwon.io/docs/advanced/map_and_reflect#map-with-shadowload)

If you want to map a section to a struct along with [ShadowLoad](https://ini.unknwon.io/docs/howto/work_with_keys#same-key-with-multiple-values), then you need to indicate `allowshadow` in the struct tag.

For example, suppose you have the following configuration:

```
[IP]
value = 192.168.31.201
value = 192.168.31.211
value = 192.168.31.221

```

You should define your struct as follows:

```
type IP struct {
   Value    []string `ini:"value,omitempty,allowshadow"`
}

```

In case you don’t need the first two tag rules, then you can just have `ini:",,allowshadow"`.

### Other Notes On Map/Reflect[](https://ini.unknwon.io/docs/advanced/map_and_reflect#other-notes-on-map%2Freflect)

Any embedded struct is treated as a section by default, and there is no automatic parent-child relations in map/reflect feature:

```
type Child struct {
    Age string
}

type Parent struct {
    Name string
    Child
}

type Config struct {
    City string
    Parent
}

```

Example configuration:

```
City = Boston

[Parent]
Name = Unknwon

[Child]
Age = 21

```

What if, yes, I’m paranoid, I want embedded struct to be in the same section. Well, all roads lead to Rome.

```
type Child struct {
    Age string
}

type Parent struct {
    Name string
    Child `ini:"Parent"`
}

type Config struct {
    City string
    Parent
}

```

Example configuration:

```
City = Boston

[Parent]
Name = Unknwon
Age = 21

```

See also [Customize name and value mappers](https://ini.unknwon.io/docs/advanced/name_and_value_mapper).
