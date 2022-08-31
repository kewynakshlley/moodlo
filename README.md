# moodlo

simple app only one endpoint that spits random feelings based on categories such as: happiness, fear, sadness and anger.


- /emotions/{category}
    * method: GET
    * description: returns random 
Response example:

```
{
    "Feeling": "feels hurt"
}
```


TODO: 

- Name of the person/character, ex: Kewyn is feeling hurt.
- Convert wordlists to a database schema.
- Add emotes to represent the emotion.
- Past/future events: Kewyn felt anger, Kewyn will feel anger
- Commandline interaction 


## Built With

* [Gorilla MUX](https://github.com/gorilla/mux) - Requests multiplexer used to HTTP routing and URL matcher

