# Calculator

```mermaid
flowchart TD
    title[<u>1</u>]
    style title fill:#FFF,stroke:#FFF
    calc --> adder
    calc --> multiplier
```

```mermaid
flowchart TD
    title[<u>2</u>]
    style title fill:#FFF,stroke:#FFF
    calc --> adder[go run .\nadder]
    calc --> multiplier
    style adder fill:#e74c3c,stroke:black
```

```mermaid
flowchart TD
    title[<u>3</u>]
    style title fill:#FFF,stroke:#FFF
    calc[weaver exec\ncalc] --> adder[weaver exec\nadder]
    calc --> multiplier
    style calc fill:#e74c3c,stroke:black
    style adder fill:#e74c3c,stroke:black
```

```mermaid
flowchart TD
    title[<u>3b</u>]
    style title fill:#FFF,stroke:#FFF
    calc[weaver kube\ncalc] --> adder[weaver kube\nadder]
    calc --> multiplier
    style calc fill:#e74c3c,stroke:black
    style adder fill:#e74c3c,stroke:black
```

```mermaid
flowchart TD
    title[<u>4</u>]
    style title fill:#FFF,stroke:#FFF
    calc[weaver exec\ncalc] --> adder[weaver exec\nadder]
    calc --> multiplier[weaver exec\nmultiplier]
    style calc fill:#e74c3c,stroke:black
    style adder fill:#e74c3c,stroke:black
    style multiplier fill:#e74c3c,stroke:black
```

```mermaid
flowchart TD
    title[<u>5</u>]
    style title fill:#FFF,stroke:#FFF
    calc[weaver kube\ncalc] --> adder[weaver kube\nadder]
    calc --> multiplier[weaver kube\nmultiplier]
    style calc fill:#e74c3c,stroke:black
    style adder fill:#e74c3c,stroke:black
    style multiplier fill:#e74c3c,stroke:black
```

|    | Go  | Proto | YAML | TOML | Docker |
| -- | --- | ----- | ---- | ---- | ------ |
| 1  | 167 | 24    | 95   | 0    | 33     |
| 2  | 184 | 24    | 116  | 0    | 33     |
| 3  | 136 | 12    | 140  | 0    | 24     |
| 3b | 136 | 12    | 31   | 9    | 11     |
| 4  | 99  | 0     | 153  | 0    | 12     |
| 5  | 99  | 0     | 0    | 7    | 0      |
