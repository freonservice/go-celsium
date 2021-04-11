# go-celsium
Library for easy named formatting translations

### Documentation

All translations with named parameters are stored in next format:

`Hello, {name}, how are you?`

Where **{name}** it's named parameter which will converting in necessary string from dictionary (see example below).

```	
text, err := Format(&api.Translation{
    Text: "Hello, {name}, how are you?",
    IdentifierNamedList: []string{"name"},
}, map[string]string{
    "name": "Jhon",
})
```

As you can see, its pretty easy to converting raw string to production ready text.

Object **api.Translation** taken from GRPC response from freon service.
